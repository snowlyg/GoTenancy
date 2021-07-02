package tests

import (
	"net/http"
	"testing"
)

func TestShippingTemplateList(t *testing.T) {
	params := []param{
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "name": ""}, length: 4},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "name": "邮费"}, length: 2},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "name": "陕西"}, length: 1},
	}
	for _, param := range params {
		shippingTemplate(t, param.args, param.length)
	}
}

func shippingTemplate(t *testing.T, params map[string]interface{}, length int) {
	auth := tenancyWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/merchant/shippingTemplate/getShippingTemplateList").
		WithJSON(params).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize")
	data.Value("pageSize").Number().Equal(10)
	data.Value("page").Number().Equal(1)
	data.Value("total").Number().Equal(length)
	if length > 0 {
		list := data.Value("list").Array()
		list.Length().Ge(0)
		first := list.First().Object()
		first.Keys().ContainsOnly(
			"id",
			"createdAt",
			"updatedAt",
			"name",
			"type",
			"appoint",
			"undelivery",
			"isDefault",
			"sort",
		)
		first.Value("id").Number().Ge(0)
	}
}

func TestShippingTemplateSelect(t *testing.T) {
	auth := tenancyWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.GET("v1/merchant/shippingTemplate/getShippingTemplateSelect").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")
	obj.Value("data").Array().First().Object().Keys().ContainsOnly(
		"id",
		"name",
	)
}
