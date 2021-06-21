package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"go.uber.org/zap"
)

// SaveConfigValue
func SaveConfigValue(ctx *gin.Context) {
	var req request.GetByConfigCate
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	var config map[string]interface{}
	if errs := ctx.ShouldBindJSON(&config); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if err := service.SaveConfigValue(config, req.Cate, ctx); err != nil {
		g.TENANCY_LOG.Error("操作失败!", zap.Any("err", err))
		response.FailWithMessage("操作失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("操作成功", ctx)
	}
}
