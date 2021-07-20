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

// GetCreateProductCategoryMap
func GetCreateProductCategoryMap(ctx *gin.Context) {
	if form, err := service.GetProductCategoryMap(0, ctx); err != nil {
		g.TENANCY_LOG.Error("获取表单失败!", zap.Any("err", err))
		response.FailWithMessage("获取表单失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(form, "获取成功", ctx)
	}
}

// GetUpdateProductCategoryMap
func GetUpdateProductCategoryMap(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if form, err := service.GetProductCategoryMap(req.Id, ctx); err != nil {
		g.TENANCY_LOG.Error("获取表单失败!", zap.Any("err", err))
		response.FailWithMessage("获取表单失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(form, "获取成功", ctx)
	}
}

// CreateProductCategory
func CreateProductCategory(ctx *gin.Context) {
	var category model.ProductCategory
	if errs := ctx.ShouldBindJSON(&category); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if returnCategory, err := service.CreateProductCategory(category, multi.GetTenancyId(ctx)); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("添加失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(returnCategory, "创建成功", ctx)
	}
}

// UpdateProductCategory
func UpdateProductCategory(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	var category model.ProductCategory
	if errs := ctx.ShouldBindJSON(&category); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if returnCategory, err := service.UpdateProductCategory(category, req.Id); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(returnCategory, "更新成功", ctx)
	}
}

// GetProductCategoryList
func GetProductCategoryList(ctx *gin.Context) {
	if list, err := service.GetProductCategoryInfoList(0); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(list, "获取成功", ctx)
	}
}

// GetClientProductCategoryList
func GetClientProductCategoryList(ctx *gin.Context) {
	if list, err := service.GetProductCategoryInfoList(multi.GetTenancyId(ctx)); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(list, "获取成功", ctx)
	}
}

// GetProductCategorySelect
func GetProductCategorySelect(ctx *gin.Context) {
	if opts, err := service.GetProductCategoriesOptions(multi.GetTenancyId(ctx)); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(opts, "获取成功", ctx)
	}
}

// GetAdminProductCategorySelect
func GetAdminProductCategorySelect(ctx *gin.Context) {
	if opts, err := service.GetProductCategoriesOptions(0); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(opts, "获取成功", ctx)
	}
}

// GetProductCategoryById
func GetProductCategoryById(ctx *gin.Context) {
	var reqId request.GetById
	if errs := ctx.ShouldBindUri(&reqId); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	category, err := service.GetProductCategoryByID(reqId.Id)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithData(category, ctx)
	}
}

// ChangeProductCategoryStatus
func ChangeProductCategoryStatus(ctx *gin.Context) {
	var changeStatus request.ChangeStatus
	if errs := ctx.ShouldBindJSON(&changeStatus); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	err := service.ChangeProductCategoryStatus(changeStatus)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("设置成功", ctx)
	}
}

// DeleteProductCategory
func DeleteProductCategory(ctx *gin.Context) {
	var reqId request.GetById
	if errs := ctx.ShouldBindUri(&reqId); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.DeleteProductCategory(reqId.Id); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}
