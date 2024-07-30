package main

import (
	"log"
	"net"

	api "github.com/yitech/go-grpc-template/api/v1/helloworld"
	pb "github.com/yitech/go-grpc-template/grpc/v1/helloworld"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &api.HelloWorld{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
