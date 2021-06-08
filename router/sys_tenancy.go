package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitTenancyRouter(Router *gin.RouterGroup) {
	TenancyRouter := Router.Group("/tenancy")
	{
		TenancyRouter.GET("/getTenancies/:code", v1.GetTenanciesByRegion)  // 获取Tenancy列表(不分页)
		TenancyRouter.GET("/getTenancyCount", v1.GetTenancyCount)          // 获取Tenancy对应状态数量
		TenancyRouter.POST("/createTenancy", v1.CreateTenancy)             // 创建Tenancy
		TenancyRouter.POST("/getTenancyList", v1.GetTenanciesList)         // 获取Tenancy列表
		TenancyRouter.GET("/getTenancyById/:id", v1.GetTenancyById)        // 获取单条Tenancy消息
		TenancyRouter.POST("/setTenancyRegion", v1.SetTenancyRegion)       // 设置商户地区
		TenancyRouter.POST("/changeTenancyStatus", v1.ChangeTenancyStatus) // 设置商户地区
		TenancyRouter.PUT("/updateTenancy/:id", v1.UpdateTenancy)          // 更新Tenancy
		TenancyRouter.DELETE("/deleteTenancy/:id", v1.DeleteTenancy)       // 删除Tenancy
	}
}
