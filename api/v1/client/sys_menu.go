package client

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"go.uber.org/zap"
)

// GetMenu 获取用户动态路由
func GetMenu(ctx *gin.Context) {
	if menus, err := service.GetMenuTree(ctx); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(menus, "获取成功", ctx)
	}
}
