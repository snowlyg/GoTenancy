package tests

import (
	"fmt"
	"testing"

	"github.com/kataras/iris/v12/httptest"
	"github.com/snowlyg/go-tenancy/source"
	"github.com/snowlyg/multi"
)

func TestAdminUserList(t *testing.T) {
	auth := baseWithLoginTester(t)
	obj := auth.POST("/v1/admin/user/getAdminList").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("获取成功")

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

	baseLogOut(auth)
}

func TestAdminUserProcess(t *testing.T) {
	auth := baseWithLoginTester(t)
	obj := auth.POST("/v1/admin/user/register").
		WithJSON(map[string]interface{}{"username": "chindeo", "password": "123456", "authorityId": source.AdminAuthorityId, "authorityType": multi.AdminAuthority}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("注册成功")

	user := obj.Value("data").Object()
	user.Value("id").Number().Ge(0)
	user.Value("userName").String().Equal("chindeo")
	user.Value("authorityId").String().Equal(source.AdminAuthorityId)
	userId := user.Value("id").Number().Raw()

	// changePassword error
	obj = auth.POST("/v1/admin/user/changePassword").
		WithJSON(map[string]interface{}{"username": "chindeo", "password": "123456", "newPassword": "456789", "authorityType": multi.AdminAuthority}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("修改成功")

	// changePassword error
	obj = auth.POST("/v1/admin/user/changePassword").
		WithJSON(map[string]interface{}{"username": "chindeo", "password": "123456", "newPassword": "456789", "authorityType": multi.AdminAuthority}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(4000)
	obj.Value("msg").String().Equal("修改失败，原密码与当前账户不符")

	// changePassword error
	obj = auth.POST("/v1/admin/user/changePassword").
		WithJSON(map[string]interface{}{"username": "chindeo", "password": "123456", "newPassword": "456789", "authorityType": 0}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(4000)
	obj.Value("msg").String().Equal("Key: 'ChangePasswordStruct.AuthorityType' Error:Field validation for 'AuthorityType' failed on the 'required' tag")

	// setUserAuthority
	obj = auth.POST("/v1/admin/user/setUserAuthority").
		WithJSON(map[string]interface{}{"id": userId, "authorityId": source.AdminAuthorityId}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("修改成功")

	// setAdminInfo
	obj = auth.PUT(fmt.Sprintf("/v1/admin/user/setUserInfo/%d", int(userId))).
		WithJSON(map[string]interface{}{"email": "admin@master.com", "phone": "13800138001"}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("设置成功")

	// setUserAuthority
	obj = auth.DELETE("/v1/admin/user/deleteUser").
		WithJSON(map[string]interface{}{"id": userId}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("删除成功")

	baseLogOut(auth)
}

func TestAdminUserRegisterError(t *testing.T) {
	auth := baseWithLoginTester(t)
	obj := auth.POST("/v1/admin/user/register").
		WithJSON(map[string]interface{}{"username": "admin", "password": "123456", "authorityId": source.AdminAuthorityId, "authorityType": multi.AdminAuthority}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(4000)
	obj.Value("msg").String().Equal("用户名已注册")

	baseLogOut(auth)
}

func TestAdminUserRegisterAuthorityTypeEmpty(t *testing.T) {
	auth := baseWithLoginTester(t)
	obj := auth.POST("/v1/admin/user/register").
		WithJSON(map[string]interface{}{"username": "admin", "password": "123456", "authorityId": source.AdminAuthorityId, "authorityType": 0}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(4000)
	obj.Value("msg").String().Equal("Key: 'Register.AuthorityType' Error:Field validation for 'AuthorityType' failed on the 'required' tag")

	baseLogOut(auth)
}

func TestAdminUserRegisterAuthorityIdEmpty(t *testing.T) {
	auth := baseWithLoginTester(t)
	obj := auth.POST("/v1/admin/user/register").
		WithJSON(map[string]interface{}{"username": "admin_authrity_id_empty", "password": "123456", "authorityId": "", "authorityType": multi.AdminAuthority}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(4000)
	obj.Value("msg").String().Equal("Key: 'Register.AuthorityId' Error:Field validation for 'AuthorityId' failed on the 'required' tag")

	baseLogOut(auth)
}
