package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"react-blog-server/apps/app/article/internal/config"
	"react-blog-server/apps/app/article/model"
)

type ServiceContext struct {
	Config       config.Config
	ArticleModel model.ArticleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:       c,
		ArticleModel: model.NewArticleModel(conn, nil),
	}
}
