package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitInitRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("/init")
	{
		ApiRouter.POST("/initdb", v1.InitDB)  // 创建Api
		ApiRouter.GET("/checkdb", v1.CheckDB) // 创建Api
	}
}
