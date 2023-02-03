package middleware

import (
	"net/http"
	"react-blog-server/common/utils"
)

var allowOrigins = []string{
	"https://www.wxwind.top",
}

func Cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		method := r.Method

		//是cors请求
		origin := r.Header.Get("Origin")
		if origin != "" && utils.IsContain(allowOrigins, origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
			//允许浏览器发送的头
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			//允许浏览器拿到的头
			w.Header().Set("Access-Control-Expose-Headers", "")
			//是否允许cookies, authorization headers 或 TLS client certificates
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("content-type", "application/json")
		}

		//是非简单请求的预检请求，直接返回204，不做后续处理
		if method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next(w, r)
	}

}
