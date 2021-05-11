package v1

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/go-tenancy/utils"
	"go.uber.org/zap"
)

// CreateTenancy
func CreateTenancy(ctx iris.Context) {
	var tenancy model.SysTenancy
	_ = ctx.ReadJSON(&tenancy)
	if err := utils.Verify(tenancy, utils.CreateTenancyVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
	}
	if returnTenancy, err := service.CreateTenancy(tenancy); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", ctx)
	} else {
		data := iris.Map{"id": returnTenancy.ID, "uuid": returnTenancy.UUID, "name": returnTenancy.Name, "tele": returnTenancy.Tele, "address": returnTenancy.Address, "businessTime": returnTenancy.BusinessTime, "sysRegionCode": returnTenancy.SysRegionCode}
		response.OkWithDetailed(data, "创建成功", ctx)
	}
}

// GetTenancyById
func GetTenancyById(ctx iris.Context) {
	var reqId request.GetById
	_ = ctx.ReadJSON(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	tenancy, err := service.GetTenancyByID(reqId.Id)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithData(tenancy, ctx)
	}
}

// UpdateTenancy
func UpdateTenancy(ctx iris.Context) {
	var tenancy model.SysTenancy
	_ = ctx.ReadJSON(&tenancy)
	if err := utils.Verify(tenancy, utils.UpdateTenancyVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
	}
	if returnTenancy, err := service.UpdateTenany(tenancy); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", ctx)
	} else {
		data := iris.Map{"id": returnTenancy.ID, "uuid": returnTenancy.UUID, "name": returnTenancy.Name, "tele": returnTenancy.Tele, "address": returnTenancy.Address, "businessTime": returnTenancy.BusinessTime, "sysRegionCode": returnTenancy.SysRegionCode}
		response.OkWithDetailed(data, "更新成功", ctx)
	}
}

// DeleteTenancy
func DeleteTenancy(ctx iris.Context) {
	var reqId request.GetById
	_ = ctx.ReadJSON(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := service.DeleteTenancy(reqId.Id); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}

// GetTenanciesList 分页获取商户列表
func GetTenanciesList(ctx iris.Context) {
	var pageInfo request.PageInfo
	_ = ctx.ReadJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if list, total, err := service.GetTenanciesInfoList(pageInfo); err != nil {
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

// GetTenanciesByRegion 根据区域获取商户列表，不分页
func GetTenanciesByRegion(ctx iris.Context) {
	code := ctx.Params().GetIntDefault("code", -1)
	if tenancies, err := service.GetTenanciesByRegion(code); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(tenancies, "获取成功", ctx)
	}
}
