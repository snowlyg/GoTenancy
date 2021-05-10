package middleware

import (
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

// 拦截器
func CasbinHandler() iris.Handler {
	return func(ctx iris.Context) {
		waitUse := multi.Get(ctx)
		if waitUse == nil {
			response.FailWithMessage("权限服务验证失败：token empty", ctx)
			ctx.StatusCode(http.StatusForbidden)
			return
		}
		obj := ctx.Path()          // 获取请求的URI
		act := ctx.Method()        // 获取请求方法
		sub := waitUse.AuthorityId // 获取用户的角色

		g.TENANCY_LOG.Debug("route path", zap.String("string", obj))

		// 判断策略中是否存在
		casbin, err := service.Casbin()
		if err != nil {
			g.TENANCY_LOG.Error("get casbin err", zap.Error(err))
			response.FailWithMessage("权限服务验证失败：casbin error", ctx)
			ctx.StatusCode(http.StatusForbidden)
			return
		}
		success, err := casbin.Enforce(sub, obj, act)
		if err != nil {
			response.FailWithMessage("权限服务验证失败：verfiy failed", ctx)
			ctx.StatusCode(http.StatusForbidden)
			return
		}
		if !success {
			response.FailWithMessage("无此操作权限", ctx)
			ctx.StatusCode(http.StatusForbidden)
			return
		}
		ctx.Next()
	}
}
