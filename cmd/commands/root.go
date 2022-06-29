package commands

import (
	"fmt"
	cfg "github.com/bhatti/GSSI/internal/config"
	"github.com/bhatti/GSSI/internal/server"
	"github.com/bhatti/GSSI/internal/version"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var cfgFile string

var config *cfg.Config
var logFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "platform-poc",
	Short: "Start an instance of Server",
	Long:  "Start an instance of Server",
	Run: func(cmd *cobra.Command, args []string) {
		_ = rootRun(cmd, args)
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"",
		"config file")

}

func rootRun(_ *cobra.Command, _ []string) error {
	_, err := server.StartServers(config)
	if err != nil {
		log.WithField("Error", err).Fatalf("could start for gRPC server")
		return err
	}
	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version, commit, date string) error {
	Version = version
	Commit = commit
	Date = date
	log.Infof("Starting root command...")
	if err := rootCmd.Execute(); err != nil {
		log.WithField("Error", err).
			Fatalf("Could not start server")

		return err
	}

	return nil
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	var err error
	if config, err = cfg.New(cfgFile); err != nil {
		panic(err)
	}
	config.Version = version.New(Version, Commit, Date)

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	if logFile != "" {
		fmt.Printf("routing logs to %s\n", logFile)
		if f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666); err == nil {
			log.SetOutput(f)
			log.SetFormatter(&log.JSONFormatter{})
		}
	}
	log.SetLevel(log.InfoLevel)
}
