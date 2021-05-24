package router

import (
	"github.com/kataras/iris/v12"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitAuthRouter(Router iris.Party) {
	Router.Get("/logout", v1.Logout) // 退出
	Router.Get("/clean", v1.Clean)   //清空授权
}
