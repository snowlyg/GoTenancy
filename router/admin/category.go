package admin

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitCategoryRouter(Router *gin.RouterGroup) {
	CategoryRouter := Router.Group("/category")
	{
		CategoryRouter.GET("/getCreateTenancyCategoryMap", v1.GetCreateTenancyCategoryMap)
		CategoryRouter.GET("/getUpdateTenancyCategoryMap/:id", v1.GetUpdateTenancyCategoryMap)
		CategoryRouter.GET("/getCategorySelect", v1.GetCategorySelect)
		CategoryRouter.POST("/createCategory", v1.CreateCategory)
		CategoryRouter.POST("/getCategoryList", v1.GetCategoryList)
		CategoryRouter.GET("/getCategoryById/:id", v1.GetCategoryById)
		CategoryRouter.POST("/changeCategoryStatus", v1.ChangeCategoryStatus)
		CategoryRouter.PUT("/updateCategory/:id", v1.UpdateCategory)
		CategoryRouter.DELETE("/deleteCategory/:id", v1.DeleteCategory)
	}
}
