package svc

import (
	"github.com/l6l6ng/go-zero-demo/book/service/search/api/internal/config"
	"github.com/l6l6ng/go-zero-demo/book/service/search/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config  config.Config
	Example rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Example: middleware.NewExampleMiddleware().Handle,
	}
}
