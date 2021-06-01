package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitProductRouter(Router *gin.RouterGroup) {
	ProductRouter := Router.Group("/product")
	{
		ProductRouter.POST("/createProduct", v1.CreateProduct)
		ProductRouter.POST("/getProductList", v1.GetProductList)
		ProductRouter.POST("/getProductById", v1.GetProductById)
		ProductRouter.PUT("/updateProduct", v1.UpdateProduct)
		ProductRouter.DELETE("/deleteProduct", v1.DeleteProduct)
	}
}
