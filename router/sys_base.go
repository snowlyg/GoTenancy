package router

import (
	"github.com/kataras/iris/v12"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
	"github.com/snowlyg/go-tenancy/middleware"
)

func InitBaseRouter(Router iris.Party) (R iris.Party) {
	BaseRouter := Router.Party("base", middleware.NeedInit())
	{
		BaseRouter.Post("login", v1.Login)
		BaseRouter.Post("captcha", v1.Captcha)
	}
	return BaseRouter
}
