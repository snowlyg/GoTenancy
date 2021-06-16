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

// GetCreateTenancyCategoryMap
func GetCreateTenancyCategoryMap(ctx *gin.Context) {
	if form, err := service.GetTenancyCategoryMap(0, ctx); err != nil {
		g.TENANCY_LOG.Error("获取表单失败!", zap.Any("err", err))
		response.FailWithMessage("获取表单失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(form, "获取成功", ctx)
	}
}

// GetUpdateTenancyCategoryMap
func GetUpdateTenancyCategoryMap(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if form, err := service.GetTenancyCategoryMap(req.Id, ctx); err != nil {
		g.TENANCY_LOG.Error("获取表单失败!", zap.Any("err", err))
		response.FailWithMessage("获取表单失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(form, "获取成功", ctx)
	}
}

// CreateCategory
func CreateCategory(ctx *gin.Context) {
	var category model.TenancyCategory
	if errs := ctx.ShouldBindJSON(&category); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if returnCategory, err := service.CreateCategory(category, ctx); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("添加失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(returnCategory, "创建成功", ctx)
	}
}

// UpdateCategory
func UpdateCategory(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	var category model.TenancyCategory
	if errs := ctx.ShouldBindJSON(&category); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if returnCategory, err := service.UpdateCategory(category, req.Id); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(returnCategory, "更新成功", ctx)
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

// GetCategorySelect
func GetCategorySelect(ctx *gin.Context) {
	if opts, err := service.GetTenacyCategoriesOptions(ctx); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(opts, "获取成功", ctx)
	}
}

// GetCategoryById
func GetCategoryById(ctx *gin.Context) {
	var reqId request.GetById
	if errs := ctx.ShouldBindUri(&reqId); errs != nil {
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

// ChangeCategoryStatus
func ChangeCategoryStatus(ctx *gin.Context) {
	var changeStatus request.ChangeStatus
	if errs := ctx.ShouldBindJSON(&changeStatus); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	err := service.ChangeCategoryStatus(changeStatus)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("设置成功", ctx)
	}
}

// DeleteCategory
func DeleteCategory(ctx *gin.Context) {
	var reqId request.GetById
	if errs := ctx.ShouldBindUri(&reqId); errs != nil {
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
