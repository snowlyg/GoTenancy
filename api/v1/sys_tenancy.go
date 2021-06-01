package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"go.uber.org/zap"
)

// CreateTenancy
func CreateTenancy(ctx *gin.Context) {
	var tenancy request.CreateSysTenancy
	if errs := ctx.ShouldBindJSON(&tenancy); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if returnTenancy, err := service.CreateTenancy(tenancy); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", ctx)
	} else {
		response.OkWithDetailed(getTenancyMap(returnTenancy), "创建成功", ctx)
	}
}

// GetTenancyById
func GetTenancyById(ctx *gin.Context) {
	var reqId request.GetById
	if errs := ctx.ShouldBindJSON(&reqId); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	tenancy, err := service.GetTenancyByID(reqId.Id)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithData(tenancy, ctx)
	}
}

// SetTenancyRegion
func SetTenancyRegion(ctx *gin.Context) {
	var regionCode request.SetRegionCode
	if errs := ctx.ShouldBindJSON(&regionCode); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	err := service.SetTenancyRegionByID(regionCode.Id, regionCode.SysRegionCode)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithMessage("设置成功", ctx)
	}
}

// UpdateTenancy
func UpdateTenancy(ctx *gin.Context) {
	var tenancy request.UpdateSysTenancy
	if errs := ctx.ShouldBindJSON(&tenancy); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if returnTenancy, err := service.UpdateTenany(tenancy); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", ctx)
	} else {
		response.OkWithDetailed(getTenancyMap(returnTenancy), "更新成功", ctx)
	}
}

// getTenancyMap
func getTenancyMap(returnTenancy model.SysTenancy) gin.H {
	return gin.H{"id": returnTenancy.ID, "uuid": returnTenancy.UUID, "name": returnTenancy.Name, "tele": returnTenancy.Tele, "address": returnTenancy.Address, "businessTime": returnTenancy.BusinessTime, "sysRegionCode": returnTenancy.SysRegionCode}
}

// DeleteTenancy
func DeleteTenancy(ctx *gin.Context) {
	var reqId request.GetById
	if errs := ctx.ShouldBindJSON(&reqId); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.DeleteTenancy(reqId.Id); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}

// GetTenanciesList 分页获取商户列表
func GetTenanciesList(ctx *gin.Context) {
	var pageInfo request.PageInfo
	if errs := ctx.ShouldBindJSON(&pageInfo); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if list, total, err := service.GetTenanciesInfoList(pageInfo); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}

// GetTenanciesByRegion 根据区域获取商户列表，不分页
func GetTenanciesByRegion(ctx *gin.Context) {
	code := ctx.DefaultQuery("code", "-1")
	if tenancies, err := service.GetTenanciesByRegion(code); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(tenancies, "获取成功", ctx)
	}
}
