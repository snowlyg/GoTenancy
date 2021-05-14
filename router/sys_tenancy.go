package router

import (
	"github.com/kataras/iris/v12"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitTenancyRouter(Router iris.Party) {
	TenancyRouter := Router.Party("/tenancy")
	{
		TenancyRouter.Get("/getTenancies/{code:int}", v1.GetTenanciesByRegion) // 获取Tenancy列表
		TenancyRouter.Post("/createTenancy", v1.CreateTenancy)                 // 创建Tenancy
		TenancyRouter.Post("/getTenancyList", v1.GetTenanciesList)             // 获取Tenancy列表
		TenancyRouter.Post("/getTenancyById", v1.GetTenancyById)               // 获取单条Tenancy消息
		TenancyRouter.Post("/setTenancyRegion", v1.SetTenancyRegion)           // 设置商户地区
		TenancyRouter.Put("/updateTenancy", v1.UpdateTenancy)                  // 更新Tenancy
		TenancyRouter.Delete("/deleteTenancy", v1.DeleteTenancy)               // 删除Tenancy
	}
}
