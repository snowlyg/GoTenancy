package tests

import (
	"net/http"
	"testing"
)

func TestCaptcha(t *testing.T) {
	e := baseTester(t)
	obj := e.POST("/v1/public/captcha").
		Expect().Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("验证码获取成功")
	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("captchaId", "picPath")

	data.Value("captchaId").String().NotEmpty()
	data.Value("picPath").String().NotEmpty()
}
