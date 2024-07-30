package api

import (
	"context"

	pb "github.com/yitech/go-grpc-template/grpc/v1/helloworld"
)

type Greeter struct {
	pb.UnimplementedGreeterServer
}

func (s *Greeter) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
