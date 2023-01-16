package handler

import (
	"net/http"
	"react-blog-server/apps/article/internal/logic"
	"react-blog-server/apps/article/internal/svc"
	"react-blog-server/apps/article/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func getArticlePathHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetArticlePathReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetArticlePathLogic(r.Context(), svcCtx)
		resp, err := l.GetArticlePath(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
