package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitTenancyRouter(Router *gin.RouterGroup) {
	TenancyRouter := Router.Group("/tenancy")
	{
		TenancyRouter.GET("/getTenancies/{code:int}", v1.GetTenanciesByRegion) // 获取Tenancy列表
		TenancyRouter.POST("/createTenancy", v1.CreateTenancy)                 // 创建Tenancy
		TenancyRouter.POST("/getTenancyList", v1.GetTenanciesList)             // 获取Tenancy列表
		TenancyRouter.POST("/getTenancyById", v1.GetTenancyById)               // 获取单条Tenancy消息
		TenancyRouter.POST("/setTenancyRegion", v1.SetTenancyRegion)           // 设置商户地区
		TenancyRouter.PUT("/updateTenancy", v1.UpdateTenancy)                  // 更新Tenancy
		TenancyRouter.DELETE("/deleteTenancy", v1.DeleteTenancy)               // 删除Tenancy
	}
}
