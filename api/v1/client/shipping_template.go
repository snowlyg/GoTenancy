package client

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

// CreateShippingTemplate
func CreateShippingTemplate(ctx *gin.Context) {
	var product model.ShippingTemplate
	if errs := ctx.ShouldBindJSON(&product); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if returnShippingTemplate, err := service.CreateShippingTemplate(product, multi.GetTenancyId(ctx)); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("添加失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(returnShippingTemplate, "创建成功", ctx)
	}
}

// UpdateShippingTemplate
func UpdateShippingTemplate(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	var product request.UpdateShippingTemplate
	if errs := ctx.ShouldBindJSON(&product); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if err := service.UpdateShippingTemplate(product, req.Id); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("更新成功", ctx)
	}
}

// GetShippingTemplateSelect
func GetShippingTemplateSelect(ctx *gin.Context) {
	if data, err := service.GetShippingTemplateInfoSelect(); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(data, "获取成功", ctx)
	}
}

// GetShippingTemplateList
func GetShippingTemplateList(ctx *gin.Context) {
	var pageInfo request.ShippingTemplatePageInfo
	if errs := ctx.ShouldBindJSON(&pageInfo); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if list, total, err := service.GetShippingTemplateInfoList(pageInfo); err != nil {
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

// GetShippingTemplateById
func GetShippingTemplateById(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	product, err := service.GetShippingTemplateByID(req.Id)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithData(product, ctx)
	}
}

// DeleteShippingTemplate
func DeleteShippingTemplate(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.DeleteShippingTemplate(req.Id); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}
