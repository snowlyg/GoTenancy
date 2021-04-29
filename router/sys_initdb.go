package router

import (
	"github.com/kataras/iris/v12"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitInitRouter(Router iris.Party) {
	ApiRouter := Router.Party("init")
	{
		ApiRouter.Post("initdb", v1.InitDB)   // 创建Api
		ApiRouter.Post("checkdb", v1.CheckDB) // 创建Api
	}
}
