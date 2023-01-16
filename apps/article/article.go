package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"react-blog-server/apps/article/internal/config"
	"react-blog-server/apps/article/internal/handler"
	"react-blog-server/apps/article/internal/svc"
	"react-blog-server/common/errorx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/article-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.ToData()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.ToData()
		default:
			return http.StatusInternalServerError, nil
		}
	})
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
