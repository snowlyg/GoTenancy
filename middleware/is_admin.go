package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/multi"
)

// IsAdmin
func IsAdmin() iris.Handler {
	return func(ctx iris.Context) {
		if !multi.IsAdmin(ctx) {
			response.ForbiddenFailWithMessage("无此操作权限", ctx)
			ctx.StopExecution()
		}
		ctx.Next()
	}
}
