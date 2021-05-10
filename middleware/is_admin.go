package middleware

import (
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/multi"
)

// IsAdmin
func IsAdmin() iris.Handler {
	return func(ctx iris.Context) {
		if !multi.IsAdmin(ctx) {
			response.FailWithMessage("无此操作权限", ctx)
			ctx.StatusCode(http.StatusForbidden)
			return
		}
		ctx.Next()
	}
}
