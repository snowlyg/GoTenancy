package v1

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/go-tenancy/utils"
	"go.uber.org/zap"
)

// CreateCategory
func CreateCategory(ctx iris.Context) {
	var category request.CreateTenancyCategory
	if errs := utils.Verify(ctx.ReadJSON(&category)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if returnCategory, err := service.CreateCategory(category); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", ctx)
	} else {
		response.OkWithDetailed(getCategoryMap(returnCategory), "创建成功", ctx)
	}
}

// UpdateCategory
func UpdateCategory(ctx iris.Context) {
	var category request.UpdateTenancyCategory
	if errs := utils.Verify(ctx.ReadJSON(&category)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if returnCategory, err := service.UpdateCategory(category); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", ctx)
	} else {
		response.OkWithDetailed(getCategoryMap(returnCategory), "更新成功", ctx)
	}
}

// getCategoryMap
func getCategoryMap(returnCategory model.TenancyCategory) context.Map {
	return iris.Map{
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
func GetCategoryList(ctx iris.Context) {
	var pageInfo request.PageInfo
	if errs := utils.Verify(ctx.ReadJSON(&pageInfo)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if list, err := service.GetCategoryInfoList(); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
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
func GetCategoryById(ctx iris.Context) {
	var reqId request.GetById
	if errs := utils.Verify(ctx.ReadJSON(&reqId)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	category, err := service.GetCategoryByID(reqId.Id)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithData(category, ctx)
	}
}

// DeleteCategory
func DeleteCategory(ctx iris.Context) {
	var reqId request.GetById
	if errs := utils.Verify(ctx.ReadJSON(&reqId)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.DeleteCategory(reqId.Id); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}
