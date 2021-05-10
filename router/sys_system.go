package router

import (
	"github.com/kataras/iris/v12"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitSystemRouter(Router iris.Party) {
	SystemRouter := Router.Party("/system")
	{
		SystemRouter.Post("/getSystemConfig", v1.GetSystemConfig) // 获取配置文件内容
		SystemRouter.Post("/setSystemConfig", v1.SetSystemConfig) // 设置配置文件内容
		SystemRouter.Post("/getServerInfo", v1.GetServerInfo)     // 获取服务器信息
		SystemRouter.Post("/reloadSystem", v1.ReloadSystem)       // 重启服务
	}
}
