package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"react-blog-server/apps/app/article/internal/logic"
	"react-blog-server/apps/app/article/internal/svc"
)

func ArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewArticleLogic(r.Context(), svcCtx)
		resp, err := l.Article()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
