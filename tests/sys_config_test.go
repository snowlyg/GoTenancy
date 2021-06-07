package tests

import (
	"net/http"
	"testing"
)

func TestConfigList(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/config/getConfigList").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize")
	data.Value("pageSize").Number().Equal(10)
	data.Value("page").Number().Equal(1)
	data.Value("total").Number().Ge(0)

	list := data.Value("list").Array()
	list.Length().Ge(0)
	first := list.First().Object()
	first.Keys().ContainsOnly("id", "type", "name", "value", "createdAt", "updatedAt")
	first.Value("id").Number().Ge(0)

}

func TestConfigProcess(t *testing.T) {
	data := map[string]interface{}{
		"name":  "YJ3s1abt7MAfT6gWVKoDresdfsdf",
		"type":  "YJ3s1abt7MAfT6gWVKoDresdfsdf",
		"value": "tRE49zaf5NCm6PidFZoaFg3u4WCHDok7fxgL63yV0pF4AMsdfsdfsdfssa",
	}
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/config/createConfig").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("创建成功")

	config := obj.Value("data").Object()
	config.Value("id").Number().Ge(0)
	config.Value("type").String().Equal(data["type"].(string))
	config.Value("value").String().Equal(data["value"].(string))
	config.Value("name").String().Equal(data["name"].(string))
	configId := config.Value("id").Number().Raw()

	update := map[string]interface{}{
		"id":    configId,
		"name":  "YJ3s1abt7MAfT6gWVKoDjnhjsfsd",
		"type":  "YJ3s1abt7MAfT6gWVKoDjnhjsfsd",
		"value": "tRE49zaf5NCm6PidFZoaFg3u4WCHDok7fxgL63yV0pF4AMsdfbnfgh",
	}

	obj = auth.PUT("/v1/admin/config/updateConfig").
		WithJSON(update).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("更新成功")
	config = obj.Value("data").Object()

	config.Value("id").Number().Ge(0)
	config.Value("name").String().Equal(update["name"].(string))
	config.Value("type").String().Equal(update["type"].(string))
	config.Value("value").String().Equal(update["value"].(string))

	getByName := map[string]interface{}{
		"name": "YJ3s1abt7MAfT6gWVKoDjnhjsfsd",
		"type": "YJ3s1abt7MAfT6gWVKoDjnhjsfsd",
	}
	obj = auth.POST("/v1/admin/config/getConfigByName").
		WithJSON(getByName).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
	config = obj.Value("data").Object()

	config.Value("id").Number().Ge(0)
	config.Value("name").String().Equal(update["name"].(string))
	config.Value("type").String().Equal(update["type"].(string))
	config.Value("value").String().Equal(update["value"].(string))

	// setUserAuthority
	obj = auth.DELETE("/v1/admin/config/deleteConfig").
		WithJSON(map[string]interface{}{"id": configId}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("删除成功")

}

func TestConfigRegisterError(t *testing.T) {
	data := map[string]interface{}{
		"name":  "token",
		"type":  "wechat",
		"value": "tRE49zaf5NCm6PidFZoaFg3u4WCHDok7fxgL63yV0pF4AM",
	}
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/config/createConfig").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("添加失败:设置名称已经使用")

}
