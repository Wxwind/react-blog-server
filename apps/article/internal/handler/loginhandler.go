package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"react-blog-server/apps/article/internal/logic"
	"react-blog-server/apps/article/internal/svc"
	"react-blog-server/apps/article/internal/types"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
