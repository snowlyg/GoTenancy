package general

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitAddressRouter(Router *gin.RouterGroup) {
	AddressRouter := Router.Group("/address")
	{
		AddressRouter.POST("/createAddress", v1.CreateAddress)
		AddressRouter.POST("/getAddressList", v1.GetAddressList)
		AddressRouter.POST("/getAddressById", v1.GetAddressById)
		AddressRouter.PUT("/updateAddress", v1.UpdateAddress)
		AddressRouter.DELETE("/deleteAddress", v1.DeleteAddress)
	}
}

func InitReceiptRouter(Router *gin.RouterGroup) {
	ReceiptRouter := Router.Group("/receipt")
	{
		ReceiptRouter.POST("/createReceipt", v1.CreateReceipt)
		ReceiptRouter.POST("/getReceiptList", v1.GetReceiptList)
		ReceiptRouter.POST("/getReceiptById", v1.GetReceiptById)
		ReceiptRouter.PUT("/updateReceipt", v1.UpdateReceipt)
		ReceiptRouter.DELETE("/deleteReceipt", v1.DeleteReceipt)
	}
}
