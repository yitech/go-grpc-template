package client

import v1 "github.com/yitech/go-grpc-template/pkg/client/v1"

type Client struct {
	address string
}

func NewClient(address string) IClient {
	return &Client{
		address: address,
	}
}

func (c *Client) V1() v1.IV1 {
	return v1.NewV1(c.address)
}
