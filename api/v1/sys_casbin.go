package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"

	"github.com/snowlyg/go-tenancy/service"

	"github.com/snowlyg/go-tenancy/model/request"

	"github.com/snowlyg/go-tenancy/model/response"

	"go.uber.org/zap"
)

// UpdateCasbin 更新角色api权限
func UpdateCasbin(ctx *gin.Context) {
	var cmr request.CasbinInReceive
	if errs := ctx.ShouldBindJSON(&cmr); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.UpdateCasbin(cmr.AuthorityId, cmr.CasbinInfos); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", ctx)
	} else {
		response.OkWithMessage("更新成功", ctx)
	}
}

// GetPolicyPathByAuthorityId 获取权限列表
func GetPolicyPathByAuthorityId(ctx *gin.Context) {
	var casbin request.CasbinInReceive
	if errs := ctx.ShouldBindJSON(&casbin); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	paths := service.GetPolicyPathByAuthorityId(casbin.AuthorityId)
	response.OkWithDetailed(response.PolicyPathResponse{Paths: paths}, "获取成功", ctx)
}
