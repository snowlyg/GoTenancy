package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitProductRouter(Router *gin.RouterGroup) {
	ProductRouter := Router.Group("/product")
	{
		ProductRouter.GET("/getEditProductFictiMap/:id", v1.GetEditProductFictiMap)
		ProductRouter.GET("/getProductFilter", v1.GetProductFilter)
		ProductRouter.PUT("/setProductFicti/:id", v1.SetProductFicti)
		ProductRouter.POST("/createProduct", v1.CreateProduct)
		ProductRouter.POST("/changeProductStatus", v1.ChangeProductStatus)
		ProductRouter.POST("/getProductList", v1.GetProductList)
		ProductRouter.GET("/getProductById/:id", v1.GetProductById)
		ProductRouter.PUT("/updateProduct/:id", v1.UpdateProduct)
		ProductRouter.DELETE("/deleteProduct/:id", v1.DeleteProduct)
	}
}
