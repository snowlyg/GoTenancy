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

//LoginTenancy 后台登录
func LoginTenancy(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if loginResponse, err := service.LoginTenancy(req.Id); err != nil {
		g.TENANCY_LOG.Error("登陆失败!", zap.Any("err", err))
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.OkWithDetailed(loginResponse, "登录成功", ctx)
	}
}

// CreateTenancy
func CreateTenancy(ctx *gin.Context) {
	var tenancy model.SysTenancy
	if errs := ctx.ShouldBindJSON(&tenancy); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if returnTenancy, err := service.CreateTenancy(tenancy); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("添加失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(returnTenancy, "创建成功", ctx)
	}
}

// GetTenancyById
func GetTenancyById(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	tenancy, err := service.GetTenancyByID(req.Id)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
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
	err := service.SetTenancyRegionByID(regionCode)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("设置成功", ctx)
	}
}

// ChangeTenancyStatus
func ChangeTenancyStatus(ctx *gin.Context) {
	var changeStatus request.ChangeStatus
	if errs := ctx.ShouldBindJSON(&changeStatus); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	err := service.ChangeTenancyStatus(changeStatus)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("设置成功", ctx)
	}
}

// UpdateTenancy
func UpdateTenancy(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	var tenancy model.SysTenancy
	if errs := ctx.ShouldBindJSON(&tenancy); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if returnTenancy, err := service.UpdateTenancy(tenancy, req.Id); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(returnTenancy, "更新成功", ctx)
	}
}

// DeleteTenancy
func DeleteTenancy(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.DeleteTenancy(req.Id); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}

// GetTenanciesList 分页获取商户列表
func GetTenanciesList(ctx *gin.Context) {
	var pageInfo request.TenancyPageInfo
	if errs := ctx.ShouldBindJSON(&pageInfo); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if list, total, err := service.GetTenanciesInfoList(pageInfo); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
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
	if tenancies, err := service.GetTenanciesByRegion(ctx.Param("code")); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(tenancies, "获取成功", ctx)
	}
}

// GetTenancyCount 获取Tenancy对应状态数量
func GetTenancyCount(ctx *gin.Context) {
	if tenancies, err := service.GetTenancyCount(); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(tenancies, "获取成功", ctx)
	}
}

// GetTenancyInfo 获取Tenancy对应状态数量
func GetTenancyInfo(ctx *gin.Context) {
	if tenancies, err := service.GetTenancyInfo(ctx); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(tenancies, "获取成功", ctx)
	}
}
