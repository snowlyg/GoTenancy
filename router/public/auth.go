package public

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitAuthRouter(Router *gin.RouterGroup) {
	Router.GET("/logout", v1.Logout) // 退出
	Router.GET("/clean", v1.Clean)   //清空授权
}
