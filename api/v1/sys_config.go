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

// CreateConfig
func CreateConfig(ctx iris.Context) {
	var config model.SysConfig
	if errs := utils.Verify(ctx.ReadJSON(&config)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if returnConfig, err := service.CreateConfig(config); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", ctx)
	} else {
		data := iris.Map{"id": returnConfig.ID, "name": returnConfig.Name, "type": returnConfig.Type, "value": returnConfig.Value}
		response.OkWithDetailed(data, "创建成功", ctx)
	}
}

// UpdateConfig
func UpdateConfig(ctx iris.Context) {
	var config model.SysConfig
	if errs := utils.Verify(ctx.ReadJSON(&config)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if returnConfig, err := service.UpdateConfig(config); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", ctx)
	} else {
		data := iris.Map{"id": returnConfig.ID, "name": returnConfig.Name, "type": returnConfig.Type, "value": returnConfig.Value}
		response.OkWithDetailed(data, "更新成功", ctx)
	}
}

// GetConfigList
func GetConfigList(ctx iris.Context) {
	var pageInfo request.PageInfo
	if errs := utils.Verify(ctx.ReadJSON(&pageInfo)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if list, total, err := service.GetConfigInfoList(pageInfo); err != nil {
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

// GetConfigByName
func GetConfigByName(ctx iris.Context) {
	var req request.GetSysConfig
	if errs := utils.Verify(ctx.ReadJSON(&req)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	config, err := service.GetConfigByName(req.Name, req.Type)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithData(config, ctx)
	}
}

// DeleteConfig
func DeleteConfig(ctx iris.Context) {
	var reqId request.GetById
	if errs := utils.Verify(ctx.ReadJSON(&reqId)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.DeleteConfig(reqId.Id); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}