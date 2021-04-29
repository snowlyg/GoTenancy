package v1

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"go.uber.org/zap"
)

// JsonInBlacklist jwt加入黑名单
func JsonInBlacklist(ctx iris.Context) {
	token := ctx.GetHeader("x-token")
	jwt := model.JwtBlacklist{Jwt: token}
	if err := service.JsonInBlacklist(jwt); err != nil {
		g.TENANCY_LOG.Error("jwt作废失败!", zap.Any("err", err))
		response.FailWithMessage("jwt作废失败", ctx)
	} else {
		response.OkWithMessage("jwt作废成功", ctx)
	}
}
