package router

import (
	"github.com/kataras/iris/v12"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitBrandRouter(Router iris.Party) {
	BrandRouter := Router.Party("/brand")
	{
		BrandRouter.Post("/createBrand", v1.CreateBrand)
		BrandRouter.Post("/getBrandList", v1.GetBrandList)
		BrandRouter.Post("/getBrandById", v1.GetBrandById)
		BrandRouter.Post("/setBrandCate", v1.SetBrandCate)
		BrandRouter.Put("/updateBrand", v1.UpdateBrand)
		BrandRouter.Delete("/deleteBrand", v1.DeleteBrand)
	}
}
