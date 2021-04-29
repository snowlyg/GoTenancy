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

// CreateSysOperationRecord 创建SysOperationRecord
func CreateSysOperationRecord(ctx iris.Context) {
	var sysOperationRecord model.SysOperationRecord
	_ = ctx.ReadJSON(&sysOperationRecord)
	if err := service.CreateSysOperationRecord(sysOperationRecord); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", ctx)
	} else {
		response.OkWithMessage("创建成功", ctx)
	}
}

// DeleteSysOperationRecord 删除SysOperationRecord
func DeleteSysOperationRecord(ctx iris.Context) {
	var sysOperationRecord model.SysOperationRecord
	_ = ctx.ReadJSON(&sysOperationRecord)
	if err := service.DeleteSysOperationRecord(sysOperationRecord); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}

// DeleteSysOperationRecordByIds 批量删除SysOperationRecord
func DeleteSysOperationRecordByIds(ctx iris.Context) {
	var IDS request.IdsReq
	_ = ctx.ReadJSON(&IDS)
	if err := service.DeleteSysOperationRecordByIds(IDS); err != nil {
		g.TENANCY_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", ctx)
	} else {
		response.OkWithMessage("批量删除成功", ctx)
	}
}

// FindSysOperationRecord 用id查询SysOperationRecord
func FindSysOperationRecord(ctx iris.Context) {
	var sysOperationRecord model.SysOperationRecord
	_ = ctx.ReadQuery(&sysOperationRecord)
	if err := utils.Verify(sysOperationRecord, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err, resysOperationRecord := service.GetSysOperationRecord(sysOperationRecord.ID); err != nil {
		g.TENANCY_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", ctx)
	} else {
		response.OkWithDetailed(iris.Map{"resysOperationRecord": resysOperationRecord}, "查询成功", ctx)
	}
}

// GetSysOperationRecordList 分页获取SysOperationRecord列表
func GetSysOperationRecordList(ctx iris.Context) {
	var pageInfo request.SysOperationRecordSearch
	_ = ctx.ReadQuery(&pageInfo)
	if err, list, total := service.GetSysOperationRecordInfoList(pageInfo); err != nil {
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
