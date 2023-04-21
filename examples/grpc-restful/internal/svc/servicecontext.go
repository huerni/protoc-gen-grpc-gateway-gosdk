package svc

import "github.com/jaronnie/protoc-gen-grpc-gateway-gosdk/grpc-restful/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
