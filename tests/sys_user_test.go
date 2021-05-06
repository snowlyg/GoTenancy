package tests

import (
	"testing"

	"github.com/iris-contrib/httpexpect/v2"
	"github.com/kataras/iris/v12/httptest"
)

func TestLogin(t *testing.T) {
	e := baseTester(t)
	obj := e.POST("/public/login").
		WithJSON(map[string]interface{}{"username": "admin", "password": "123456"}).
		Expect().Status(httptest.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("登录成功")
	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("user", "token", "expiresAt")
	data.Value("token").NotNull()
	data.Value("expiresAt").NotNull()

	token := data.Value("token").String().Raw()
	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})

	obj = auth.GET("/v1/user/logout").
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("退出登录")
}

func TestLoginWithErrorUsername(t *testing.T) {
	e := baseTester(t)
	obj := e.POST("/public/login").
		WithJSON(map[string]interface{}{"username": "error_username", "password": "123456"}).
		Expect().Status(httptest.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(7)
	obj.Value("msg").String().Equal("用户名不存在或者密码错误")
	obj.Value("data").Object().Empty()
}

func TestLoginWithErrorPassword(t *testing.T) {
	e := baseTester(t)
	obj := e.POST("/public/login").
		WithJSON(map[string]interface{}{"username": "admin", "password": "error_pwd"}).
		Expect().Status(httptest.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(7)
	obj.Value("msg").String().Equal("用户名不存在或者密码错误")
	obj.Value("data").Object().Empty()
}

func TestLoginWithErrorUsernameAndPassword(t *testing.T) {
	e := baseTester(t)
	obj := e.POST("/public/login").
		WithJSON(map[string]interface{}{"username": "error_username", "password": "error_pwd"}).
		Expect().Status(httptest.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(7)
	obj.Value("msg").String().Equal("用户名不存在或者密码错误")
	obj.Value("data").Object().Empty()
}
