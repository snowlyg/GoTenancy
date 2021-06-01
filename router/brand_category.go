package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitBrandCategoryRouter(Router *gin.RouterGroup) {
	BrandCategoryRouter := Router.Group("/brandCategory")
	{
		BrandCategoryRouter.POST("/createBrandCategory", v1.CreateBrandCategory)
		BrandCategoryRouter.POST("/getBrandCategoryList", v1.GetBrandCategoryList)
		BrandCategoryRouter.POST("/getBrandCategoryById", v1.GetBrandCategoryById)
		BrandCategoryRouter.PUT("/updateBrandCategory", v1.UpdateBrandCategory)
		BrandCategoryRouter.DELETE("/deleteBrandCategory", v1.DeleteBrandCategory)
	}
}
