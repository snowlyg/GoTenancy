package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"go.uber.org/zap"
)

// CreateSysOperationRecord 创建SysOperationRecord
func CreateSysOperationRecord(ctx *gin.Context) {
	var sysOperationRecord model.SysOperationRecord
	_ = ctx.ShouldBindJSON(&sysOperationRecord)
	if err := service.CreateSysOperationRecord(sysOperationRecord); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", ctx)
	} else {
		response.OkWithMessage("创建成功", ctx)
	}
}

// DeleteSysOperationRecord 删除SysOperationRecord
func DeleteSysOperationRecord(ctx *gin.Context) {
	var sysOperationRecord model.SysOperationRecord
	_ = ctx.ShouldBindJSON(&sysOperationRecord)
	if err := service.DeleteSysOperationRecord(sysOperationRecord); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}

// DeleteSysOperationRecordByIds 批量删除SysOperationRecord
func DeleteSysOperationRecordByIds(ctx *gin.Context) {
	var IDS request.IdsReq
	_ = ctx.ShouldBindJSON(&IDS)
	if err := service.DeleteSysOperationRecordByIds(IDS); err != nil {
		g.TENANCY_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", ctx)
	} else {
		response.OkWithMessage("批量删除成功", ctx)
	}
}

// FindSysOperationRecord 用id查询SysOperationRecord
func FindSysOperationRecord(ctx *gin.Context) {
	var sysOperationRecord model.SysOperationRecord
	if errs := ctx.ShouldBindJSON(&sysOperationRecord); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if resysOperationRecord, err := service.GetSysOperationRecord(sysOperationRecord.ID); err != nil {
		g.TENANCY_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", ctx)
	} else {
		response.OkWithDetailed(gin.H{"resysOperationRecord": resysOperationRecord}, "查询成功", ctx)
	}
}

// GetSysOperationRecordList 分页获取SysOperationRecord列表
func GetSysOperationRecordList(ctx *gin.Context) {
	var pageInfo request.SysOperationRecordSearch
	_ = ctx.ShouldBindQuery(&pageInfo)
	if list, total, err := service.GetSysOperationRecordInfoList(pageInfo); err != nil {
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
