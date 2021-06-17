package admin

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("/user")
	{
		UserRouter.POST("/registerAdmin", v1.RegisterAdmin)       // 注册
		UserRouter.POST("/registerTenancy", v1.RegisterTenancy)   // 注册
		UserRouter.POST("/changePassword", v1.ChangePassword)     // 修改密码
		UserRouter.POST("/getAdminList", v1.GetAdminList)         // 分页获取管理员列表
		UserRouter.POST("/getTenancyList", v1.GetTenancyList)     // 分页获取商户列表
		UserRouter.POST("/getGeneralList", v1.GetGeneralList)     // 分页获取普通用户列表
		UserRouter.POST("/setUserAuthority", v1.SetUserAuthority) // 设置用户权限
		UserRouter.DELETE("/deleteUser", v1.DeleteUser)           // 删除用户
		UserRouter.PUT("/setUserInfo/:user_id", v1.SetUserInfo)   // 设置用户信息
	}
}
