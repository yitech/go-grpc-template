package helloworld

import (
	"context"

	pb "github.com/yitech/go-grpc-template/grpc/v1/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GreeterClient is the client API for Greeter service.
type GreeterClient struct {
	pb.GreeterClient
}

// NewGreeterClient creates a new GreeterClient.
func NewGreeterClient(address string) (IHelloWorld, error) {
	// Set up a connection to the server.
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &GreeterClient{pb.NewGreeterClient(conn)}, nil
}

func (c *GreeterClient) SayHello(ctx context.Context, name string) (string, error) {
	// Contact the server and print out its response.
	r, err := c.GreeterClient.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		return "", err
	}
	return r.GetMessage(), nil
}
