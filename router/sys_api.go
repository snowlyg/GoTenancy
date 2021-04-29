package router

import (
	"github.com/kataras/iris/v12"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
	"github.com/snowlyg/go-tenancy/middleware"
)

func InitApiRouter(Router iris.Party) {
	ApiRouter := Router.Party("/api", middleware.OperationRecord())
	{
		ApiRouter.Post("/createApi", v1.CreateApi)               // 创建Api
		ApiRouter.Post("/deleteApi", v1.DeleteApi)               // 删除Api
		ApiRouter.Post("/getApiList", v1.GetApiList)             // 获取Api列表
		ApiRouter.Post("/getApiById", v1.GetApiById)             // 获取单条Api消息
		ApiRouter.Post("/updateApi", v1.UpdateApi)               // 更新api
		ApiRouter.Post("/getAllApis", v1.GetAllApis)             // 获取所有api
		ApiRouter.Delete("/deleteApisByIds", v1.DeleteApisByIds) // 删除选中api
	}
}
