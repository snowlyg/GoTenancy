package public

import (
	"github.com/gin-gonic/gin"
	public "github.com/snowlyg/go-tenancy/api/v1/public"
	"github.com/snowlyg/go-tenancy/middleware"
)

func InitAuthRouter(Router *gin.RouterGroup) {
	Router.GET("/logout", public.Logout) // 退出
	Router.GET("/clean", public.Clean)   //清空授权
}

func InitInitRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("/init")
	{
		ApiRouter.POST("/initdb", public.InitDB)  // 创建Api
		ApiRouter.GET("/checkdb", public.CheckDB) // 创建Api
	}
}
func InitPublicRouter(Router *gin.RouterGroup) (R *gin.RouterGroup) {
	BaseRouter := Router.Group("/public", middleware.NeedInit())
	{
		BaseRouter.POST("/admin/login", public.AdminLogin)
		BaseRouter.POST("/merchant/login", public.ClientLogin)
		BaseRouter.GET("/captcha", public.Captcha)
		BaseRouter.GET("/region/:p_code", public.Region)
		BaseRouter.GET("/getRegionList", public.RegionList)
	}
	return BaseRouter
}
