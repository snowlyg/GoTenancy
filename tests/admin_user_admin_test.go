package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/source"
)

func TestAdminUserList(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/user/getAdminList").
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

	list := data.Value("list").Array()
	list.Length().Ge(0)
	first := list.First().Object()
	first.Keys().ContainsOnly("id", "userName", "email", "phone", "nickName", "headerImg", "authorityName", "authorityType", "authorityId", "defaultRouter", "createdAt", "updatedAt")
	first.Value("id").Number().Ge(0)
}

func TestAdminLoginUser(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	// changePassword success
	obj := auth.POST("v1/admin/user/changePassword").
		WithJSON(map[string]interface{}{"username": "admin", "password": "123456", "newPassword": "456789", "confirmPassword": "456789"}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("修改成功")

	// changePassword error
	obj = auth.POST("v1/admin/user/changePassword").
		WithJSON(map[string]interface{}{"username": "admin", "password": "123456", "newPassword": "456789", "confirmPassword": "456789"}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("修改失败，原密码与当前账户不符")

	// changeProfile success
	obj = auth.POST("v1/admin/user/changeProfile").
		WithJSON(map[string]interface{}{"nickName": "admin", "phone": "123456"}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("修改成功")

	// changeProfile success
	obj = auth.POST("v1/admin/user/changeProfile").
		WithJSON(map[string]interface{}{"nickName": "超级管理员", "phone": "123456"}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("修改成功")

	// changePassword success
	obj = auth.POST("v1/admin/user/changePassword").
		WithJSON(map[string]interface{}{"username": "admin", "password": "456789", "newPassword": "123456", "confirmPassword": "123456"}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("修改成功")
}

func TestAdminUserProcess(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/user/registerAdmin").
		WithJSON(map[string]interface{}{"username": "chindeo", "password": "123456", "authorityId": source.AdminAuthorityId}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("注册成功")

	user := obj.Value("data").Object()
	user.Value("id").Number().Ge(0)
	user.Value("userName").String().Equal("chindeo")
	user.Value("authorityId").String().Equal(source.AdminAuthorityId)
	userId := user.Value("id").Number().Raw()
	if userId > 0 {

		// setUserAuthority
		obj = auth.POST("v1/admin/user/setUserAuthority").
			WithJSON(map[string]interface{}{"id": userId, "authorityId": source.AdminAuthorityId}).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("修改成功")

		// setAdminInfo
		obj = auth.PUT(fmt.Sprintf("v1/admin/user/setUserInfo/%d", int(userId))).
			WithJSON(map[string]interface{}{"email": "admin@admin.com", "phone": "13800138001"}).
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

func TestAdminUserRegisterError(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/user/registerAdmin").
		WithJSON(map[string]interface{}{"username": "admin", "password": "123456", "authorityId": source.AdminAuthorityId}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("用户名已注册")

}

func TestAdminUserRegisterAuthorityIdEmpty(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/user/registerAdmin").
		WithJSON(map[string]interface{}{"username": "admin_authrity_id_empty", "password": "123456", "authorityId": ""}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("Key: 'Register.AuthorityId' Error:Field validation for 'AuthorityId' failed on the 'required' tag")

}
