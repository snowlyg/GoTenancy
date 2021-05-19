package router

import (
	"github.com/kataras/iris/v12"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitAddressRouter(Router iris.Party) {
	AddressRouter := Router.Party("/address")
	{
		AddressRouter.Post("/createAddress", v1.CreateAddress)
		AddressRouter.Post("/getAddressList", v1.GetAddressList)
		AddressRouter.Post("/getAddressById", v1.GetAddressById)
		AddressRouter.Put("/updateAddress", v1.UpdateAddress)
		AddressRouter.Delete("/deleteAddress", v1.DeleteAddress)
	}
}
