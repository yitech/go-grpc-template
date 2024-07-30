package v1

import (
	"github.com/yitech/go-grpc-template/pkg/client/v1/helloworld"
	"google.golang.org/grpc"
)

type V1 struct {
	conn *grpc.ClientConn
}

func NewV1(conn *grpc.ClientConn) IV1 {
	return &V1{
		conn: conn,
	}
}

func (v *V1) HelloWorld() helloworld.IHelloWorld {
	return helloworld.NewGreeterClient(v.conn)
}
