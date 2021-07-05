package client

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"go.uber.org/zap"
)

// GetProductFilter
func GetProductFilter(ctx *gin.Context) {
	if filters, err := service.GetProductFilter(ctx); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(filters, "获取成功", ctx)
	}
}

// CreateProduct
func CreateProduct(ctx *gin.Context) {
	var product request.CreateProduct
	if errs := ctx.ShouldBindJSON(&product); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if returnProduct, err := service.CreateProduct(product, ctx); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("添加失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(returnProduct, "创建成功", ctx)
	}
}

// UpdateProduct
func UpdateProduct(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	var product request.UpdateProduct
	if errs := ctx.ShouldBindJSON(&product); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if err := service.UpdateProduct(product, req.Id); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("更新成功", ctx)
	}
}

// ChangeProductIsShow
func ChangeProductIsShow(ctx *gin.Context) {
	var changeStatus request.ChangeProductIsShow
	if errs := ctx.ShouldBindJSON(&changeStatus); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	err := service.ChangeProductIsShow(changeStatus)
	if err != nil {
		g.TENANCY_LOG.Error("设置失败!", zap.Any("err", err))
		response.FailWithMessage("设置失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("设置成功", ctx)
	}
}

// GetProductList
func GetProductList(ctx *gin.Context) {
	var pageInfo request.ProductPageInfo
	if errs := ctx.ShouldBindJSON(&pageInfo); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if list, total, err := service.GetProductInfoList(pageInfo, ctx); err != nil {
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

// GetProductById
func GetProductById(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	product, err := service.GetProductByID(req.Id)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithData(product, ctx)
	}
}

// DeleteProduct
func DeleteProduct(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.DeleteProduct(req.Id); err != nil {
		g.TENANCY_LOG.Error("操作失败!", zap.Any("err", err))
		response.FailWithMessage("操作失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("操作成功", ctx)
	}
}

// RestoreProduct
func RestoreProduct(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.RestoreProduct(req.Id); err != nil {
		g.TENANCY_LOG.Error("操作失败!", zap.Any("err", err))
		response.FailWithMessage("操作失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("操作成功", ctx)
	}
}

// DestoryProduct
func DestoryProduct(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.ForceDeleteProduct(req.Id); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}
