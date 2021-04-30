package middleware

import (
	"fmt"
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
)

// 拦截器
func CasbinHandler() iris.Handler {
	return func(ctx iris.Context) {
		waitUse := jwt.Get(ctx).(*request.CustomClaims)
		// 获取请求的URI
		obj := ctx.FullRequestURI()
		// 获取请求方法
		act := ctx.Method()
		// 获取用户的角色
		sub := waitUse.AuthorityId
		e := service.Casbin()
		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		fmt.Printf("%s : %t", g.TENANCY_CONFIG.System.Env, success)
		if g.TENANCY_CONFIG.System.Env == "dev" || success {
			ctx.Next()
		} else {
			response.FailWithDetailed(iris.Map{}, "权限不足", ctx)
			ctx.StatusCode(http.StatusForbidden)
			return
		}
	}
}
