package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"go.uber.org/zap"
)

// CreateAttrTemplate
func CreateAttrTemplate(ctx *gin.Context) {
	var attrTemplate model.TenancyAttrTemplate
	if errs := ctx.ShouldBindJSON(&attrTemplate); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if returnAttrTemplate, err := service.CreateAttrTemplate(attrTemplate, ctx); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("添加失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(getAttrTemplateMap(returnAttrTemplate), "创建成功", ctx)
	}
}

// UpdateAttrTemplate
func UpdateAttrTemplate(ctx *gin.Context) {
	var attrTemplate model.TenancyAttrTemplate
	if errs := ctx.ShouldBindJSON(&attrTemplate); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if returnAttrTemplate, err := service.UpdateAttrTemplate(attrTemplate); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(getAttrTemplateMap(returnAttrTemplate), "更新成功", ctx)
	}
}

// getAttrTemplateMap
func getAttrTemplateMap(returnAttrTemplate model.TenancyAttrTemplate) gin.H {
	return gin.H{
		"id":            returnAttrTemplate.ID,
		"templateName":  returnAttrTemplate.TemplateName,
		"templateValue": returnAttrTemplate.TemplateValue,
		"sysTenancyId":  returnAttrTemplate.SysTenancyID,
	}
}

// GetAttrTemplateList
func GetAttrTemplateList(ctx *gin.Context) {
	var pageInfo request.PageInfo
	if errs := ctx.ShouldBindJSON(&pageInfo); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if list, total, err := service.GetAttrTemplateInfoList(pageInfo); err != nil {
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

// GetAttrTemplateById
func GetAttrTemplateById(ctx *gin.Context) {
	var reqId request.GetById
	if errs := ctx.ShouldBindJSON(&reqId); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	attrTemplate, err := service.GetAttrTemplateByID(reqId.Id)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithData(attrTemplate, ctx)
	}
}

// DeleteAttrTemplate
func DeleteAttrTemplate(ctx *gin.Context) {
	var reqId request.GetById
	if errs := ctx.ShouldBindJSON(&reqId); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.DeleteAttrTemplate(reqId.Id); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}
