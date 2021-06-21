package public

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/response"
	"go.uber.org/zap"
)

var store = base64Captcha.DefaultMemStore

// Captcha 生成验证码
func Captcha(ctx *gin.Context) {
	//字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(g.TENANCY_CONFIG.Captcha.ImgHeight, g.TENANCY_CONFIG.Captcha.ImgWidth, g.TENANCY_CONFIG.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		g.TENANCY_LOG.Error("验证码获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(response.SysCaptchaResponse{
			CaptchaId: id,
			PicPath:   b64s,
		}, "获取成功", ctx)
	}
}
