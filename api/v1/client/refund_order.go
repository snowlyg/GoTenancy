package client

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"go.uber.org/zap"
)

// GetRefundOrderList
func GetRefundOrderList(ctx *gin.Context) {
	var pageInfo request.RefundOrderPageInfo
	if errs := ctx.ShouldBindJSON(&pageInfo); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if list, stat, total, err := service.GetRefundOrderInfoList(pageInfo, ctx); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(gin.H{
			"stat":     stat,
			"list":     list,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}

// GetRefundOrderRecord
func GetRefundOrderRecord(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	var pageInfo request.PageInfo
	if errs := ctx.ShouldBindJSON(&pageInfo); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if records, total, err := service.GetRefundOrderRecord(req.Id, pageInfo); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(gin.H{
			"list":     records,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}

func GetRefundOrderRemarkMap(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if chart, err := service.GetRefundOrderRemarkMap(req.Id, ctx); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithData(chart, ctx)
	}
}

// RemarkRefundOrder
func RemarkRefundOrder(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	var remark request.OrderRemark
	if errs := ctx.ShouldBindJSON(&remark); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.RemarkRefundOrder(req.Id, remark, ctx); err != nil {
		g.TENANCY_LOG.Error("操作失败!", zap.Any("err", err))
		response.FailWithMessage("操作失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("操作成功", ctx)
	}
}

// DeleteRefundOrder
func DeleteRefundOrder(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.DeleteRefundOrder(req.Id); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}
