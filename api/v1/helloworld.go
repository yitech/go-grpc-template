package api

import (
	"context"

	pb "github.com/yitech/go-grpc-template/grpc/v1/helloworld"
)

// HelloWorld implements the GreeterServer interface.
type HelloWorld struct {
	pb.UnimplementedGreeterServer
}

// SayHello sends a greeting.
func (s *HelloWorld) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
