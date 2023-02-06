package logic

import (
	"context"
	"react-blog-server/common/errorx"

	"react-blog-server/apps/article/internal/svc"
	"react-blog-server/apps/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleListLogic {
	return &GetArticleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleListLogic) GetArticleList() (resp *types.GetArticleListResp, err error) {
	articles, err := l.svcCtx.ArticleModel.Find(l.ctx, 10)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.DATABASE_MYSQL_INTERNAL_ERROR, err.Error())
	}
	var res []*types.Article
	for _, a := range articles {

		res = append(res, &types.Article{
			Title:       a.Title,
			ArticleId:   a.Id,
			ImageURL:    a.ImageUrl,
			Desc:        a.Describes.String,
			PublishTime: a.PublishTime.Format("2006-01-02"),
			UpdateTime:  a.UpdateTime.Format("2006-01-02"),
		})
	}
	return &types.GetArticleListResp{Data: res, Meta: types.Meta{Status: 200, Msg: "succeed"}}, nil
}
