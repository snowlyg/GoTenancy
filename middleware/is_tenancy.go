package middleware

import (
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/multi"
)

// IsTenancy
func IsTenancy() iris.Handler {
	return func(ctx iris.Context) {
		if !multi.IsTenancy(ctx) {
			response.FailWithMessage("无此操作权限", ctx)
			ctx.StatusCode(http.StatusForbidden)
			return
		}
		ctx.Next()
	}
}
