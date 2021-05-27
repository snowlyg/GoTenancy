package router

import (
	"github.com/kataras/iris/v12"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitCategoryRouter(Router iris.Party) {
	CategoryRouter := Router.Party("/category")
	{
		CategoryRouter.Post("/createCategory", v1.CreateCategory)
		CategoryRouter.Post("/getCategoryList", v1.GetCategoryList)
		CategoryRouter.Post("/getCategoryById", v1.GetCategoryById)
		CategoryRouter.Put("/updateCategory", v1.UpdateCategory)
		CategoryRouter.Delete("/deleteCategory", v1.DeleteCategory)
	}
}
