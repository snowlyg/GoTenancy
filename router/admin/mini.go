package admin

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitMiniRouter(Router *gin.RouterGroup) {
	MiniRouter := Router.Group("/mini")
	{
		MiniRouter.POST("/createMini", v1.CreateMini)
		MiniRouter.POST("/getMiniList", v1.GetMiniList)
		MiniRouter.POST("/getMiniById", v1.GetMiniById)
		MiniRouter.PUT("/updateMini", v1.UpdateMini)
		MiniRouter.DELETE("/deleteMini", v1.DeleteMini)
	}
}
