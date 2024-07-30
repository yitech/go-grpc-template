package helloworld

import (
	"context"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/bmizerany/assert"
	"github.com/stretchr/testify/require"
	api "github.com/yitech/go-grpc-template/api/v1/helloworld"
	pb "github.com/yitech/go-grpc-template/grpc/v1/helloworld"
)

func TestHelloWorld_SayHello(t *testing.T) {
	// Create a listener on a random port
	lis, err := net.Listen("tcp", ":0")
	require.NoError(t, err)
	defer lis.Close()

	// Create a gRPC server and register the HelloWorld service
	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &api.HelloWorld{})

	// Start the server in a goroutine
	go grpcServer.Serve(lis)
	defer grpcServer.Stop()

	// Create a client to connect to the server
	conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	client := NewGreeterClient(conn)

	require.NoError(t, err)
	defer client.Close()

	// Create a context
	ctx := context.Background()

	// Make a request
	resp, err := client.SayHello(ctx, "test")
	require.NoError(t, err)

	// Assert the response
	assert.Equal(t, "Hello test", resp)
}
