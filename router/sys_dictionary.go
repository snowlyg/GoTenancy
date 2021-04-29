package router

import (
	"github.com/kataras/iris/v12"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
	"github.com/snowlyg/go-tenancy/middleware"
)

func InitSysDictionaryRouter(Router iris.Party) {
	SysDictionaryRouter := Router.Party("/sysDictionary", middleware.OperationRecord())
	{
		SysDictionaryRouter.Post("/createSysDictionary", v1.CreateSysDictionary)   // 新建SysDictionary
		SysDictionaryRouter.Delete("/deleteSysDictionary", v1.DeleteSysDictionary) // 删除SysDictionary
		SysDictionaryRouter.Put("/updateSysDictionary", v1.UpdateSysDictionary)    // 更新SysDictionary
		SysDictionaryRouter.Get("/findSysDictionary", v1.FindSysDictionary)        // 根据ID获取SysDictionary
		SysDictionaryRouter.Get("/getSysDictionaryList", v1.GetSysDictionaryList)  // 获取SysDictionary列表
	}
}
