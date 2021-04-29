package router

import (
	"github.com/kataras/iris/v12"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
	"github.com/snowlyg/go-tenancy/middleware"
)

func InitEmailRouter(Router iris.Party) {
	UserRouter := Router.Party("email", middleware.OperationRecord())
	{
		UserRouter.Post("emailTest", v1.EmailTest) // 发送测试邮件
	}
}
