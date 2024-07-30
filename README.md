<img src="https://go.dev/images/gophers/pink.svg" width="150" height="150">

# Go gRPC Template

This repository provides a basic template for building gRPC services in Go. It includes a simple gRPC server and client example, along with a Makefile for easy management of the build process.

## Directory Structure
```bash
.
├── api
│ └── v1
│ └── helloworld.go # Generated gRPC code
├── bin # Directory for binaries
├── cmd
│ ├── client
│ │ └── main.go # gRPC client implementation
│ └── server
│ └── main.go # gRPC server implementation
├── compile_proto.sh # Script to compile protobuf files
├── go.mod # Go module file
├── go.sum # Go dependencies file
├── grpc
│ └── v1
│ └── helloworld # Generated gRPC code
│ ├── helloworld_grpc.pb.go
│ └── helloworld.pb.go
├── LICENSE # License file
├── Makefile # Makefile for build automation
├── proto
│ └── v1
│ └── helloworld.proto # Protobuf file
└── README.md # This README file
```

## Getting Started

### Prerequisites

- Go 1.16 or higher
- Protocol Buffers compiler (`protoc`)
- Go plugins for `protoc`

### Install Protocol Buffers Compiler

Follow the [installation guide for `protoc`](https://grpc.io/docs/protoc-installation/).

### Install Go Plugins for `protoc`

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
Make sure $GOPATH/bin is in your PATH.

### Compile Protobuf Files
Use the compile_proto.sh script to compile the protobuf files:

```bash
make proto
```
### Build and Run
Use the Makefile to build the server and client binaries.

#### Build
```bash
make build
```
#### Run Server
```bash
./bin/server
```
#### Run Client
```bash
./bin/client
```

### Project Layout
- proto/v1/helloworld.proto: Protobuf definition for the helloworld service.
- grpc/v1/helloworld: Generated Go code for the helloworld service.
- cmd/server/main.go: Entrypoint of the gRPC server.
- cmd/client/main.go: Entrypoint of the gRPC client.
- compile_proto.sh: Script to compile protobuf files.
- Makefile: Makefile to automate the build process.

## License
This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgements
- gRPC
- Protocol Buffers
