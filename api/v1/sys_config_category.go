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

// GetCreateConfigCategoryMap
func GetCreateConfigCategoryMap(ctx *gin.Context) {
	if form, err := service.GetConfigCategoryMap(0, ctx); err != nil {
		g.TENANCY_LOG.Error("获取表单失败!", zap.Any("err", err))
		response.FailWithMessage("获取表单失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(form, "获取成功", ctx)
	}
}

// GetUpdateConfigCategoryMap
func GetUpdateConfigCategoryMap(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if form, err := service.GetConfigCategoryMap(req.Id, ctx); err != nil {
		g.TENANCY_LOG.Error("获取表单失败!", zap.Any("err", err))
		response.FailWithMessage("获取表单失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(form, "获取成功", ctx)
	}
}

// CreateConfigCategory
func CreateConfigCategory(ctx *gin.Context) {
	var cate model.SysConfigCategory
	if errs := ctx.ShouldBindJSON(&cate); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if returnCate, err := service.CreateConfigCategory(cate); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("添加失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(gin.H{
			"id":     returnCate.ID,
			"name":   returnCate.Name,
			"sort":   returnCate.Sort,
			"key":    returnCate.Key,
			"info":   returnCate.Info,
			"icon":   returnCate.Icon,
			"status": returnCate.Status,
		}, "创建成功", ctx)
	}
}

// UpdateConfigCategory
func UpdateConfigCategory(ctx *gin.Context) {
	var cate model.SysConfigCategory
	if errs := ctx.ShouldBindJSON(&cate); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if returnCate, err := service.UpdateConfigCategory(cate, ctx.Param("id")); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(gin.H{
			"name":   returnCate.Name,
			"sort":   returnCate.Sort,
			"key":    returnCate.Key,
			"info":   returnCate.Info,
			"icon":   returnCate.Icon,
			"status": returnCate.Status,
		}, "更新成功", ctx)
	}
}

// GetConfigCategoryList
func GetConfigCategoryList(ctx *gin.Context) {
	if list, err := service.GetConfigCategoriesInfoList(); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:  list,
			Total: 0,
		}, "获取成功", ctx)
	}
}

// ChangeConfigCategoryStatus
func ChangeConfigCategoryStatus(ctx *gin.Context) {
	var changeStatus request.ChangeStatus
	if errs := ctx.ShouldBindJSON(&changeStatus); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	err := service.ChangeConfigCategoryStatus(changeStatus)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("设置成功", ctx)
	}
}

// GetConfigCategoryById
func GetConfigCategoryById(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	configCategory, err := service.GetConfigCategoryByID(req.Id)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithData(configCategory, ctx)
	}
}

// DeleteConfigCategory
func DeleteConfigCategory(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.DeleteConfigCategory(req.Id); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}
