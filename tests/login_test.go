package tests

import (
	"net/http"
	"testing"
)

func TestLoginWithErrorUsername(t *testing.T) {
	e := baseTester(t)
	obj := e.POST("v1/public/login").
		WithJSON(map[string]interface{}{"username": "error_username", "password": "123456", "authorityType": 1}).
		Expect().Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("用户名或者密码错误")
	obj.Value("data").Object().Empty()
}

func TestLoginWithErrorPassword(t *testing.T) {
	e := baseTester(t)
	obj := e.POST("v1/public/login").
		WithJSON(map[string]interface{}{"username": "admin", "password": "error_pwd", "authorityType": 1}).
		Expect().Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("用户名或者密码错误")
	obj.Value("data").Object().Empty()
}

func TestLoginWithErrorUsernameAndPassword(t *testing.T) {
	e := baseTester(t)
	obj := e.POST("v1/public/login").
		WithJSON(map[string]interface{}{"username": "error_username", "password": "error_pwd", "authorityType": 1}).
		Expect().Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("用户名或者密码错误")
	obj.Value("data").Object().Empty()
}

func TestLoginWithErrorAuthorityType(t *testing.T) {
	e := baseTester(t)
	obj := e.POST("v1/public/login").
		WithJSON(map[string]interface{}{"username": "error_username", "password": "error_pwd", "authorityType": 3}).
		Expect().Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("用户名或者密码错误")
	obj.Value("data").Object().Empty()
}

func TestLoginWithEmptyAuthorityType(t *testing.T) {
	e := baseTester(t)
	obj := e.POST("v1/public/login").
		WithJSON(map[string]interface{}{"username": "admin", "password": "123456", "authorityType": 0}).
		Expect().Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("Key: 'Login.AuthorityType' Error:Field validation for 'AuthorityType' failed on the 'required' tag")
	obj.Value("data").Object().Empty()
}
