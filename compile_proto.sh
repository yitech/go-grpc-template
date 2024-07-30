#!/bin/bash

protoc \
    --go_out=. \
    --go-grpc_out=. \
    proto/v1/helloworld.proto