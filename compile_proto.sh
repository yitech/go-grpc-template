#!/bin/bash

protoc \
    --go_out=. \
    --go-grpc_out=. \
    proto/helloworld.proto