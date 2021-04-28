package middleware

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

// 处理跨域请求,支持options访问
func Cors() iris.Handler {
	return func(ctx iris.Context) {
		method := ctx.Method()
		origin := ctx.GetHeader("Origin")
		ctx.Header("Access-Control-Allow-Origin", origin)
		ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			ctx.StatusCode(http.StatusNoContent)
		}
		// 处理请求
		ctx.Next()
	}
}
