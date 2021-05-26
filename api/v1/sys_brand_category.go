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

// CreateBrandCategory
func CreateBrandCategory(ctx iris.Context) {
	var brandCategory request.CreateSysBrandCategory
	if errs := utils.Verify(ctx.ReadJSON(&brandCategory)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if returnBrandCategory, err := service.CreateBrandCategory(brandCategory); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", ctx)
	} else {
		response.OkWithDetailed(getBrandCategoryMap(returnBrandCategory), "创建成功", ctx)
	}
}

// UpdateBrandCategory
func UpdateBrandCategory(ctx iris.Context) {
	var brandCategory request.UpdateSysBrandCategory
	if errs := utils.Verify(ctx.ReadJSON(&brandCategory)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if returnBrandCategory, err := service.UpdateBrandCategory(brandCategory); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", ctx)
	} else {
		response.OkWithDetailed(getBrandCategoryMap(returnBrandCategory), "更新成功", ctx)
	}
}

// getBrandCategoryMap
func getBrandCategoryMap(returnBrandCategory model.SysBrandCategory) context.Map {
	return iris.Map{
		"id":       returnBrandCategory.ID,
		"cateName": returnBrandCategory.CateName,
		"pid":      returnBrandCategory.Pid,
		"sort":     returnBrandCategory.Sort,
		"path":     returnBrandCategory.Path,
		"isShow":   returnBrandCategory.IsShow,
		"level":    returnBrandCategory.Level,
	}
}

// GetBrandCategoryList
func GetBrandCategoryList(ctx iris.Context) {
	var pageInfo request.PageInfo
	if errs := utils.Verify(ctx.ReadJSON(&pageInfo)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if list, total, err := service.GetBrandCategoryInfoList(pageInfo); err != nil {
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

// GetBrandCategoryById
func GetBrandCategoryById(ctx iris.Context) {
	var reqId request.GetById
	if errs := utils.Verify(ctx.ReadJSON(&reqId)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	brandCategory, err := service.GetBrandCategoryByID(reqId.Id)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithData(brandCategory, ctx)
	}
}

// DeleteBrandCategory
func DeleteBrandCategory(ctx iris.Context) {
	var reqId request.GetById
	if errs := utils.Verify(ctx.ReadJSON(&reqId)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.DeleteBrandCategory(reqId.Id); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}
