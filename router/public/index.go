package public

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
	"github.com/snowlyg/go-tenancy/middleware"
)

func InitAuthRouter(Router *gin.RouterGroup) {
	Router.GET("/logout", v1.Logout) // 退出
	Router.GET("/clean", v1.Clean)   //清空授权
}

func InitInitRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("/init")
	{
		ApiRouter.POST("/initdb", v1.InitDB)  // 创建Api
		ApiRouter.GET("/checkdb", v1.CheckDB) // 创建Api
	}
}
func InitPublicRouter(Router *gin.RouterGroup) (R *gin.RouterGroup) {
	BaseRouter := Router.Group("/public", middleware.NeedInit())
	{
		BaseRouter.POST("/admin/login", v1.AdminLogin)
		BaseRouter.POST("/client/login", v1.ClientLogin)
		BaseRouter.GET("/captcha", v1.Captcha)
		BaseRouter.GET("/region/:p_code", v1.Region)
		BaseRouter.GET("/getRegionList", v1.RegionList)
	}
	return BaseRouter
}
