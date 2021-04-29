package router

import (
	"github.com/kataras/iris/v12"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
	"github.com/snowlyg/go-tenancy/middleware"
)

func InitJwtRouter(Router iris.Party) {
	ApiRouter := Router.Party("jwt", middleware.OperationRecord())
	{
		ApiRouter.Post("jsonInBlacklist", v1.JsonInBlacklist) // jwt加入黑名单
	}
}
