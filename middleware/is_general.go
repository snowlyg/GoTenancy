package middleware

import (
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/multi"
)

// IsGeneral
func IsGeneral() iris.Handler {
	return func(ctx iris.Context) {
		if !multi.IsGeneral(ctx) {
			response.FailWithMessage("无此操作权限", ctx)
			ctx.StatusCode(http.StatusForbidden)
			return
		}
		ctx.Next()
	}
}
