package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"go.uber.org/zap"
)

// GetUpdateMediaMap
func GetUpdateMediaMap(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if form, err := service.GetMediaMap(req.Id, ctx); err != nil {
		g.TENANCY_LOG.Error("获取表单失败!", zap.Any("err", err))
		response.FailWithMessage("获取表单失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(form, "获取成功", ctx)
	}
}

func UploadFile(ctx *gin.Context) {
	noSave := ctx.DefaultQuery("noSave", "0")
	_, header, err := ctx.Request.FormFile("file")
	if err != nil {
		g.TENANCY_LOG.Error("接收文件失败!", zap.Any("err", err))
		response.FailWithMessage("接收文件失败", ctx)
		return
	}
	file, err := service.UploadFile(header, noSave, ctx) // 文件上传后拿到文件路径
	if err != nil {
		g.TENANCY_LOG.Error("修改数据库链接失败!", zap.Any("err", err))
		response.FailWithMessage("修改数据库链接失败", ctx)
		return
	}
	response.OkWithDetailed(gin.H{"src": file.Url, "id": file.ID}, "上传成功", ctx)
}

func DeleteFile(ctx *gin.Context) {
	var file request.DeleteMedia
	_ = ctx.ShouldBindJSON(&file)
	if err := service.DeleteFile(file.Ids); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败:"+err.Error(), ctx)
		return
	}
	response.OkWithMessage("删除成功", ctx)
}

func UpdateMediaName(ctx *gin.Context) {
	var req request.UpdateMediaName
	_ = ctx.ShouldBindJSON(&req)
	if err := service.UpdateMediaName(req, ctx.Param("id")); err != nil {
		g.TENANCY_LOG.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败:"+err.Error(), ctx)
		return
	}
	response.OkWithMessage("修改成功", ctx)
}

func GetFileList(ctx *gin.Context) {
	var pageInfo request.MediaPageInfo
	_ = ctx.ShouldBindJSON(&pageInfo)
	list, total, err := service.GetFileRecordInfoList(pageInfo, ctx)
	if err != nil {
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
