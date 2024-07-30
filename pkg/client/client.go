package client

import (
	v1 "github.com/yitech/go-grpc-template/pkg/client/v1"
	"google.golang.org/grpc"
)

type Client struct {
	conn *grpc.ClientConn
}

func NewClient(conn *grpc.ClientConn) IClient {
	return &Client{
		conn: conn,
	}
}

func (c *Client) V1() v1.IV1 {
	return v1.NewV1(c.conn)
}
