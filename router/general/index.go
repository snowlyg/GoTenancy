package general

import (
	"github.com/gin-gonic/gin"
	general "github.com/snowlyg/go-tenancy/api/v1/general"
)

func InitAddressRouter(Router *gin.RouterGroup) {
	AddressRouter := Router.Group("/address")
	{
		AddressRouter.POST("/createAddress", general.CreateAddress)
		AddressRouter.POST("/getAddressList", general.GetAddressList)
		AddressRouter.POST("/getAddressById", general.GetAddressById)
		AddressRouter.PUT("/updateAddress", general.UpdateAddress)
		AddressRouter.DELETE("/deleteAddress", general.DeleteAddress)
	}
}

func InitReceiptRouter(Router *gin.RouterGroup) {
	ReceiptRouter := Router.Group("/receipt")
	{
		ReceiptRouter.POST("/createReceipt", general.CreateReceipt)
		ReceiptRouter.POST("/getReceiptList", general.GetReceiptList)
		ReceiptRouter.POST("/getReceiptById", general.GetReceiptById)
		ReceiptRouter.PUT("/updateReceipt", general.UpdateReceipt)
		ReceiptRouter.DELETE("/deleteReceipt", general.DeleteReceipt)
	}
}
