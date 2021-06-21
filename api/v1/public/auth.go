package public

//ClientLogin 后台登录
func ClientLogin(ctx *gin.Context) {
	var L request.Login
	if errs := ctx.ShouldBindJSON(&L); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if store.Verify(L.CaptchaId, L.Captcha, true) || g.TENANCY_CONFIG.System.Env == "dev" {
		U := &model.SysUser{Username: L.Username, Password: L.Password}
		if loginResponse, err := service.Login(U, multi.TenancyAuthority); err != nil {
			g.TENANCY_LOG.Error("登陆失败!", zap.Any("err", err))
			response.FailWithMessage(err.Error(), ctx)
		} else {
			response.OkWithDetailed(loginResponse, "登录成功", ctx)
		}
	} else {
		response.FailWithMessage("验证码错误", ctx)
	}
}

// Logout 退出登录
func Logout(ctx *gin.Context) {
	token := multi.GetVerifiedToken(ctx)
	if token == nil {
		response.FailWithMessage("授权凭证为空", ctx)
		return
	}
	err := service.DelToken(string(token))
	if err != nil {
		g.TENANCY_LOG.Error("del token", zap.Any("err", err))
		response.FailWithMessage("退出失败", ctx)
		return
	}
	response.OkWithMessage("退出登录", ctx)
}

// Clean 清空 token
func Clean(ctx *gin.Context) {
	waitUse := multi.Get(ctx)
	if waitUse == nil {
		response.FailWithMessage("清空TOKEN失败", ctx)
		return
	}
	err := service.CleanToken(waitUse.ID)
	if err != nil {
		g.TENANCY_LOG.Error("清空TOKEN失败", zap.Any("err", err))
		response.FailWithMessage("清空TOKEN失败", ctx)
		return
	}
	response.OkWithMessage("TOKEN清空", ctx)
}