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

// CreateAutoLabel
func CreateAutoLabel(ctx *gin.Context) {
	var label request.LabelRule
	if errs := ctx.ShouldBindJSON(&label); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	label.SysTenancyID = multi.GetTenancyId(ctx)
	if returnUserLabel, err := service.CreateAutoLabel(label); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(returnUserLabel, "创建成功", ctx)
	}
}

// UpdateAutoLabel
func UpdateAutoLabel(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	var label request.LabelRule
	if errs := ctx.ShouldBindJSON(&label); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	label.SysTenancyID = multi.GetTenancyId(ctx)
	if err := service.UpdateAutoLabel(label, req.Id); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("更新成功", ctx)
	}
}

// GetAutoLabelList
func GetAutoLabelList(ctx *gin.Context) {
	var pageInfo request.UserLabelPageInfo
	if errs := ctx.ShouldBindJSON(&pageInfo); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if list, total, err := service.GetAutoUserLabelInfoList(pageInfo, multi.GetTenancyId(ctx)); err != nil {
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

// DeleteAutoLabel
func DeleteAutoLabel(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.DeleteAutoLabel(req.Id, multi.GetTenancyId(ctx)); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}
