package admin

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitEmailRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("/email")
	{
		UserRouter.POST("/emailTest", v1.EmailTest) // 发送测试邮件
	}
}
