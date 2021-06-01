package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/response"
)

// NeedInit
func NeedInit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if g.TENANCY_DB == nil {
			response.OkWithDetailed(gin.H{
				"needInit": true,
			}, "前往初始化数据库", ctx)
			ctx.Abort()
		} else {
			ctx.Next()
		}
		// 处理请求
	}
}
