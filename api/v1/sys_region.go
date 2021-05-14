package v1

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"go.uber.org/zap"
)

func Region(ctx iris.Context) {
	pCode := ctx.Params().GetIntDefault("p_code", -1)
	if regions, err := service.GetRegion(pCode); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(regions, "获取成功", ctx)
	}
}

func RegionList(ctx iris.Context) {
	if regions, err := service.GetRegionList(); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(regions, "获取成功", ctx)
	}
}
