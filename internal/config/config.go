package config

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/bhatti/GSSI/internal/auth"
	"github.com/bhatti/GSSI/internal/version"
	"github.com/gomodule/redigo/redis"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"time"
)

// PersistenceProvider defines enum for persistence provider
type PersistenceProvider string

const (
	// RedisPersistenceProvider uses redis
	RedisPersistenceProvider PersistenceProvider = "REDIS"

	// MongoPersistenceProvider uses MongoDB
	MongoPersistenceProvider PersistenceProvider = "MONGO"

	// MemoryPersistenceProvider uses in-memory
	MemoryPersistenceProvider PersistenceProvider = "MEMORY"
)

// RedisConfig redis config
type RedisConfig struct {
	Host       string        `yaml:"host" mapstructure:"host"`
	Port       int           `yaml:"port" mapstructure:"port"`
	Password   string        `yaml:"password" mapstructure:"password"`
	TTLMinutes time.Duration `yaml:"ttl_minutes" mapstructure:"ttl_minutes"`
	PoolSize   int           `yaml:"pool_size" mapstructure:"pool_size"`
	MaxPopWait time.Duration `yaml:"max_subscription_wait" mapstructure:"max_subscription_wait"`
	Pool       *redis.Pool   `yaml:"-"`
}

// MongoConfig -- mongo config
type MongoConfig struct {
	Name     string          `yaml:"name"`
	URL      string          `yaml:"url"`
	MaxLimit int32           `yaml:"max_limit"`
	Client   *mongo.Client   `yaml:"-"`
	Database *mongo.Database `yaml:"-"`
}

// Config -- Default Config
type Config struct {
	Mongo               MongoConfig         `yaml:"mongo" env:"MONGO"`
	Redis               RedisConfig         `yaml:"redis" env:"REDIS"`
	GrpcSasl            bool                `yaml:"grpc_sasl"`
	ListenPort          string              `yaml:"listen_port"`
	Debug               bool                `yaml:"debug"`
	Dir                 string              `yaml:"dir" env:"CONFIG_DIR"`
	Url                 string              `yaml:"url" env:"BASE_URL"`
	PersistenceProvider PersistenceProvider `yaml:"persistence_provider" env:"REPOSITORY_PROVIDER"`
	Version             *version.Info       `yaml:"-"`
}

// New -- initializes the Default Configuration
func New(configFile string) (*Config, error) {
	viper.SetDefault("debug", "false")
	viper.SetDefault("dir", "")
	viper.SetDefault("listen_port", "127.0.0.1:7777")
	viper.SetDefault("max_limit", "10000")
	viper.SetDefault("url", "")
	viper.SetDefault("persistence_provider", "MEMORY")
	viper.SetDefault("mongo.name", "tuple_space_db")
	viper.SetDefault("mongo.url", "mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb")
	viper.SetEnvPrefix("")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			return nil, err
		}
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.AddConfigPath("./config")
		viper.AddConfigPath("../config")
		viper.AddConfigPath("../../config")
	}
	var err error
	if err = viper.ReadInConfig(); err == nil {
		logrus.Infof("using config file: %s", viper.ConfigFileUsed())
	} else {
		logrus.Infof("could not read config %s", err)
	}
	var config Config
	if err = viper.Unmarshal(&config); err != nil {
		logrus.Fatalf("unable to decode into struct, %v", err)
		return nil, err
	}
	confDir := os.Getenv("CONFIG_DIR")
	if confDir != "" {
		config.Dir = confDir
	}
	if val := os.Getenv("BASE_URL"); val != "" {
		config.Url = strings.TrimSpace(val)
	}

	if config.PersistenceProvider == MongoPersistenceProvider {
		config.Mongo.Client, err = mongo.NewClient(options.Client().ApplyURI(config.Mongo.URL))
		if err != nil {
			return nil, err
		}
		err = config.Mongo.Client.Connect(context.Background())
		if err != nil {
			return nil, err
		}
		config.Mongo.Database = config.Mongo.Client.Database(config.Mongo.Name)
	}
	if err = config.validate(); err != nil {
		return nil, err
	}
	return &config, nil
}

func (c *Config) Authorizer() (auth.Authorizer, error) {
	aclFile, err := c.ACLModelFile()
	if err != nil {
		debug.PrintStack()
		return nil, err
	}
	policyFile, err := c.ACLPolicyFile()
	if err != nil {
		return nil, err
	}
	return auth.New(aclFile, policyFile), nil
}

func (c *Config) TLSClient() (tlsC TLSConfig, err error) {
	if !c.GrpcSasl {
		return
	}
	var certFile string
	var keyFile string
	var caFile string
	certFile, err = c.ClientCertFile()
	if err != nil {
		return
	}
	keyFile, err = c.ClientKeyFile()
	if err != nil {
		return
	}
	caFile, err = c.CAFile()
	if err != nil {
		return
	}
	tlsC = TLSConfig{
		CertFile: certFile,
		KeyFile:  keyFile,
		CAFile:   caFile,
	}
	return
}

func (c *Config) TLSRootClient() (tlsC TLSConfig, err error) {
	if !c.GrpcSasl {
		return
	}
	var certFile string
	var keyFile string
	var caFile string
	certFile, err = c.ClientRootCertFile()
	if err != nil {
		return
	}
	keyFile, err = c.ClientRootKeyFile()
	if err != nil {
		return
	}
	caFile, err = c.CAFile()
	if err != nil {
		return
	}
	tlsC = TLSConfig{
		CertFile: certFile,
		KeyFile:  keyFile,
		CAFile:   caFile,
	}
	return
}

func (c *Config) TLSNobodyClient() (tlsC TLSConfig, err error) {
	if !c.GrpcSasl {
		return
	}
	var certFile string
	var keyFile string
	var caFile string
	certFile, err = c.ClientNobodyCertFile()
	if err != nil {
		return
	}
	keyFile, err = c.ClientNobodyKeyFile()
	if err != nil {
		return
	}
	caFile, err = c.CAFile()
	if err != nil {
		return
	}
	tlsC = TLSConfig{
		CertFile: certFile,
		KeyFile:  keyFile,
		CAFile:   caFile,
	}
	return
}

func (c *Config) SetupTLSClient() (tlsConfig *tls.Config, err error) {
	tlsC, err := c.TLSClient()
	if err != nil {
		return nil, err
	}
	return tlsC.SetupTLS()
}

func (c *Config) SetupTLSServer(addr string) (tlsConfig *tls.Config, err error) {
	certFile, err := c.ServerCertFile()
	if err != nil {
		return nil, err
	}
	keyFile, err := c.ServerKeyFile()
	if err != nil {
		return nil, err
	}
	caFile, err := c.CAFile()
	if err != nil {
		return nil, err
	}
	tlsC := TLSConfig{
		CertFile:      certFile,
		KeyFile:       keyFile,
		CAFile:        caFile,
		ServerAddress: addr,
		Server:        true,
	}
	return tlsC.SetupTLS()
}
func (c *Config) CAFile() (string, error) {
	return configFile(c.Dir, "ca.pem")
}

func (c *Config) ServerCertFile() (string, error) {
	return configFile(c.Dir, "server.pem")
}

func (c *Config) ServerKeyFile() (string, error) {
	return configFile(c.Dir, "server-key.pem")
}

func (c *Config) ClientCertFile() (string, error) {
	return configFile(c.Dir, "client.pem")
}

func (c *Config) ClientKeyFile() (string, error) {
	return configFile(c.Dir, "client-key.pem")
}

func (c *Config) ClientRootCertFile() (string, error) {
	return configFile(c.Dir, "root-client.pem")
}

func (c *Config) ClientRootKeyFile() (string, error) {
	return configFile(c.Dir, "root-client-key.pem")
}

func (c *Config) ClientNobodyCertFile() (string, error) {
	return configFile(c.Dir, "nobody-client.pem")
}

func (c *Config) ClientNobodyKeyFile() (string, error) {
	return configFile(c.Dir, "nobody-client-key.pem")
}

func (c *Config) ACLModelFile() (string, error) {
	return configFile(c.Dir, "model.conf")
}

func (c *Config) ACLPolicyFile() (string, error) {
	return configFile(c.Dir, "policy.csv")
}

func configFile(dir string, filename string) (string, error) {
	f := filepath.Join(dir, filename)
	if st, err := os.Stat(f); err == nil {
		if st.IsDir() || !st.Mode().IsRegular() {
			return "", fmt.Errorf("%s is directory", f)
		}
		return f, nil
	} else {
		return "", fmt.Errorf("failed to find '%s' config file due to %s", f, err)
	}
}

func (c *MongoConfig) validate() error {
	if val := os.Getenv("MONGO_URL"); val != "" {
		c.URL = val
	}
	if c.Name == "" {
		c.Name = "tuple_space_db"
	}
	return nil
}

func (c *Config) validate() error {
	if c.PersistenceProvider == MongoPersistenceProvider {
		if err := c.Mongo.validate(); err != nil {
			return err
		}
	} else if c.PersistenceProvider == RedisPersistenceProvider {
		if err := c.Redis.validate(); err != nil {
			return err
		}
	}
	if c.Url == "" {
		b, _ := yaml.Marshal(c)
		fmt.Printf("%s\n", b)
		return fmt.Errorf("failed to find base-url: %s", os.Getenv("BASE_URL"))
	}

	return nil
}

// Validate - validates
func (c *RedisConfig) validate() error {
	if c.Host == "" {
		return errors.New("redis host is not set")
	}
	if c.Port == 0 {
		c.Port = 6379
	}
	if c.MaxPopWait == 0 {
		c.MaxPopWait = 60 * time.Second
	}
	if c.Host == "" || c.Port == 0 {
		return fmt.Errorf("redis is not configured %s:%d", c.Host, c.Port)
	}
	hostPort := fmt.Sprintf("%s:%d", c.Host, c.Port)
	c.Pool = &redis.Pool{
		MaxIdle:   c.PoolSize,
		MaxActive: c.PoolSize,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", hostPort)
			if err != nil {
				return nil, err
			}
			if c.Password != "" {
				if _, err := conn.Do("AUTH", c.Password); err != nil {
					_ = conn.Close()
					return nil, err
				}
			}
			return conn, err
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
	}
	return nil
}
