package router

import (
	"github.com/kataras/iris/v12"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitMiniRouter(Router iris.Party) {
	MiniRouter := Router.Party("/mini")
	{
		MiniRouter.Post("/createMini", v1.CreateMini)
		MiniRouter.Post("/getMiniList", v1.GetMiniList)
		MiniRouter.Post("/getMiniById", v1.GetMiniById)
		MiniRouter.Put("/updateMini", v1.UpdateMini)
		MiniRouter.Delete("/deleteMini", v1.DeleteMini)
	}
}
