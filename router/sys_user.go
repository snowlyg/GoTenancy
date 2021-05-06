package router

import (
	"github.com/kataras/iris/v12"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
	"github.com/snowlyg/go-tenancy/middleware"
)

func InitUserRouter(Router iris.Party) {
	UserRouter := Router.Party("/user", middleware.OperationRecord())
	{
		UserRouter.Get("/logout", v1.Logout)
		UserRouter.Get("/clean", v1.Clean)
		UserRouter.Post("/register", v1.Register)
		UserRouter.Post("/changePassword", v1.ChangePassword)     // 修改密码
		UserRouter.Post("/getUserList", v1.GetUserList)           // 分页获取用户列表
		UserRouter.Post("/setUserAuthority", v1.SetUserAuthority) // 设置用户权限
		UserRouter.Delete("/deleteUser", v1.DeleteUser)           // 删除用户
		UserRouter.Put("/setUserInfo", v1.SetUserInfo)            // 设置用户信息
	}
}
