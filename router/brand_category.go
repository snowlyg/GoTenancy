package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitBrandCategoryRouter(Router *gin.RouterGroup) {
	BrandCategoryRouter := Router.Group("/brandCategory")
	{
		BrandCategoryRouter.GET("/getCreateBrandCategoryMap", v1.GetCreateBrandCategoryMap)
		BrandCategoryRouter.GET("/getUpdateBrandCategoryMap/:id", v1.GetUpdateBrandCategoryMap)
		BrandCategoryRouter.POST("/createBrandCategory", v1.CreateBrandCategory)
		BrandCategoryRouter.POST("/getBrandCategoryList", v1.GetBrandCategoryList)
		BrandCategoryRouter.GET("/getBrandCategoryById/:id", v1.GetBrandCategoryById)
		BrandCategoryRouter.POST("/changeBrandCategoryStatus", v1.ChangeBrandCategoryStatus)
		BrandCategoryRouter.PUT("/updateBrandCategory/:id", v1.UpdateBrandCategory)
		BrandCategoryRouter.DELETE("/deleteBrandCategory/:id", v1.DeleteBrandCategory)
	}
}
