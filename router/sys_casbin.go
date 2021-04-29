package router

import (
	"github.com/kataras/iris/v12"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
	"github.com/snowlyg/go-tenancy/middleware"
)

func InitCasbinRouter(Router iris.Party) {
	CasbinRouter := Router.Party("/casbin", middleware.OperationRecord())
	{
		CasbinRouter.Post("/updateCasbin", v1.UpdateCasbin)
		CasbinRouter.Post("/getPolicyPathByAuthorityId", v1.GetPolicyPathByAuthorityId)
	}
}
