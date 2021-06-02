package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitAuthorityRouter(Router *gin.RouterGroup) {
	AuthorityRouter := Router.Group("/authority")
	{
		AuthorityRouter.POST("/createAuthority", v1.CreateAuthority)                 // 创建角色
		AuthorityRouter.DELETE("/deleteAuthority", v1.DeleteAuthority)               // 删除角色
		AuthorityRouter.PUT("/updateAuthority", v1.UpdateAuthority)                  // 更新角色
		AuthorityRouter.POST("/copyAuthority", v1.CopyAuthority)                     // 更新角色
		AuthorityRouter.POST("/getAuthorityList", v1.GetAuthorityList)               // 获取角色列表
		AuthorityRouter.POST("/getAdminAuthorityList", v1.GetAdminAuthorityList)     // 获取员工角色列表
		AuthorityRouter.POST("/getTenancyAuthorityList", v1.GetTenancyAuthorityList) // 获取商户角色列表
		AuthorityRouter.POST("/getGeneralAuthorityList", v1.GetGeneralAuthorityList) // 获取普通用户角色列表
		AuthorityRouter.POST("/setDataAuthority", v1.SetDataAuthority)               // 设置角色资源权限
	}
}
