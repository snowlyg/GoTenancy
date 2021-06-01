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

func UploadFile(ctx *gin.Context) {
	var file model.TenancyMedia
	noSave := ctx.DefaultQuery("noSave", "0")
	path := ctx.DefaultPostForm("path", "")
	_, header, err := ctx.Request.FormFile("file")
	if err != nil {
		g.TENANCY_LOG.Error("接收文件失败!", zap.Any("err", err))
		response.FailWithMessage("接收文件失败", ctx)
		return
	}
	file, err = service.UploadFile(header, noSave, path, ctx) // 文件上传后拿到文件路径
	if err != nil {
		g.TENANCY_LOG.Error("修改数据库链接失败!", zap.Any("err", err))
		response.FailWithMessage("修改数据库链接失败", ctx)
		return
	}
	response.OkWithDetailed(response.TenancyMedia{File: file}, "上传成功", ctx)
}

func DeleteFile(ctx *gin.Context) {
	var file model.TenancyMedia
	_ = ctx.ShouldBindJSON(&file)
	if err := service.DeleteFile(file); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", ctx)
		return
	}
	response.OkWithMessage("删除成功", ctx)
}

func GetFileList(ctx *gin.Context) {
	var pageInfo request.PageInfo
	_ = ctx.ShouldBindJSON(&pageInfo)
	list, total, err := service.GetFileRecordInfoList(pageInfo, ctx)
	if err != nil {
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
