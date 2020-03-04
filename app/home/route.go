package home

import (
	"github.com/betacraft/yaag/irisyaag"
	"github.com/kataras/iris/v12"
	"go-tenancy/app/home/homecontrollers"
	"go-tenancy/app/home/homemiddlware"
	"go-tenancy/config/db"
)

func Register(api *iris.Application) {
	app := api.Party("/", homemiddlware.CrsAuth()).AllowMethods(iris.MethodOptions)
	{
		app.HandleDir("/static", "resources/app/static")
		app.Get("/", func(ctx iris.Context) { // 首页模块
			_ = ctx.View("app/index.html")
		})

		v1 := app.Party("/v1")
		{
			v1.Post("/admin/login", homecontrollers.UserLogin)
			v1.PartyFunc("/admin", func(app iris.Party) {
				v1.Use(irisyaag.New())
				casbinMiddleware := homemiddlware.New(db.GetCasbinEnforcer())         //casbin for gorm                                                   // <- IMPORTANT, register the middleware.
				app.Use(homemiddlware.JwtHandler().Serve, casbinMiddleware.ServeHTTP) //登录验证
				app.Get("/logout", homecontrollers.UserLogout).Name = "退出"

				app.PartyFunc("/users", func(users iris.Party) {
					users.Get("/", homecontrollers.GetAllUsers).Name = "用户列表"
					users.Get("/{id:uint}", homecontrollers.GetUser).Name = "用户详情"
					users.Post("/", homecontrollers.CreateUser).Name = "创建用户"
					users.Put("/{id:uint}", homecontrollers.UpdateUser).Name = "编辑用户"
					users.Delete("/{id:uint}", homecontrollers.DeleteUser).Name = "删除用户"
					users.Get("/profile", homecontrollers.GetProfile).Name = "个人信息"
				})
				app.PartyFunc("/roles", func(roles iris.Party) {
					roles.Get("/", homecontrollers.GetAllRoles).Name = "角色列表"
					roles.Get("/{id:uint}", homecontrollers.GetRole).Name = "角色详情"
					roles.Post("/", homecontrollers.CreateRole).Name = "创建角色"
					roles.Put("/{id:uint}", homecontrollers.UpdateRole).Name = "编辑角色"
					roles.Delete("/{id:uint}", homecontrollers.DeleteRole).Name = "删除角色"
				})
				app.PartyFunc("/permissions", func(permissions iris.Party) {
					permissions.Get("/", homecontrollers.GetAllPermissions).Name = "权限列表"
					permissions.Get("/{id:uint}", homecontrollers.GetPermission).Name = "权限详情"
					//permissions.Post("/", controllers.CreatePermission).Name = "创建权限"
					//permissions.Put("/{id:uint}", controllers.UpdatePermission).Name = "编辑权限"
					//permissions.Delete("/{id:uint}", controllers.DeletePermission).Name = "删除权限"
				})
			})
		}
	}
}
