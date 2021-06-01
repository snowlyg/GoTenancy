package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/multi"
)

// IsAdmin
func IsAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !multi.IsAdmin(ctx) && !multi.IsTenancy(ctx) {
			response.ForbiddenFailWithMessage("无此操作权限", ctx)
			ctx.Abort()
		}
		ctx.Next()
	}
}
