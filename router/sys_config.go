package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitConfigRouter(Router *gin.RouterGroup) {
	ConfigRouter := Router.Group("/config")
	{
		ConfigRouter.GET("/getConfigMap/:category", v1.GetConfigMap)
		ConfigRouter.GET("/getCreateConfigMap", v1.GetCreateConfigMap)
		ConfigRouter.GET("/getUpdateConfigMap/:id", v1.GetUpdateConfigMap)
		ConfigRouter.POST("/getConfigList", v1.GetConfigList)
		ConfigRouter.POST("/createConfig", v1.CreateConfig)
		ConfigRouter.GET("/getConfigByKey/:key", v1.GetConfigByKey)
		ConfigRouter.GET("/getConfigByID/:id", v1.GetConfigByID)
		ConfigRouter.POST("/changeConfigStatus", v1.ChangeConfigStatus)
		ConfigRouter.PUT("/updateConfig/:id", v1.UpdateConfig)
		ConfigRouter.DELETE("/deleteConfig/:id", v1.DeleteConfig)
	}
}
