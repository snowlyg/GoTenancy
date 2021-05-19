package v1

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/multi"

	"github.com/snowlyg/go-tenancy/utils"

	"github.com/snowlyg/go-tenancy/model/request"

	"github.com/snowlyg/go-tenancy/model/response"

	"github.com/snowlyg/go-tenancy/model"

	"go.uber.org/zap"
)

// CreateAuthority 创建角色
func CreateAuthority(ctx iris.Context) {
	var authority model.SysAuthority
	if errs := utils.Verify(ctx.ReadJSON(&authority)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if authBack, err := service.CreateAuthority(authority); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(response.SysAuthorityResponse{Authority: authBack}, "创建成功", ctx)
	}
}

// CopyAuthority 拷贝角色
func CopyAuthority(ctx iris.Context) {
	var copyInfo response.SysAuthorityCopyResponse
	if errs := utils.Verify(ctx.ReadJSON(&copyInfo)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if errs := utils.Verify(ctx.ReadJSON(&copyInfo.Authority)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if authBack, err := service.CopyAuthority(copyInfo); err != nil {
		g.TENANCY_LOG.Error("拷贝失败!", zap.Any("err", err))
		response.FailWithMessage("拷贝失败"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(response.SysAuthorityResponse{Authority: authBack}, "拷贝成功", ctx)
	}
}

// DeleteAuthority 删除角色
func DeleteAuthority(ctx iris.Context) {
	var authority request.DeleteAuthority
	if errs := utils.Verify(ctx.ReadJSON(&authority)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.DeleteAuthority(&authority); err != nil { // 删除角色之前需要判断是否有用户正在使用此角色
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败"+err.Error(), ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}

// UpdateAuthority 更新角色信息
func UpdateAuthority(ctx iris.Context) {
	var auth model.SysAuthority
	if errs := utils.Verify(ctx.ReadJSON(&auth)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if authority, err := service.UpdateAuthority(auth); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(response.SysAuthorityResponse{Authority: authority}, "更新成功", ctx)
	}
}

// GetAdminAuthorityList 分页获取角色列表
func GetAdminAuthorityList(ctx iris.Context) {
	var pageInfo request.PageInfo
	if errs := utils.Verify(ctx.ReadJSON(&pageInfo)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if list, total, err := service.GetAuthorityInfoList(pageInfo, multi.AdminAuthority); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}

// GetTenancyAuthorityList 分页获取角色列表
func GetTenancyAuthorityList(ctx iris.Context) {
	var pageInfo request.PageInfo
	if errs := utils.Verify(ctx.ReadJSON(&pageInfo)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if list, total, err := service.GetAuthorityInfoList(pageInfo, multi.TenancyAuthority); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}

// GetGeneralAuthorityList 分页获取角色列表
func GetGeneralAuthorityList(ctx iris.Context) {
	var pageInfo request.PageInfo
	if errs := utils.Verify(ctx.ReadJSON(&pageInfo)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if list, total, err := service.GetAuthorityInfoList(pageInfo, multi.GeneralAuthority); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}

// GetAuthorityList 分页获取角色列表
func GetAuthorityList(ctx iris.Context) {
	var pageInfo request.PageInfo
	if errs := utils.Verify(ctx.ReadJSON(&pageInfo)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if list, total, err := service.GetAuthorityInfoList(pageInfo, multi.NoneAuthority); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}

// SetDataAuthority 设置角色资源权限
func SetDataAuthority(ctx iris.Context) {
	var auth request.SetDataAuthority
	if errs := utils.Verify(ctx.ReadJSON(&auth)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.SetDataAuthority(auth); err != nil {
		g.TENANCY_LOG.Error("设置失败!", zap.Any("err", err))
		response.FailWithMessage("设置失败"+err.Error(), ctx)
	} else {
		response.OkWithMessage("设置成功", ctx)
	}
}
