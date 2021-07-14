package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/g"
)

func TestExpressList(t *testing.T) {
	params := []param{
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "name": ""}, length: 34},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "name": "汤氏物流"}, length: 1},
	}
	for _, param := range params {
		expressList(t, param.args, param.length)
	}

}

func expressList(t *testing.T, params map[string]interface{}, length int) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/express/getExpressList").
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

	list := data.Value("list").Array()
	list.Length().Ge(0)
	first := list.First().Object()
	first.Keys().ContainsOnly("id", "name", "status", "code", "sort", "createdAt", "updatedAt")
}

func TestExpressProcess(t *testing.T) {
	data := map[string]interface{}{
		"name":   "sdfsdfs34234",
		"code":   "sdfsdfs34234",
		"sort":   1,
		"status": 1,
	}
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/express/createExpress").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("创建成功")

	express := obj.Value("data").Object()
	express.Value("id").Number().Ge(0)
	express.Value("name").String().Equal(data["name"].(string))
	express.Value("status").Number().Equal(data["status"].(int))
	express.Value("code").String().Equal(data["code"].(string))
	express.Value("sort").Number().Equal(data["sort"].(int))
	expressId := express.Value("id").Number().Raw()
	if expressId > 0 {

		update := map[string]interface{}{
			"name":   "sdfsdfs213213",
			"code":   "sdfsdfs213213",
			"sort":   1,
			"status": 1,
		}

		obj = auth.PUT(fmt.Sprintf("v1/admin/express/updateExpress/%d", int(expressId))).
			WithJSON(update).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("更新成功")
		express = obj.Value("data").Object()

		express.Value("id").Number().Ge(0)
		express.Value("name").String().Equal(update["name"].(string))
		express.Value("status").Number().Equal(update["status"].(int))
		express.Value("code").String().Equal(update["code"].(string))
		express.Value("sort").Number().Equal(update["sort"].(int))

		obj = auth.GET(fmt.Sprintf("v1/admin/express/getExpressById/%d", int(expressId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("操作成功")
		express = obj.Value("data").Object()

		express.Value("id").Number().Ge(0)
		express.Value("name").String().Equal(update["name"].(string))
		express.Value("status").Number().Equal(update["status"].(int))
		express.Value("code").String().Equal(update["code"].(string))
		express.Value("sort").Number().Equal(update["sort"].(int))

		obj = auth.POST("v1/admin/express/changeExpressStatus").
			WithJSON(map[string]interface{}{
				"id":     expressId,
				"status": g.StatusTrue,
			}).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("设置成功")

		obj = auth.GET("v1/admin/express/getCreateExpressMap").
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		obj = auth.GET(fmt.Sprintf("v1/admin/express/getUpdateExpressMap/%d", int(expressId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		// deleteExpress
		obj = auth.DELETE(fmt.Sprintf("v1/admin/express/deleteExpress/%d", int(expressId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("删除成功")
	}
}
