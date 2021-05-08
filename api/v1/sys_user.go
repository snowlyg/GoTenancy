package v1

import (
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/middleware"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/go-tenancy/utils"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

// Login 用户登录
func Login(ctx iris.Context) {
	var L request.Login
	_ = ctx.ReadJSON(&L)

	if err := utils.Verify(L, utils.GetLoginVerify()); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	if store.Verify(L.CaptchaId, L.Captcha, true) || g.TENANCY_CONFIG.System.Env == "dev" {
		U := &model.SysUser{Username: L.Username, Password: L.Password}
		if err, user := service.Login(U); err != nil {
			g.TENANCY_LOG.Error("登陆失败! 用户名不存在或者密码错误", zap.Any("err", err))
			response.FailWithMessage("用户名不存在或者密码错误", ctx)
		} else {
			tokenNext(ctx, *user)
		}
	} else {
		response.FailWithMessage("验证码错误", ctx)
	}
}

// tokenNext 登录以后签发jwt
func tokenNext(ctx iris.Context, user model.SysUser) {
	claims := &multi.CustomClaims{
		ID:            strconv.FormatUint(uint64(user.ID), 10),
		Username:      user.Username,
		AuthorityId:   user.AuthorityId,
		AuthorityType: user.Authority.AuthorityType,
		LoginType:     multi.LoginTypeWeb,
		AuthType:      multi.AuthPwd,
		CreationDate:  time.Now().Local().Unix(),
		ExpiresIn:     multi.RedisSessionTimeoutWeb.Milliseconds(),
	}

	token, _, err := middleware.CreateToken(claims)
	if err != nil {
		g.TENANCY_LOG.Error("获取token失败", zap.Any("err", err))
		response.FailWithMessage("获取token失败", ctx)
		return
	}
	response.OkWithDetailed(response.LoginResponse{
		Token: token,
	}, "登录成功", ctx)
}

// Logout 退出登录
func Logout(ctx iris.Context) {
	token := multi.GetVerifiedToken(ctx)
	if token == nil {
		response.FailWithMessage("授权凭证为空", ctx)
		return
	}
	err := middleware.DelToken(string(token))
	if err != nil {
		g.TENANCY_LOG.Error("del token", zap.Any("err", err))
		response.FailWithMessage("退出失败", ctx)
		return
	}
	response.OkWithMessage("退出登录", ctx)
}

// Clean 清空 token
func Clean(ctx iris.Context) {
	waitUse := multi.Get(ctx)
	if waitUse == nil {
		response.FailWithMessage("清空TOKEN失败", ctx)
		return
	}
	err := middleware.CleanToken(waitUse.ID)
	if err != nil {
		g.TENANCY_LOG.Error("清空TOKEN失败", zap.Any("err", err))
		response.FailWithMessage("清空TOKEN失败", ctx)
		return
	}
	response.OkWithMessage("TOKEN清空", ctx)
}

// Register 用户注册账号
func Register(ctx iris.Context) {
	var R request.Register
	_ = ctx.ReadJSON(&R)
	if err := utils.Verify(R, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	user := &model.SysUser{Username: R.Username, Password: R.Password, AuthorityId: R.AuthorityId}
	err, userReturn := service.Register(*user)
	if err != nil {
		g.TENANCY_LOG.Error("注册失败", zap.Any("err", err))
		response.FailWithDetailed(response.SysUserResponse{User: userReturn}, "注册失败", ctx)
	} else {
		response.OkWithDetailed(response.SysUserResponse{User: userReturn}, "注册成功", ctx)
	}
}

// ChangePassword 用户修改密码
func ChangePassword(ctx iris.Context) {
	var user request.ChangePasswordStruct
	_ = ctx.ReadJSON(&user)
	if err := utils.Verify(user, utils.ChangePasswordVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	U := &model.SysUser{Username: user.Username, Password: user.Password}
	if err, _ := service.ChangePassword(U, user.NewPassword); err != nil {
		g.TENANCY_LOG.Error("修改失败", zap.Any("err", err))
		response.FailWithMessage("修改失败，原密码与当前账户不符", ctx)
	} else {
		response.OkWithMessage("修改成功", ctx)
	}
}

// GetUserList 分页获取用户列表
func GetUserList(ctx iris.Context) {
	var pageInfo request.PageInfo
	_ = ctx.ReadJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err, list, total := service.GetUserInfoList(pageInfo); err != nil {
		g.TENANCY_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
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
func SetUserAuthority(ctx iris.Context) {
	var sua request.SetUserAuth
	_ = ctx.ReadJSON(&sua)
	if UserVerifyErr := utils.Verify(sua, utils.SetUserAuthorityVerify); UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), ctx)
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
func DeleteUser(ctx iris.Context) {
	var reqId request.GetById
	_ = ctx.ReadJSON(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	jwtId := ctx.GetID()
	if jwtId == uint(reqId.Id) {
		response.FailWithMessage("删除失败, 自杀失败", ctx)
		return
	}
	if err := service.DeleteUser(reqId.Id); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}

// SetUserInfo 设置用户信息
func SetUserInfo(ctx iris.Context) {
	userId := ctx.Params().GetIntDefault("user_id", 0)
	err, user := service.FindUserById(userId)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	if user.IsAdmin() {
		var admin model.SysAdminInfo
		_ = ctx.ReadJSON(&admin)
		if err := utils.Verify(admin, utils.IdVerify); err != nil {
			response.FailWithMessage(err.Error(), ctx)
			return
		}
		if err, ReqUser := service.SetUserAdminInfo(admin, user.AdminInfo.ID > 0); err != nil {
			g.TENANCY_LOG.Error("设置失败", zap.Any("err", err))
			response.FailWithMessage("设置失败", ctx)
		} else {
			response.OkWithDetailed(iris.Map{"adminInfo": ReqUser}, "设置成功", ctx)
		}
	} else if user.IsTenancy() {
		var tenancy model.SysTenancyInfo
		_ = ctx.ReadJSON(&tenancy)
		if err := utils.Verify(tenancy, utils.IdVerify); err != nil {
			response.FailWithMessage(err.Error(), ctx)
			return
		}
		if err, ReqUser := service.SetUserTenancyInfo(tenancy, user.TenancyInfo.ID > 0); err != nil {
			g.TENANCY_LOG.Error("设置失败", zap.Any("err", err))
			response.FailWithMessage("设置失败", ctx)
		} else {
			response.OkWithDetailed(iris.Map{"tenancyInfo": ReqUser}, "设置成功", ctx)
		}
	} else if user.IsGeneral() {
		var general model.SysGeneralInfo
		_ = ctx.ReadJSON(&general)
		if err := utils.Verify(general, utils.IdVerify); err != nil {
			response.FailWithMessage(err.Error(), ctx)
			return
		}
		if err, ReqUser := service.SetUserGeneralInfo(general, user.GeneralInfo.ID > 0); err != nil {
			g.TENANCY_LOG.Error("设置失败", zap.Any("err", err))
			response.FailWithMessage("设置失败", ctx)
		} else {
			response.OkWithDetailed(iris.Map{"generalInfo": ReqUser}, "设置成功", ctx)
		}
	} else {
		g.TENANCY_LOG.Error("未知角色", zap.Any("err", user.AuthorityType()))
		response.FailWithMessage("未知角色", ctx)
	}
}

// GetClaims returns the current authorized client claims.
func GetClaims(ctx iris.Context) *multi.CustomClaims {
	waitUse := multi.Get(ctx)
	if waitUse == nil {
		g.TENANCY_LOG.Error("从Context中获取用户ID失败, 请检查路由是否使用multi中间件")
	}
	return waitUse
}

// getUserID 从Context中获取用户ID
func getUserID(ctx iris.Context) string {
	if claims := GetClaims(ctx); claims == nil {
		g.TENANCY_LOG.Error("从Context中获取用户ID失败, 请检查路由是否使用multi中间件")
		return ""
	} else {
		return claims.ID
	}
}

// getUserAuthorityId 从Context中获取用户角色id
func getUserAuthorityId(ctx iris.Context) string {
	if claims := GetClaims(ctx); claims == nil {
		g.TENANCY_LOG.Error("从Context中获取用户UUID失败, 请检查路由是否使用multi中间件")
		return ""
	} else {
		return claims.AuthorityId
	}
}
