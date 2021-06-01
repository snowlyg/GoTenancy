package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitBrandRouter(Router *gin.RouterGroup) {
	BrandRouter := Router.Group("/brand")
	{
		BrandRouter.POST("/createBrand", v1.CreateBrand)
		BrandRouter.POST("/getBrandList", v1.GetBrandList)
		BrandRouter.POST("/getBrandById", v1.GetBrandById)
		BrandRouter.POST("/setBrandCate", v1.SetBrandCate)
		BrandRouter.PUT("/updateBrand", v1.UpdateBrand)
		BrandRouter.DELETE("/deleteBrand", v1.DeleteBrand)
	}
}
