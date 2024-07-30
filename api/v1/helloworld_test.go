package api

import (
	"context"
	"testing"

	pb "github.com/yitech/go-grpc-template/grpc/v1/helloworld"
)

func TestSayHello(t *testing.T) {
	server := &HelloWorld{}
	req := &pb.HelloRequest{Name: "test"}
	resp, err := server.SayHello(context.Background(), req)
	if err != nil {
		t.Fatalf("SayHello failed: %v", err)
	}
	expected := "Hello test"
	if resp.GetMessage() != expected {
		t.Errorf("unexpected response: got %v, want %v", resp.GetMessage(), expected)
	}
}
