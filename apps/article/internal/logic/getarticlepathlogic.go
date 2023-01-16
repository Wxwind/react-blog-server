package logic

import (
	"context"
	"react-blog-server/apps/article/internal/svc"
	"react-blog-server/apps/article/internal/types"
	"react-blog-server/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticlePathLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticlePathLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticlePathLogic {
	return &GetArticlePathLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticlePathLogic) GetArticlePath(req *types.GetArticlePathReq) (resp *types.GetArticlePathResp, err error) {
	res, err := l.svcCtx.ArticleModel.FindOne(l.ctx, req.ParticleId)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.DATABASE_MYSQL_NOT_EXISTS_ARTICLE, err.Error())
	}
	return &types.GetArticlePathResp{
		Data: res.MarkdownUrl,
		Meta: types.Meta{Status: 200, Msg: "succeed"},
	}, nil
}
