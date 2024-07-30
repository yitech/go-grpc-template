package client

import v1 "github.com/yitech/go-grpc-template/pkg/client/v1"

// V1 is the client API for v1 service.
type IClient interface {
	V1() v1.IV1
}
