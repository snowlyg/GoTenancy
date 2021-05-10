package v1

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/go-tenancy/utils"

	"github.com/snowlyg/go-tenancy/model/response"

	"github.com/snowlyg/go-tenancy/service"

	"github.com/snowlyg/go-tenancy/model/request"

	"github.com/snowlyg/go-tenancy/model"

	"github.com/snowlyg/go-tenancy/g"
	"go.uber.org/zap"
)

// CreateApi 创建基础api
func CreateApi(ctx iris.Context) {
	var api model.SysApi
	_ = ctx.ReadJSON(&api)
	if err := utils.Verify(api, utils.ApiVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := service.CreateApi(api); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", ctx)
	} else {
		response.OkWithMessage("创建成功", ctx)
	}
}

// DeleteApi 删除api
func DeleteApi(ctx iris.Context) {
	var api model.SysApi
	_ = ctx.ReadJSON(&api)
	if err := utils.Verify(api.TENANCY_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := service.DeleteApi(api); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}

// GetApiList 分页获取API列表
func GetApiList(ctx iris.Context) {
	var pageInfo request.SearchApiParams
	_ = ctx.ReadJSON(&pageInfo)
	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err, list, total := service.GetAPIInfoList(pageInfo.SysApi, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc); err != nil {
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

// GetApiById 根据id获取api
func GetApiById(ctx iris.Context) {
	var idInfo request.GetById
	_ = ctx.ReadJSON(&idInfo)
	if err := utils.Verify(idInfo, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	err, api := service.GetApiById(idInfo.Id)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithData(response.SysAPIResponse{Api: api}, ctx)
	}
}

// UpdateApi 更新基础api
func UpdateApi(ctx iris.Context) {
	var api model.SysApi
	_ = ctx.ReadJSON(&api)
	if err := utils.Verify(api, utils.ApiVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := service.UpdateApi(api); err != nil {
		g.TENANCY_LOG.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败", ctx)
	} else {
		response.OkWithMessage("修改成功", ctx)
	}
}

// GetAllApis 获取所有的Api 不分页
func GetAllApis(ctx iris.Context) {
	if err, apis := service.GetAllApis(); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(response.SysAPIListResponse{Apis: apis}, "获取成功", ctx)
	}
}

// DeleteApisByIds 删除选中Api
func DeleteApisByIds(ctx iris.Context) {
	var ids request.IdsReq
	_ = ctx.ReadJSON(&ids)
	if err := service.DeleteApisByIds(ids); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}
