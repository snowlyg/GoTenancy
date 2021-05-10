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
	obj := auth.POST("/v1/user/getAdminList").
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
	data.Value("list").Array().Length().Ge(0)

	user := data.Value("list").Array().First().Object()
	// user.Keys().ContainsOnly("ID")
	user.Value("ID").Number().Ge(0)

	baseLogOut(auth)
}

func TestAdminUserProcess(t *testing.T) {
	auth := baseWithLoginTester(t)
	obj := auth.POST("/v1/user/register").
		WithJSON(map[string]interface{}{"username": "chindeo", "password": "123456", "authority_id": multi.AdminAuthority}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("注册成功")

	user := obj.Value("data").Object().Value("user").Object()
	// user.Keys().ContainsOnly("ID")
	user.Value("ID").Number().Ge(0)
	user.Value("userName").String().Equal("chindeo")
	user.Value("authorityId").String().Equal(source.AdminAuthorityId)
	userId := user.Value("ID").Number().Raw()

	// changePassword error
	obj = auth.POST("/v1/user/changePassword").
		WithJSON(map[string]interface{}{"username": "chindeo", "password": "123456", "new_password": "456789"}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("修改成功")

	// changePassword error
	obj = auth.POST("/v1/user/changePassword").
		WithJSON(map[string]interface{}{"username": "chindeo", "password": "123456", "new_password": "456789"}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(7)
	obj.Value("msg").String().Equal("修改失败，原密码与当前账户不符")

	// setUserAuthority
	obj = auth.POST("/v1/user/setUserAuthority").
		WithJSON(map[string]interface{}{"id": userId, "authority_id": multi.AdminAuthority}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("修改成功")

	// setAdminInfo
	obj = auth.PUT(fmt.Sprintf("/v1/user/setAdminInfo/%f", userId)).
		WithJSON(map[string]interface{}{"id": 1, "email": "admin@master.com", "phone": "13800138001"}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("设置成功")

	// setUserAuthority
	obj = auth.DELETE("/v1/user/deleteUser").
		WithJSON(map[string]interface{}{"id": userId}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("删除成功")

	baseLogOut(auth)
}

func TestAdminUserRegisterError(t *testing.T) {
	auth := baseWithLoginTester(t)
	obj := auth.POST("/v1/user/register").
		WithJSON(map[string]interface{}{"username": "admin", "password": "123456", "authority_id": multi.AdminAuthority}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(7)
	obj.Value("msg").String().Equal("注册失败")

	baseLogOut(auth)
}
