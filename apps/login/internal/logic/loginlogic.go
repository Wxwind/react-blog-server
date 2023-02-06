package logic

import (
	"context"
	"errors"
	"react-blog-server/apps/login/internal/svc"
	"react-blog-server/apps/login/internal/types"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("参数错误")
	}

	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	if user.Password == req.Password {

	}
	return
}
