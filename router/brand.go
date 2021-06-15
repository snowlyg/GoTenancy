package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitBrandRouter(Router *gin.RouterGroup) {
	BrandRouter := Router.Group("/brand")
	{
		BrandRouter.GET("/getCreateBrandMap", v1.GetCreateBrandMap)
		BrandRouter.GET("/getUpdateBrandMap/:id", v1.GetUpdateBrandMap)
		BrandRouter.POST("/createBrand", v1.CreateBrand)
		BrandRouter.POST("/getBrandList", v1.GetBrandList)
		BrandRouter.GET("/getBrandById/:id", v1.GetBrandById)
		BrandRouter.POST("/changeBrandStatus", v1.ChangeBrandStatus)
		BrandRouter.PUT("/updateBrand/:id", v1.UpdateBrand)
		BrandRouter.DELETE("/deleteBrand/:id", v1.DeleteBrand)
	}
}
