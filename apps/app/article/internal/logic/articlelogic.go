package logic

import (
	"context"
	"react-blog-server/apps/app/article/internal/svc"
	"react-blog-server/apps/app/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleLogic {
	return &ArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleLogic) Article() (resp *types.GetArticleListResp, err error) {
	articles, err := l.svcCtx.ArticleModel.Find(l.ctx, 10)
	if err != nil {
		return nil, err
	}
	var res []*types.Article
	for _, a := range articles {

		res = append(res, &types.Article{
			Title:       a.Title,
			ParticleId:  a.Id,
			ImageURL:    a.ImageUrl,
			Desc:        a.Describes.String,
			PublishTime: a.PublishTime.Format("2006-01-02"),
			UpdateTime:  a.UpdateTime.Format("2006-01-02"),
		})
	}
	return &types.GetArticleListResp{Data: res, Meta: types.Meta{Status: 200, Msg: "succeed"}}, nil
}
