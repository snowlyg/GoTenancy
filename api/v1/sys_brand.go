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

// GetCreateBrandMap
func GetCreateBrandMap(ctx *gin.Context) {
	if form, err := service.GetBrandMap(0, ctx); err != nil {
		g.TENANCY_LOG.Error("获取表单失败!", zap.Any("err", err))
		response.FailWithMessage("获取表单失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(form, "获取成功", ctx)
	}
}

// GetUpdateBrandMap
func GetUpdateBrandMap(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if form, err := service.GetBrandMap(req.Id, ctx); err != nil {
		g.TENANCY_LOG.Error("获取表单失败!", zap.Any("err", err))
		response.FailWithMessage("获取表单失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(form, "获取成功", ctx)
	}
}

// CreateBrand
func CreateBrand(ctx *gin.Context) {
	var brand model.SysBrand
	if errs := ctx.ShouldBindJSON(&brand); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if returnBrand, err := service.CreateBrand(brand); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("添加失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(returnBrand, "创建成功", ctx)
	}
}

// UpdateBrand
func UpdateBrand(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	var brand model.SysBrand
	if errs := ctx.ShouldBindJSON(&brand); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if returnBrand, err := service.UpdateBrand(brand, req.Id); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(returnBrand, "更新成功", ctx)
	}
}

// ChangeBrandStatus
func ChangeBrandStatus(ctx *gin.Context) {
	var changeStatus request.ChangeStatus
	if errs := ctx.ShouldBindJSON(&changeStatus); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	err := service.ChangeBrandStatus(changeStatus)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("设置成功", ctx)
	}
}

// GetBrandList
func GetBrandList(ctx *gin.Context) {
	var pageInfo request.BrandPageInfo
	if errs := ctx.ShouldBindJSON(&pageInfo); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if list, total, err := service.GetBrandInfoList(pageInfo); err != nil {
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

// GetBrandById
func GetBrandById(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	brand, err := service.GetBrandByID(req.Id)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithData(brand, ctx)
	}
}

// DeleteBrand
func DeleteBrand(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.DeleteBrand(req.Id); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}
