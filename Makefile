# Makefile

# Define project-specific variables
MODULE = github.com/yitech/grpc-go-template
PROTO_FILES = $(PROTO_DIR)/helloworld.proto
GO_OUT_DIR = .

# Define output binary names
SERVER_BIN = bin/server
CLIENT_BIN = bin/client

# Go tools
PROTOC_GEN_GO = protoc-gen-go
PROTOC_GEN_GO_GRPC = protoc-gen-go-grpc

# Coverage file
COVERAGE_FILE = coverage.out
COVERAGE_HTML = coverage.html

.PHONY: all clean proto build server client

# Default target
all: proto build

# Clean up generated files and binaries
clean:
	rm -rf $(GO_OUT_DIR)/grpc/*
	rm -f $(SERVER_BIN) $(CLIENT_BIN)
	rm -f $(COVERAGE_FILE) $(COVERAGE_HTML)

# Generate Go code from protobuf
proto:
	sh compile_proto.sh

# Build both server and client binaries
build: server client

# Build server binary
server:
	go build -o $(SERVER_BIN) cmd/server/*.go

# Build client binary
client:
	go build -o $(CLIENT_BIN) cmd/client/*.go

lint:
	golangci-lint run

test:
	go test -coverprofile=$(COVERAGE_FILE) ./api/...
	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)