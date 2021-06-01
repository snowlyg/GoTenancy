package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitCategoryRouter(Router *gin.RouterGroup) {
	CategoryRouter := Router.Group("/category")
	{
		CategoryRouter.POST("/createCategory", v1.CreateCategory)
		CategoryRouter.POST("/getCategoryList", v1.GetCategoryList)
		CategoryRouter.POST("/getCategoryById", v1.GetCategoryById)
		CategoryRouter.PUT("/updateCategory", v1.UpdateCategory)
		CategoryRouter.DELETE("/deleteCategory", v1.DeleteCategory)
	}
}
