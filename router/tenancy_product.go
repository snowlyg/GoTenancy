package router

import (
	"github.com/kataras/iris/v12"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitProductRouter(Router iris.Party) {
	ProductRouter := Router.Party("/product")
	{
		ProductRouter.Post("/createProduct", v1.CreateProduct)
		ProductRouter.Post("/getProductList", v1.GetProductList)
		ProductRouter.Post("/getProductById", v1.GetProductById)
		ProductRouter.Put("/updateProduct", v1.UpdateProduct)
		ProductRouter.Delete("/deleteProduct", v1.DeleteProduct)
	}
}
