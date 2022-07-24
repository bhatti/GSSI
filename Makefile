BRANCH := $(shell git rev-parse --symbolic-full-name --abbrev-ref HEAD)
SUMMARY := $(shell git describe --always --long --dirty)
VERSION := $(shell cat VERSION)
CONFIG_PATH=${HOME}/gssi-grpc-config

.PHONY: init
init:
	mkdir -p ${CONFIG_PATH}
	cp fixtures/model.conf ${CONFIG_PATH}
	cp fixtures/policy.csv ${CONFIG_PATH}

.PHONY: gencert

gencert: init
	cfssl gencert \
		-initca fixtures/ca-csr.json | cfssljson -bare ca

	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=fixtures/ca-config.json \
		-profile=server \
		fixtures/server-csr.json | cfssljson -bare server
	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=fixtures/ca-config.json \
		-profile=client \
		fixtures/client-csr.json | cfssljson -bare client
	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=fixtures/ca-config.json \
		-profile=client \
		-cn="root" \
		fixtures/client-csr.json | cfssljson -bare root-client
	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=fixtures/ca-config.json \
		-profile=client \
		-cn="nobody" \
		fixtures/client-csr.json | cfssljson -bare nobody-client
	mv *.pem *.csr ${CONFIG_PATH}

.PHONY: gencert
test: init
	CGO_ENABLED=1 go test -race ./internal/...

.PHONY: compile
compile: model-compile
	protoc `find api -name "*.proto"` \
		--go_out=. \
		--go-grpc_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--proto_path=.
	#protoc --js_out=import_style=commonjs,binary:. api/v1/*/*.proto --proto_path=.
model-compile:
	protoc `find api -name "*.proto"` \
                --go_out=. \
                --go_opt=paths=source_relative \
                --proto_path=.

vendor:
	go mod vendor

build: compile
	go build -mod=vendor -ldflags "-X main.Version=$(VERSION) -X main.Commit=$(BRANCH) -X main.Date=$(VERSION)"  -o server cmd/main.go

release-actual:
	gox -mod=vendor -ldflags= -osarch="linux/amd64 darwin/amd64" -output="release/server_{{.OS}}_{{.Arch}}" ./...


release: vendor release-actual


.PHONY: compile build release vendor
