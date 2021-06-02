package v1

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"

	"go.uber.org/zap"
)

// GetSystemConfig 获取配置文件内容
func GetSystemConfig(ctx *gin.Context) {
	response.OkWithDetailed(response.SysConfigResponse{Config: service.GetSystemConfig()}, "获取成功", ctx)
}

// SetSystemConfig 设置配置文件内容
func SetSystemConfig(ctx *gin.Context) {
	var sys model.System
	_ = ctx.ShouldBindJSON(&sys)
	if err := service.SetSystemConfig(sys); err != nil {
		g.TENANCY_LOG.Error("设置失败!", zap.Any("err", err))
		response.FailWithMessage("设置失败", ctx)
	} else {
		response.OkWithData("设置成功", ctx)
	}
}

// ReloadSystem 重启系统
func ReloadSystem(ctx *gin.Context) {
	if runtime.GOOS == "windows" {
		response.FailWithMessage("系统不支持", ctx)
		return
	}
	pid := os.Getpid()
	cmd := exec.Command("kill", "-1", strconv.Itoa(pid))
	err := cmd.Run()
	if err != nil {
		g.TENANCY_LOG.Error("重启系统失败!", zap.Any("err", err))
		response.FailWithMessage("重启系统失败", ctx)
		return
	}
	response.OkWithMessage("重启系统成功", ctx)
}

// GetServerInfo 获取服务器信息
func GetServerInfo(ctx *gin.Context) {
	if server, err := service.GetServerInfo(); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
		return
	} else {
		response.OkWithDetailed(gin.H{"server": server}, "获取成功", ctx)
	}

}
