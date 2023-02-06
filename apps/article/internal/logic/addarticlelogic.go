package logic

import (
	"context"

	"react-blog-server/apps/article/internal/svc"
	"react-blog-server/apps/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddArticleLogic {
	return &AddArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddArticleLogic) AddArticle(req *types.AddArticleReq) (resp *types.AddArticleResp, err error) {
	// todo: add your logic here and delete this line

	return
}
