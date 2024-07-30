package helloworld

import "context"

type IHelloWorld interface {
	// SayHello sends a greeting to the server.
	SayHello(context.Context, string) (string, error)
	Close() error
}
