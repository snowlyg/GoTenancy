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
