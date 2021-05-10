package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/response"
)

// NeedInit
func NeedInit() iris.Handler {
	return func(ctx iris.Context) {
		if g.TENANCY_DB == nil {
			response.OkWithDetailed(iris.Map{
				"needInit": true,
			}, "前往初始化数据库", ctx)
			ctx.StopExecution()
		} else {
			ctx.Next()
		}
		// 处理请求
	}
}
