package tests

import (
	"testing"

	"github.com/kataras/iris/v12/httptest"
)

func TestMiniList(t *testing.T) {
	auth := baseWithLoginTester(t)
	obj := auth.POST("/v1/admin/mini/getMiniList").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("获取成功")

	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize")
	data.Value("pageSize").Number().Equal(10)
	data.Value("page").Number().Equal(1)
	data.Value("total").Number().Ge(0)

	list := data.Value("list").Array()
	list.Length().Ge(0)
	first := list.First().Object()
	first.Keys().ContainsOnly("id", "uuid", "name", "appId", "appSecret", "remark", "createdAt", "updatedAt")
	first.Value("id").Number().Ge(0)

	baseLogOut(auth)
}

func TestMiniProcess(t *testing.T) {
	data := map[string]interface{}{
		"name":      "中德澳线上点餐商城",
		"appId":     "YJ3s1abt7MAfT6gWVKoDresdfsdf",
		"appSecret": "tRE49zaf5NCm6PidFZoaFg3u4WCHDok7fxgL63yV0pF4AMsdfsdfsdfssa",
		"remark":    "中德澳线上点餐商城",
	}
	auth := baseWithLoginTester(t)
	obj := auth.POST("/v1/admin/mini/createMini").
		WithJSON(data).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("创建成功")

	mini := obj.Value("data").Object()
	mini.Value("id").Number().Ge(0)
	mini.Value("uuid").String().NotEmpty()
	mini.Value("name").String().Equal(data["name"].(string))
	mini.Value("appId").String().Equal(data["appId"].(string))
	mini.Value("appSecret").String().Equal(data["appSecret"].(string))
	mini.Value("remark").String().Equal(data["remark"].(string))
	miniId := mini.Value("id").Number().Raw()

	update := map[string]interface{}{
		"id":        miniId,
		"name":      "中德澳线上点餐商城1",
		"appId":     "YJ3s1abt7MAfT6gWVKoDjnhjsfsd",
		"appSecret": "tRE49zaf5NCm6PidFZoaFg3u4WCHDok7fxgL63yV0pF4AMsdfbnfgh",
		"remark":    "中德澳线上点餐商城1",
	}

	obj = auth.PUT("/v1/admin/mini/updateMini").
		WithJSON(update).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("更新成功")
	mini = obj.Value("data").Object()

	mini.Value("id").Number().Ge(0)
	mini.Value("uuid").String().NotEmpty()
	mini.Value("name").String().Equal(update["name"].(string))
	mini.Value("appId").String().Equal(update["appId"].(string))
	mini.Value("appSecret").String().Equal(update["appSecret"].(string))
	mini.Value("remark").String().Equal(update["remark"].(string))

	// setUserAuthority
	obj = auth.DELETE("/v1/admin/mini/deleteMini").
		WithJSON(map[string]interface{}{"id": miniId}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("删除成功")

	baseLogOut(auth)
}

func TestMiniRegisterError(t *testing.T) {
	data := map[string]interface{}{
		"name":      "中德澳上线护理商城",
		"appId":     "YJ3s1abt7MAfT6gWVKoD",
		"appSecret": "tRE49zaf5NCm6PidFZoaFg3u4WCHDok7fxgL63yV0pF4AM",
		"remark":    "中德澳上线护理商城",
	}
	auth := baseWithLoginTester(t)
	obj := auth.POST("/v1/admin/mini/createMini").
		WithJSON(data).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(7)
	obj.Value("msg").String().Equal("创建失败")

	baseLogOut(auth)
}
