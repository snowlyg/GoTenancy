package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitAttrTemplateRouter(Router *gin.RouterGroup) {
	AttrTemplateRouter := Router.Group("/attr_template")
	{
		AttrTemplateRouter.POST("/createAttrTemplate", v1.CreateAttrTemplate)
		AttrTemplateRouter.POST("/getAttrTemplateList", v1.GetAttrTemplateList)
		AttrTemplateRouter.POST("/getAttrTemplateById", v1.GetAttrTemplateById)
		AttrTemplateRouter.PUT("/updateAttrTemplate", v1.UpdateAttrTemplate)
		AttrTemplateRouter.DELETE("/deleteAttrTemplate", v1.DeleteAttrTemplate)
	}
}
