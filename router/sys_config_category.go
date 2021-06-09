package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitConfigCategoryRouter(Router *gin.RouterGroup) {
	ConfigCategoryRouter := Router.Group("/configCategory")
	{
		ConfigCategoryRouter.GET("/getCreateConfigCategoryMap", v1.GetCreateConfigCategoryMap)
		ConfigCategoryRouter.GET("/getUpdateConfigCategoryMap/:id", v1.GetUpdateConfigCategoryMap)
		ConfigCategoryRouter.GET("/getConfigCategoryList", v1.GetConfigCategoryList)
		ConfigCategoryRouter.POST("/createConfigCategory", v1.CreateConfigCategory)
		ConfigCategoryRouter.GET("/getConfigCategoryById/:id", v1.GetConfigCategoryById)
		ConfigCategoryRouter.PUT("/updateConfigCategory/:id", v1.UpdateConfigCategory)
		ConfigCategoryRouter.POST("/changeConfigCategoryStatus", v1.ChangeConfigCategoryStatus)
		ConfigCategoryRouter.DELETE("/deleteConfigCategory/:id", v1.DeleteConfigCategory)
	}
}
