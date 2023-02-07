package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"react-blog-server/apps/article/internal/config"
	"react-blog-server/apps/article/model"
)

type ServiceContext struct {
	Config       config.Config
	ArticleModel model.ArticleModel
	UserModel    model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:       c,
		ArticleModel: model.NewArticleModel(conn),
		UserModel:    model.NewUserModel(conn),
	}
}
