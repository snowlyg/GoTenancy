package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"go.uber.org/zap"
)

// EmailTest 发送测试邮件
func EmailTest(ctx *gin.Context) {
	if err := service.EmailTest(); err != nil {
		g.TENANCY_LOG.Error("发送失败!", zap.Any("err", err))
		response.FailWithMessage("发送失败", ctx)
	} else {
		response.OkWithData("发送成功", ctx)
	}
}
