package helloworld

import (
	"context"

	pb "github.com/yitech/go-grpc-template/grpc/v1/helloworld"
	"google.golang.org/grpc"
)

// GreeterClient is the client API for Greeter service.
type GreeterClient struct {
	conn   *grpc.ClientConn
	client pb.GreeterClient
}

// NewGreeterClient creates a new GreeterClient.
func NewGreeterClient(conn *grpc.ClientConn) IHelloWorld {
	// Set up a connection to the server.
	client := pb.NewGreeterClient(conn)
	return &GreeterClient{
		conn:   conn,
		client: client,
	}
}

func (c *GreeterClient) SayHello(ctx context.Context, name string) (string, error) {
	// Contact the server and print out its response.
	r, _ := c.client.SayHello(ctx, &pb.HelloRequest{Name: name})
	return r.GetMessage(), nil
}

func (c *GreeterClient) Close() error {
	return c.conn.Close()
}
