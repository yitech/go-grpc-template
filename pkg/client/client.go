package client

import (
	v1 "github.com/yitech/go-grpc-template/pkg/client/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn *grpc.ClientConn
}

func NewClient(address string) (IClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &Client{
		conn: conn,
	}, nil
}

func (c *Client) V1() v1.IV1 {
	return v1.NewV1(c.conn)
}

func (c *Client) Close() error {
	// Close all clients
	return c.conn.Close()
}
