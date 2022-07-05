package svc

import (
	"github.com/l6l6ng/go-zero-demo/book/service/user/api/internal/config"
	"github.com/l6l6ng/go-zero-demo/book/service/user/api/internal/middleware"
	"github.com/l6l6ng/go-zero-demo/book/service/user/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
	Rbac rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
		Rbac: middleware.NewRbacMiddleware().Handle,
	}
}
