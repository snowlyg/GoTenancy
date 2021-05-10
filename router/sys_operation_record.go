package router

import (
	"github.com/kataras/iris/v12"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitSysOperationRecordRouter(Router iris.Party) {
	SysOperationRecordRouter := Router.Party("/sysOperationRecord")
	{
		SysOperationRecordRouter.Post("/createSysOperationRecord", v1.CreateSysOperationRecord)             // 新建SysOperationRecord
		SysOperationRecordRouter.Delete("/deleteSysOperationRecord", v1.DeleteSysOperationRecord)           // 删除SysOperationRecord
		SysOperationRecordRouter.Delete("/deleteSysOperationRecordByIds", v1.DeleteSysOperationRecordByIds) // 批量删除SysOperationRecord
		SysOperationRecordRouter.Get("/findSysOperationRecord", v1.FindSysOperationRecord)                  // 根据ID获取SysOperationRecord
		SysOperationRecordRouter.Get("/getSysOperationRecordList", v1.GetSysOperationRecordList)            // 获取SysOperationRecord列表

	}
}
