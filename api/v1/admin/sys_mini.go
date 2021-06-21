package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"go.uber.org/zap"
)

// CreateMini
func CreateMini(ctx *gin.Context) {
	var mini request.CreateSysMini
	if errs := ctx.ShouldBindJSON(&mini); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if returnMini, err := service.CreateMini(mini); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("添加失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(getMiniMap(returnMini), "创建成功", ctx)
	}
}

// UpdateMini
func UpdateMini(ctx *gin.Context) {
	var mini request.UpdateSysMini
	if errs := ctx.ShouldBindJSON(&mini); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if returnMini, err := service.UpdateMini(mini); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(getMiniMap(returnMini), "更新成功", ctx)
	}
}

// getMiniMap
func getMiniMap(returnMini model.SysMini) gin.H {
	return gin.H{"id": returnMini.ID, "name": returnMini.Name, "appId": returnMini.AppID, "appSecret": returnMini.AppSecret, "uuid": returnMini.UUID, "remark": returnMini.Remark}
}

// GetMiniList
func GetMiniList(ctx *gin.Context) {
	var pageInfo request.PageInfo
	if errs := ctx.ShouldBindJSON(&pageInfo); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if list, total, err := service.GetMiniInfoList(pageInfo); err != nil {
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

// GetMiniById
func GetMiniById(ctx *gin.Context) {
	var reqId request.GetById
	if errs := ctx.ShouldBindJSON(&reqId); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	mini, err := service.GetMiniByID(reqId.Id)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithData(mini, ctx)
	}
}

// DeleteMini
func DeleteMini(ctx *gin.Context) {
	var reqId request.GetById
	if errs := ctx.ShouldBindJSON(&reqId); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.DeleteMini(reqId.Id); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}
