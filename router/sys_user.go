package router

import (
	"github.com/kataras/iris/v12"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitUserRouter(Router iris.Party) {
	UserRouter := Router.Party("/user")
	{
		UserRouter.Get("/logout", v1.Logout)                         // 退出
		UserRouter.Get("/clean", v1.Clean)                           //清空授权
		UserRouter.Post("/register", v1.Register)                    // 注册
		UserRouter.Post("/changePassword", v1.ChangePassword)        // 修改密码
		UserRouter.Post("/getAdminList", v1.GetAdminList)            // 分页获取管理员列表
		UserRouter.Post("/getTenancyList", v1.GetTenancyList)        // 分页获取商户列表
		UserRouter.Post("/getGeneralList", v1.GetGeneralList)        // 分页获取普通用户列表
		UserRouter.Post("/setUserAuthority", v1.SetUserAuthority)    // 设置用户权限
		UserRouter.Delete("/deleteUser", v1.DeleteUser)              // 删除用户
		UserRouter.Put("/setUserInfo/{user_id:int}", v1.SetUserInfo) // 设置用户信息
	}
}
