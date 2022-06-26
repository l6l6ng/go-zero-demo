package svc

import (
	"github.com/l6l6ng/go-zero-demo/book/service/user/api/internal/config"
	"github.com/l6l6ng/go-zero-demo/book/service/user/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	userModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		userModel: model.NewUserModel(conn, c.CacheRedis),
	}
}
