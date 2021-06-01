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

// CreateBrand
func CreateBrand(ctx *gin.Context) {
	var brand request.CreateSysBrand
	if errs := ctx.ShouldBindJSON(&brand); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if returnBrand, err := service.CreateBrand(brand); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", ctx)
	} else {
		response.OkWithDetailed(getBrandMap(returnBrand), "创建成功", ctx)
	}
}

// UpdateBrand
func UpdateBrand(ctx *gin.Context) {
	var brand request.UpdateSysBrand
	if errs := ctx.ShouldBindJSON(&brand); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if returnBrand, err := service.UpdateBrand(brand); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", ctx)
	} else {
		response.OkWithDetailed(getBrandMap(returnBrand), "更新成功", ctx)
	}
}

// getBrandMap
func getBrandMap(returnBrand model.SysBrand) gin.H {
	return gin.H{"id": returnBrand.ID, "brandName": returnBrand.BrandName, "sort": returnBrand.Sort, "pic": returnBrand.Pic, "isShow": returnBrand.IsShow}
}

// GetBrandList
func GetBrandList(ctx *gin.Context) {
	var pageInfo request.PageInfo
	if errs := ctx.ShouldBindJSON(&pageInfo); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if list, total, err := service.GetBrandInfoList(pageInfo); err != nil {
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

// SetBrandCate
func SetBrandCate(ctx *gin.Context) {
	var setSysBrand request.SetSysBrand
	if errs := ctx.ShouldBindJSON(&setSysBrand); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	err := service.SetBrandCate(setSysBrand)
	if err != nil {
		g.TENANCY_LOG.Error("设置失败!", zap.Any("err", err))
		response.FailWithMessage("设置失败", ctx)
	} else {
		response.OkWithMessage("设置成功", ctx)
	}
}

// GetBrandById
func GetBrandById(ctx *gin.Context) {
	var reqId request.GetById
	if errs := ctx.ShouldBindJSON(&reqId); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	brand, err := service.GetBrandByID(reqId.Id)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithData(brand, ctx)
	}
}

// DeleteBrand
func DeleteBrand(ctx *gin.Context) {
	var reqId request.GetById
	if errs := ctx.ShouldBindJSON(&reqId); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.DeleteBrand(reqId.Id); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}
