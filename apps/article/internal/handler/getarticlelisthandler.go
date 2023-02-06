package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"react-blog-server/apps/article/internal/logic"
	"react-blog-server/apps/article/internal/svc"
)

func getArticleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetArticleListLogic(r.Context(), svcCtx)
		resp, err := l.GetArticleList()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
