package middleware

import (
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
)

// 拦截器
func CasbinHandler() iris.Handler {
	return func(ctx iris.Context) {
		claims := ctx.Values().Get("claims")
		waitUse := claims.(*request.CustomClaims)
		// 获取请求的URI
		obj := ctx.FullRequestURI()
		// 获取请求方法
		act := ctx.Method()
		// 获取用户的角色
		sub := waitUse.AuthorityId
		e := service.Casbin()
		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if g.TENANCY_CONFIG.System.Env == "develop" || success {
			ctx.Next()
		} else {
			response.FailWithDetailed(iris.Map{}, "权限不足", ctx)
			ctx.StatusCode(http.StatusForbidden)
			return
		}
	}
}
