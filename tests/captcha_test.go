package tests

import (
	"testing"

	"github.com/kataras/iris/v12/httptest"
)

func TestCaptcha(t *testing.T) {
	e := baseTester(t)
	obj := e.POST("/v1/public/captcha").
		Expect().Status(httptest.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("验证码获取成功")
	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("captchaId", "picPath")

	data.Value("captchaId").String().NotEmpty()
	data.Value("picPath").String().NotEmpty()
}
