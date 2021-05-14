package router

import (
	"github.com/kataras/iris/v12"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
	"github.com/snowlyg/go-tenancy/middleware"
)

func InitPublicRouter(Router iris.Party) (R iris.Party) {
	BaseRouter := Router.Party("/public", middleware.NeedInit())
	{
		BaseRouter.Post("/login", v1.Login)
		BaseRouter.Post("/captcha", v1.Captcha)
		BaseRouter.Get("/region/{p_code:int}", v1.Region)
		BaseRouter.Get("/getRegionList", v1.RegionList)
	}
	return BaseRouter
}
