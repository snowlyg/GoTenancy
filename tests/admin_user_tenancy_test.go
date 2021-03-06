package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/source"
	"github.com/snowlyg/multi"
)

func TestTenancyUserList(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/user/getTenancyList").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize")
	data.Value("pageSize").Number().Equal(10)
	data.Value("page").Number().Equal(1)
	data.Value("total").Number().Equal(1)
	data.Value("list").Array().Length().Ge(0)

	list := data.Value("list").Array()
	list.Length().Ge(0)
	first := list.First().Object()
	first.Keys().ContainsOnly("id", "userName", "email", "phone", "nickName", "headerImg", "authorityName", "authorityType", "authorityId", "tenancyName", "tenancyId", "defaultRouter", "createdAt", "updatedAt")
	first.Value("id").Number().Ge(0)
}

func TestTenancyUserProcess(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/user/registerTenancy").
		WithJSON(map[string]interface{}{"username": "admin", "password": "123456", "authorityId": source.TenancyAuthorityId, "authorityType": multi.TenancyAuthority}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("注册成功")

	user := obj.Value("data").Object()
	user.Value("id").Number().Ge(0)
	user.Value("userName").String().Equal("admin")
	user.Value("authorityId").String().Equal(source.TenancyAuthorityId)
	userId := user.Value("id").Number().Raw()
	if userId > 0 {
		// setUserAuthority
		obj = auth.POST("v1/admin/user/setUserAuthority").
			WithJSON(map[string]interface{}{"id": userId, "authorityId": source.TenancyAuthorityId}).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("修改成功")

		// setTenancyInfo
		obj = auth.PUT(fmt.Sprintf("v1/admin/user/setUserInfo/%d", int(userId))).
			WithJSON(map[string]interface{}{"email": "tenancy@master.com", "phone": "13800138001"}).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("设置成功")

		// setUserAuthority
		obj = auth.DELETE("v1/admin/user/deleteUser").
			WithJSON(map[string]interface{}{"id": userId}).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("删除成功")
	}

}

func TestTenancyUserRegisterError(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/user/registerTenancy").
		WithJSON(map[string]interface{}{"username": "a303176530", "password": "123456", "authorityId": source.TenancyAuthorityId, "authorityType": multi.TenancyAuthority}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("用户名已注册")

}

func TestTenancyUserRegisterAuthorityIdEmpty(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/user/registerTenancy").
		WithJSON(map[string]interface{}{"username": "chindeo_tenancy", "password": "123456", "authorityId": "", "authorityType": multi.TenancyAuthority}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("Key: 'Register.AuthorityId' Error:Field validation for 'AuthorityId' failed on the 'required' tag")

}
