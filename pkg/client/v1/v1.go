package v1

import "github.com/yitech/go-grpc-template/pkg/client/v1/helloworld"

type V1 struct {
	address string
}

func NewV1(address string) IV1 {
	return &V1{
		address: address,
	}
}

func (v *V1) HelloWorld() helloworld.IHelloWorld {
	hw, err := helloworld.NewGreeterClient(v.address)
	if err != nil {
		return nil
	}
	return hw
}
