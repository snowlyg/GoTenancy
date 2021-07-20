package client

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

// UpdateClientTenancy
func UpdateClientTenancy(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	var tenancy request.UpdateClientTenancy
	if errs := ctx.ShouldBindJSON(&tenancy); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.UpdateClientTenancy(tenancy, req.Id); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("更新成功", ctx)
	}
}

// GetTenancyCopyCount
func GetTenancyCopyCount(ctx *gin.Context) {
	if copyProductNum, err := service.GetTenancyCopyCount(multi.GetTenancyId(ctx)); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(gin.H{
			"count": copyProductNum,
		}, "获取成功", ctx)
	}
}

// GetTenancyInfo 获取Tenancy对应状态数量
func GetTenancyInfo(ctx *gin.Context) {
	if tenancies, err := service.GetTenancyInfo(multi.GetTenancyId(ctx)); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(tenancies, "获取成功", ctx)
	}
}

// GetUpdateTenancyMap 获取Tenancy对应状态数量
func GetUpdateTenancyMap(ctx *gin.Context) {
	if tenancies, err := service.GetUpdateTenancyMap(ctx); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(tenancies, "获取成功", ctx)
	}
}
