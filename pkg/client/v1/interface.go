package v1

import "github.com/yitech/go-grpc-template/pkg/client/v1/helloworld"

type IV1 interface {
	// SayHello sends a greeting to the server.
	HelloWorld() helloworld.IHelloWorld
}
