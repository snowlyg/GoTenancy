package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitConfigRouter(Router *gin.RouterGroup) {
	ConfigRouter := Router.Group("/config")
	{
		ConfigRouter.POST("/createConfig", v1.CreateConfig)
		ConfigRouter.POST("/getConfigList", v1.GetConfigList)
		ConfigRouter.POST("/getConfigByName", v1.GetConfigByName)
		ConfigRouter.PUT("/updateConfig", v1.UpdateConfig)
		ConfigRouter.DELETE("/deleteConfig", v1.DeleteConfig)
	}
}
