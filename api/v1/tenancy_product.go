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

// CreateProduct
func CreateProduct(ctx iris.Context) {
	var product request.CreateTenancyProduct
	if errs := utils.Verify(ctx.ReadJSON(&product)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if returnProduct, err := service.CreateProduct(product, ctx); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", ctx)
	} else {
		response.OkWithDetailed(getProductMap(returnProduct), "创建成功", ctx)
	}
}

// UpdateProduct
func UpdateProduct(ctx iris.Context) {
	var product request.UpdateTenancyProduct
	if errs := utils.Verify(ctx.ReadJSON(&product)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if returnProduct, err := service.UpdateProduct(product); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", ctx)
	} else {
		response.OkWithDetailed(getProductMap(returnProduct), "更新成功", ctx)
	}
}

// getProductMap
func getProductMap(returnProduct model.TenancyProduct) context.Map {
	return iris.Map{
		"id":                returnProduct.ID,
		"storeName":         returnProduct.StoreName,
		"storeInfo":         returnProduct.StoreInfo,
		"keyword":           returnProduct.Keyword,
		"barCode":           returnProduct.BarCode,
		"isShow":            returnProduct.IsShow,
		"status":            returnProduct.Status,
		"tenancyStatus":     returnProduct.TenancyStatus,
		"unitName":          returnProduct.UnitName,
		"sort":              returnProduct.Sort,
		"rank":              returnProduct.Rank,
		"sales":             returnProduct.Sales,
		"price":             returnProduct.Price,
		"cost":              returnProduct.Cost,
		"otPrice":           returnProduct.OtPrice,
		"stock":             returnProduct.Stock,
		"isHot":             returnProduct.IsHot,
		"isBenefit":         returnProduct.IsBenefit,
		"isBest":            returnProduct.IsBest,
		"isNew":             returnProduct.IsNew,
		"isGood":            returnProduct.IsGood,
		"productType":       returnProduct.ProductType,
		"ficti":             returnProduct.Ficti,
		"browse":            returnProduct.Browse,
		"codePath":          returnProduct.CodePath,
		"videoLink":         returnProduct.VideoLink,
		"specType":          returnProduct.SpecType,
		"extensionType":     returnProduct.ExtensionType,
		"refusal":           returnProduct.Refusal,
		"rate":              returnProduct.Rate,
		"replyCount":        returnProduct.ReplyCount,
		"giveCouponIds":     returnProduct.GiveCouponIDs,
		"isGiftBag":         returnProduct.IsGiftBag,
		"careCount":         returnProduct.CareCount,
		"image":             returnProduct.Image,
		"sliderImage":       returnProduct.SliderImage,
		"oldId":             returnProduct.OldID,
		"tempId":            returnProduct.TempID,
		"sysTenancyId":      returnProduct.SysTenancyID,
		"sysBrandId":        returnProduct.SysBrandID,
		"tenancyCategoryId": returnProduct.TenancyCategoryID,
	}
}

// GetProductList
func GetProductList(ctx iris.Context) {
	var pageInfo request.PageInfo
	if errs := utils.Verify(ctx.ReadJSON(&pageInfo)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if list, total, err := service.GetProductInfoList(pageInfo, ctx); err != nil {
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

// GetProductById
func GetProductById(ctx iris.Context) {
	var reqId request.GetById
	if errs := utils.Verify(ctx.ReadJSON(&reqId)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	product, err := service.GetProductByID(reqId.Id)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithData(product, ctx)
	}
}

// DeleteProduct
func DeleteProduct(ctx iris.Context) {
	var reqId request.GetById
	if errs := utils.Verify(ctx.ReadJSON(&reqId)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.DeleteProduct(reqId.Id); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}
