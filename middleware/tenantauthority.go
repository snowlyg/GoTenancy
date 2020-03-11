package middleware

import (
	"github.com/kataras/iris/v12"
	"go-tenancy/config/auth"
)

func TenantAuthorize(ctx iris.Context) {
	var currentUser interface{}
	//Get current user from request
	currentUser = auth.TenantAuthority.Auth.GetCurrentUser(ctx.Request())

	// 没有角色并且账号不是 nil，或者账号拥有角色
	if currentUser != nil {
		ctx.Next()
		return
	}

	// 返回权限未定义
	auth.TenantAuthority.AccessDeniedHandler(ctx.ResponseWriter(), ctx.Request())
}

func TenantAuthorizeloggedInHalfHour(ctx iris.Context) {
	var currentUser interface{}
	//Get current user from request
	currentUser = auth.TenantAuthority.Auth.GetCurrentUser(ctx.Request())

	// 没有角色并且账号不是 nil，或者账号拥有角色
	if auth.TenantAuthority.Role.HasRole(ctx.Request(), currentUser, "logged_in_half_hour") {
		ctx.Next()
		return
	}

	// 返回权限未定义
	auth.TenantAuthority.AccessDeniedHandler(ctx.ResponseWriter(), ctx.Request())
}
