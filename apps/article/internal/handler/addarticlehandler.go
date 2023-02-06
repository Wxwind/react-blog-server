package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"react-blog-server/apps/article/internal/logic"
	"react-blog-server/apps/article/internal/svc"
	"react-blog-server/apps/article/internal/types"
)

func addArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddArticleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddArticleLogic(r.Context(), svcCtx)
		resp, err := l.AddArticle(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
