package router

import (
	"github.com/kataras/iris/v12"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitCasbinRouter(Router iris.Party) {
	CasbinRouter := Router.Party("/casbin")
	{
		CasbinRouter.Post("/updateCasbin", v1.UpdateCasbin)
		CasbinRouter.Post("/getPolicyPathByAuthorityId", v1.GetPolicyPathByAuthorityId)
	}
}
