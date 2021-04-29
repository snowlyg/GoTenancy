package router

import (
	"github.com/kataras/iris/v12"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
	"github.com/snowlyg/go-tenancy/middleware"
)

func InitSysDictionaryDetailRouter(Router iris.Party) {
	SysDictionaryDetailRouter := Router.Party("/sysDictionaryDetail", middleware.OperationRecord())
	{
		SysDictionaryDetailRouter.Post("/createSysDictionaryDetail", v1.CreateSysDictionaryDetail)   // 新建SysDictionaryDetail
		SysDictionaryDetailRouter.Delete("/deleteSysDictionaryDetail", v1.DeleteSysDictionaryDetail) // 删除SysDictionaryDetail
		SysDictionaryDetailRouter.Put("/updateSysDictionaryDetail", v1.UpdateSysDictionaryDetail)    // 更新SysDictionaryDetail
		SysDictionaryDetailRouter.Get("/findSysDictionaryDetail", v1.FindSysDictionaryDetail)        // 根据ID获取SysDictionaryDetail
		SysDictionaryDetailRouter.Get("/getSysDictionaryDetailList", v1.GetSysDictionaryDetailList)  // 获取SysDictionaryDetail列表
	}
}
