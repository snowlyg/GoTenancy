package tests

import (
	"testing"

	"github.com/kataras/iris/v12/httptest"
)

func TestLoginWithErrorUsername(t *testing.T) {
	e := baseTester(t)
	obj := e.POST("/v1/public/login").
		WithJSON(map[string]interface{}{"username": "error_username", "password": "123456", "authorityType": 1}).
		Expect().Status(httptest.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(7)
	obj.Value("msg").String().Equal("用户名不存在或者密码错误")
	obj.Value("data").Object().Empty()
}

func TestLoginWithErrorPassword(t *testing.T) {
	e := baseTester(t)
	obj := e.POST("/v1/public/login").
		WithJSON(map[string]interface{}{"username": "admin", "password": "error_pwd", "authorityType": 1}).
		Expect().Status(httptest.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(7)
	obj.Value("msg").String().Equal("用户名不存在或者密码错误")
	obj.Value("data").Object().Empty()
}

func TestLoginWithErrorUsernameAndPassword(t *testing.T) {
	e := baseTester(t)
	obj := e.POST("/v1/public/login").
		WithJSON(map[string]interface{}{"username": "error_username", "password": "error_pwd", "authorityType": 1}).
		Expect().Status(httptest.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(7)
	obj.Value("msg").String().Equal("用户名不存在或者密码错误")
	obj.Value("data").Object().Empty()
}

func TestLoginWithErrorAuthorityType(t *testing.T) {
	e := baseTester(t)
	obj := e.POST("/v1/public/login").
		WithJSON(map[string]interface{}{"username": "error_username", "password": "error_pwd", "authorityType": 3}).
		Expect().Status(httptest.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(7)
	obj.Value("msg").String().Equal("用户名不存在或者密码错误")
	obj.Value("data").Object().Empty()
}

func TestLoginWithEmptyAuthorityType(t *testing.T) {
	e := baseTester(t)
	obj := e.POST("/v1/public/login").
		WithJSON(map[string]interface{}{"username": "admin", "password": "123456", "authorityType": 0}).
		Expect().Status(httptest.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(7)
	obj.Value("msg").String().Equal("AuthorityType值不能为空")
	obj.Value("data").Object().Empty()
}
