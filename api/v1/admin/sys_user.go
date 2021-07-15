package admin

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

// RegisterAdmin 员工注册
func RegisterAdmin(ctx *gin.Context) {
	var R request.Register
	if errs := ctx.ShouldBindJSON(&R); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	user := &model.SysUser{Username: R.Username, Password: R.Password, AuthorityId: R.AuthorityId}
	userReturn, err := service.Register(*user, multi.AdminAuthority)
	if err != nil {
		g.TENANCY_LOG.Error("注册失败", zap.Any("err", err))
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.OkWithDetailed(gin.H{"id": userReturn.ID, "userName": userReturn.Username, "authorityId": userReturn.AuthorityId}, "注册成功", ctx)
	}
}

// RegisterTenancy 商户注册
func RegisterTenancy(ctx *gin.Context) {
	var R request.Register
	if errs := ctx.ShouldBindJSON(&R); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	user := &model.SysUser{Username: R.Username, Password: R.Password, AuthorityId: R.AuthorityId}
	userReturn, err := service.Register(*user, multi.TenancyAuthority)
	if err != nil {
		g.TENANCY_LOG.Error("注册失败", zap.Any("err", err))
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.OkWithDetailed(gin.H{"id": userReturn.ID, "userName": userReturn.Username, "authorityId": userReturn.AuthorityId}, "注册成功", ctx)
	}
}

// ChangePassword 用户修改密码
func ChangePassword(ctx *gin.Context) {
	var user request.ChangePassword
	if errs := ctx.ShouldBindJSON(&user); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if user.NewPassword != user.ConfirmPassword {
		response.FailWithMessage("两次输入密码不一致", ctx)
		return
	}
	U := &model.SysUser{Username: multi.GetUsername(ctx), Password: user.Password}
	err := service.ChangePassword(U, user.NewPassword, multi.GetAuthorityType(ctx))
	if err != nil {
		g.TENANCY_LOG.Error("修改失败", zap.Any("err", err))
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.OkWithMessage("修改成功", ctx)
	}
}

// ChangeProfile 用户修改信息
func ChangeProfile(ctx *gin.Context) {
	var user request.ChangeProfile
	if errs := ctx.ShouldBindJSON(&user); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	err := service.ChangeProfile(user, multi.GetUserId(ctx))
	if err != nil {
		g.TENANCY_LOG.Error("修改失败", zap.Any("err", err))
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.OkWithMessage("修改成功", ctx)
	}
}

// GetAdminList 分页获取用户列表
func GetAdminList(ctx *gin.Context) {
	var pageInfo request.PageInfo
	if errs := ctx.ShouldBindJSON(&pageInfo); errs != nil {
		fmt.Printf("ShouldBindJSON %v\n\n", errs)
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if list, total, err := service.GetAdminInfoList(pageInfo); err != nil {
		g.TENANCY_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}

// GetTenancyList 分页获取用户列表
func GetTenancyList(ctx *gin.Context) {
	var pageInfo request.PageInfo
	if errs := ctx.ShouldBindJSON(&pageInfo); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if list, total, err := service.GetTenancyInfoList(pageInfo); err != nil {
		g.TENANCY_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}

// GetGeneralList 分页获取用户列表
func GetGeneralList(ctx *gin.Context) {
	var pageInfo request.UserPageInfo
	if errs := ctx.ShouldBindJSON(&pageInfo); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if list, total, err := service.GetGeneralInfoList(pageInfo); err != nil {
		g.TENANCY_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}

// SetUserAuthority 设置用户权限
func SetUserAuthority(ctx *gin.Context) {
	var sua request.SetUserAuth
	if errs := ctx.ShouldBindJSON(&sua); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.SetUserAuthority(sua.Id, sua.AuthorityId); err != nil {
		g.TENANCY_LOG.Error("修改失败", zap.Any("err", err))
		response.FailWithMessage("修改失败", ctx)
	} else {
		response.OkWithMessage("修改成功", ctx)
	}
}

// DeleteUser 删除用户
func DeleteUser(ctx *gin.Context) {
	var reqId request.GetById
	if errs := ctx.ShouldBindJSON(&reqId); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	jwtId := multi.GetUserId(ctx)
	if jwtId == reqId.Id {
		response.FailWithMessage("删除失败, 自杀失败", ctx)
		return
	}
	if err := service.DeleteUser(reqId.Id); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}

// SetUserInfo 设置用户信息
func SetUserInfo(ctx *gin.Context) {
	userId := ctx.Param("user_id")
	user, err := service.FindUserById(userId)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	if user.IsAdmin() {
		var admin model.AdminInfo
		_ = ctx.ShouldBindJSON(&admin)
		if _, err := service.SetUserAdminInfo(admin, user.AdminInfo.ID, userId); err != nil {
			g.TENANCY_LOG.Error("设置失败", zap.Any("err", err))
			response.FailWithMessage("设置失败", ctx)
		} else {
			response.OkWithMessage("设置成功", ctx)
		}
	} else if user.IsTenancy() {
		var tenancy model.TenancyInfo
		_ = ctx.ShouldBindJSON(&tenancy)
		tenancy.ID = user.TenancyInfo.ID
		if _, err := service.SetUserTenancyInfo(tenancy, user.TenancyInfo.ID, userId); err != nil {
			g.TENANCY_LOG.Error("设置失败", zap.Any("err", err))
			response.FailWithMessage("设置失败", ctx)
		} else {
			response.OkWithMessage("设置成功", ctx)
		}
		//TODO::不能修改用户信息
		// } else if user.IsGeneral() {
		// 	var general model.SysGeneralInfo
		// 	_ = ctx.ShouldBindJSON(&general)
		// 	general.ID = user.GeneralInfo.ID
		// 	if _, err := service.SetUserGeneralInfo(general, user.GeneralInfo.ID, userId); err != nil {
		// 		g.TENANCY_LOG.Error("设置失败", zap.Any("err", err))
		// 		response.FailWithMessage("设置失败", ctx)
		// 	} else {
		// 		response.OkWithMessage("设置成功", ctx)
		// 	}
	} else {
		g.TENANCY_LOG.Error("未知角色", zap.Any("err", user.AuthorityType()))
		response.FailWithMessage("未知角色", ctx)
	}
}
