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

// CreateBrand
func CreateBrand(ctx iris.Context) {
	var brand request.CreateSysBrand
	if errs := utils.Verify(ctx.ReadJSON(&brand)); errs != nil {
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
func UpdateBrand(ctx iris.Context) {
	var brand request.UpdateSysBrand
	if errs := utils.Verify(ctx.ReadJSON(&brand)); errs != nil {
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
func getBrandMap(returnBrand model.SysBrand) context.Map {
	return iris.Map{"id": returnBrand.ID, "brandName": returnBrand.BrandName, "sort": returnBrand.Sort, "pic": returnBrand.Pic, "isShow": returnBrand.IsShow}
}

// GetBrandList
func GetBrandList(ctx iris.Context) {
	var pageInfo request.PageInfo
	if errs := utils.Verify(ctx.ReadJSON(&pageInfo)); errs != nil {
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

// GetBrandById
func GetBrandById(ctx iris.Context) {
	var reqId request.GetById
	if errs := utils.Verify(ctx.ReadJSON(&reqId)); errs != nil {
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
func DeleteBrand(ctx iris.Context) {
	var reqId request.GetById
	if errs := utils.Verify(ctx.ReadJSON(&reqId)); errs != nil {
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
