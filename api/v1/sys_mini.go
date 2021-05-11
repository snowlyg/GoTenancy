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

// CreateMini
func CreateMini(ctx iris.Context) {
	var mini model.SysMini
	_ = ctx.ReadJSON(&mini)
	if err := utils.Verify(mini, utils.CreateMiniVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
	}
	if returnMini, err := service.CreateMini(mini); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", ctx)
	} else {
		data := iris.Map{"id": returnMini.ID, "name": returnMini.Name, "appId": returnMini.AppID, "appSecret": returnMini.AppSecret, "uuid": returnMini.UUID, "remark": returnMini.Remark}
		response.OkWithDetailed(data, "创建成功", ctx)
	}
}

// UpdateMini
func UpdateMini(ctx iris.Context) {
	var mini model.SysMini
	_ = ctx.ReadJSON(&mini)
	if err := utils.Verify(mini, utils.UpdateMiniVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
	}
	if returnMini, err := service.UpdateMini(mini); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", ctx)
	} else {
		data := iris.Map{"id": returnMini.ID, "name": returnMini.Name, "appId": returnMini.AppID, "appSecret": returnMini.AppSecret, "uuid": returnMini.UUID, "remark": returnMini.Remark}
		response.OkWithDetailed(data, "更新成功", ctx)
	}
}

// GetMiniList
func GetMiniList(ctx iris.Context) {
	var pageInfo request.PageInfo
	_ = ctx.ReadJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if list, total, err := service.GetMiniInfoList(pageInfo); err != nil {
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

// GetMiniById
func GetMiniById(ctx iris.Context) {
	var reqId request.GetById
	_ = ctx.ReadJSON(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	mini, err := service.GetMiniByID(reqId.Id)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithData(mini, ctx)
	}
}

// DeleteMini
func DeleteMini(ctx iris.Context) {
	var reqId request.GetById
	_ = ctx.ReadJSON(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := service.DeleteMini(reqId.Id); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}
