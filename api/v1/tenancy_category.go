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

// CreateCategory
func CreateCategory(ctx *gin.Context) {
	var category request.CreateTenancyCategory
	if errs := ctx.ShouldBindJSON(&category); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if returnCategory, err := service.CreateCategory(category, ctx); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", ctx)
	} else {
		response.OkWithDetailed(getCategoryMap(returnCategory), "创建成功", ctx)
	}
}

// UpdateCategory
func UpdateCategory(ctx *gin.Context) {
	var category request.UpdateTenancyCategory
	if errs := ctx.ShouldBindJSON(&category); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if returnCategory, err := service.UpdateCategory(category); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(getCategoryMap(returnCategory), "更新成功", ctx)
	}
}

// getCategoryMap
func getCategoryMap(returnCategory model.TenancyCategory) gin.H {
	return gin.H{
		"id":       returnCategory.ID,
		"cateName": returnCategory.CateName,
		"pid":      returnCategory.Pid,
		"sort":     returnCategory.Sort,
		"path":     returnCategory.Path,
		"isShow":   returnCategory.IsShow,
		"level":    returnCategory.Level,
		"pic":      returnCategory.Pic,
	}
}

// GetCategoryList
func GetCategoryList(ctx *gin.Context) {
	var pageInfo request.PageInfo
	if errs := ctx.ShouldBindJSON(&pageInfo); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if list, err := service.GetCategoryInfoList(ctx); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    0,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}

// GetCategoryById
func GetCategoryById(ctx *gin.Context) {
	var reqId request.GetById
	if errs := ctx.ShouldBindJSON(&reqId); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	category, err := service.GetCategoryByID(reqId.Id)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithData(category, ctx)
	}
}

// DeleteCategory
func DeleteCategory(ctx *gin.Context) {
	var reqId request.GetById
	if errs := ctx.ShouldBindJSON(&reqId); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.DeleteCategory(reqId.Id); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}
