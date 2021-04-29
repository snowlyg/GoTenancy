package v1

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/utils"

	"github.com/snowlyg/go-tenancy/service"

	"github.com/snowlyg/go-tenancy/model/response"

	"github.com/snowlyg/go-tenancy/model/request"

	"github.com/snowlyg/go-tenancy/model"

	"go.uber.org/zap"
)

// CreateSysDictionaryDetail 创建SysDictionaryDetail
func CreateSysDictionaryDetail(ctx iris.Context) {
	var detail model.SysDictionaryDetail
	_ = ctx.ReadJSON(&detail)
	if err := service.CreateSysDictionaryDetail(detail); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", ctx)
	} else {
		response.OkWithMessage("创建成功", ctx)
	}
}

// DeleteSysDictionaryDetail  删除SysDictionaryDetail
func DeleteSysDictionaryDetail(ctx iris.Context) {
	var detail model.SysDictionaryDetail
	_ = ctx.ReadJSON(&detail)
	if err := service.DeleteSysDictionaryDetail(detail); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}

// UpdateSysDictionaryDetail 更新SysDictionaryDetail
func UpdateSysDictionaryDetail(ctx iris.Context) {
	var detail model.SysDictionaryDetail
	_ = ctx.ReadJSON(&detail)
	if err := service.UpdateSysDictionaryDetail(&detail); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", ctx)
	} else {
		response.OkWithMessage("更新成功", ctx)
	}
}

// FindSysDictionaryDetail 用id查询SysDictionaryDetail
func FindSysDictionaryDetail(ctx iris.Context) {
	var detail model.SysDictionaryDetail
	_ = ctx.ReadQuery(&detail)
	if err := utils.Verify(detail, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err, resysDictionaryDetail := service.GetSysDictionaryDetail(detail.ID); err != nil {
		g.TENANCY_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", ctx)
	} else {
		response.OkWithDetailed(iris.Map{"resysDictionaryDetail": resysDictionaryDetail}, "查询成功", ctx)
	}
}

// GetSysDictionaryDetailList 分页获取SysDictionaryDetail列表
func GetSysDictionaryDetailList(ctx iris.Context) {
	var pageInfo request.SysDictionaryDetailSearch
	_ = ctx.ReadQuery(&pageInfo)
	if err, list, total := service.GetSysDictionaryDetailInfoList(pageInfo); err != nil {
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
