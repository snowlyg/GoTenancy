package router

import (
	"github.com/kataras/iris/v12"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitBrandCategoryRouter(Router iris.Party) {
	BrandCategoryRouter := Router.Party("/brandCategory")
	{
		BrandCategoryRouter.Post("/createBrandCategory", v1.CreateBrandCategory)
		BrandCategoryRouter.Post("/getBrandCategoryList", v1.GetBrandCategoryList)
		BrandCategoryRouter.Post("/getBrandCategoryById", v1.GetBrandCategoryById)
		BrandCategoryRouter.Put("/updateBrandCategory", v1.UpdateBrandCategory)
		BrandCategoryRouter.Delete("/deleteBrandCategory", v1.DeleteBrandCategory)
	}
}
