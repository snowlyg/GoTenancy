package tests

import (
	"testing"

	"github.com/kataras/iris/v12/httptest"
)

func TestGeneralUserList(t *testing.T) {
	auth := baseWithLoginTester(t)
	obj := auth.POST("/v1/admin/user/getGeneralList").
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

	list := data.Value("list").Array()
	list.Length().Ge(0)
	first := list.First().Object()
	first.Keys().ContainsOnly("id", "userName", "email", "phone", "nickName", "avatarUrl", "authorityName", "authorityType", "authorityId", "sex", "subscribe", "openId", "unionId", "country", "provice", "city", "idCard", "isAuth", "realName", "birthday", "createdAt", "updatedAt")
	first.Value("id").Number().Ge(0)

	baseLogOut(auth)
}
