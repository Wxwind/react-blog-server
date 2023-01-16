package handler

import (
	"net/http"
	"react-blog-server/apps/article/internal/logic"
	"react-blog-server/apps/article/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
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
