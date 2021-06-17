package admin

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitCasbinRouter(Router *gin.RouterGroup) {
	CasbinRouter := Router.Group("/casbin")
	{
		CasbinRouter.POST("/updateCasbin", v1.UpdateCasbin)
		CasbinRouter.POST("/getPolicyPathByAuthorityId", v1.GetPolicyPathByAuthorityId)
	}
}
