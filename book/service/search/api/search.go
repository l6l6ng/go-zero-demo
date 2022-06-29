package main

import (
	"flag"
	"fmt"
	"github.com/l6l6ng/go-zero-demo/book/service/search/api/internal/handler"
	"github.com/l6l6ng/go-zero-demo/book/service/search/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"github.com/l6l6ng/go-zero-demo/book/service/search/api/internal/config"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/search-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	//全局中间件
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			logx.Info("global middleware2222222")
			next(w, r)
		}
	})

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
