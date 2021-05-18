package router

import (
	"github.com/kataras/iris/v12"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitConfigRouter(Router iris.Party) {
	ConfigRouter := Router.Party("/config")
	{
		ConfigRouter.Post("/createConfig", v1.CreateConfig)
		ConfigRouter.Post("/getConfigList", v1.GetConfigList)
		ConfigRouter.Post("/getConfigByName", v1.GetConfigByName)
		ConfigRouter.Put("/updateConfig", v1.UpdateConfig)
		ConfigRouter.Delete("/deleteConfig", v1.DeleteConfig)
	}
}
