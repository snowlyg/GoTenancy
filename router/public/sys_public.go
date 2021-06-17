package public

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
	"github.com/snowlyg/go-tenancy/middleware"
)

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
