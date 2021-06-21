package public

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"go.uber.org/zap"
)

// InitDB 初始化用户数据库
func InitDB(ctx *gin.Context) {
	if g.TENANCY_DB != nil {
		g.TENANCY_LOG.Error("非法访问")
		response.FailWithMessage("非法访问", ctx)
		return
	}
	var dbInfo request.InitDB
	if err := ctx.ShouldBindJSON(&dbInfo); err != nil {
		g.TENANCY_LOG.Error("参数校验不通过", zap.Any("err", err))
		response.FailWithMessage("参数校验不通过", ctx)
		return
	}
	if err := service.InitDB(dbInfo); err != nil {
		g.TENANCY_LOG.Error("自动创建数据库失败", zap.Any("err", err))
		response.FailWithMessage("自动创建数据库失败，请查看后台日志", ctx)
		return
	}
	response.OkWithData("自动创建数据库成功", ctx)
}

// CheckDB 初始化用户数据库
func CheckDB(ctx *gin.Context) {
	if g.TENANCY_DB != nil {
		g.TENANCY_LOG.Info("数据库无需初始化")
		response.OkWithDetailed(gin.H{
			"needInit": false,
		}, "数据库无需初始化", ctx)
		return
	} else {
		g.TENANCY_LOG.Info("前往初始化数据库")
		response.OkWithDetailed(gin.H{
			"needInit": true,
		}, "前往初始化数据库", ctx)
		return
	}
}
