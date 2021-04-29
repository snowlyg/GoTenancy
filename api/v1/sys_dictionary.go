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

// CreateSysDictionary 创建SysDictionary
func CreateSysDictionary(ctx iris.Context) {
	var dictionary model.SysDictionary
	_ = ctx.ReadJSON(&dictionary)
	if err := service.CreateSysDictionary(dictionary); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", ctx)
	} else {
		response.OkWithMessage("创建成功", ctx)
	}
}

// DeleteSysDictionary 删除SysDictionary
func DeleteSysDictionary(ctx iris.Context) {
	var dictionary model.SysDictionary
	_ = ctx.ReadJSON(&dictionary)
	if err := service.DeleteSysDictionary(dictionary); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}

// UpdateSysDictionary 更新SysDictionary
func UpdateSysDictionary(ctx iris.Context) {
	var dictionary model.SysDictionary
	_ = ctx.ReadJSON(&dictionary)
	if err := service.UpdateSysDictionary(&dictionary); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", ctx)
	} else {
		response.OkWithMessage("更新成功", ctx)
	}
}

// FindSysDictionary 用id查询SysDictionary
func FindSysDictionary(ctx iris.Context) {
	var dictionary model.SysDictionary
	_ = ctx.ReadQuery(&dictionary)
	if err, sysDictionary := service.GetSysDictionary(dictionary.Type, dictionary.ID); err != nil {
		g.TENANCY_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", ctx)
	} else {
		response.OkWithDetailed(iris.Map{"resysDictionary": sysDictionary}, "查询成功", ctx)
	}
}

// GetSysDictionaryList 分页获取SysDictionary列表
func GetSysDictionaryList(ctx iris.Context) {
	var pageInfo request.SysDictionarySearch
	_ = ctx.ReadQuery(&pageInfo)
	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err, list, total := service.GetSysDictionaryInfoList(pageInfo); err != nil {
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
