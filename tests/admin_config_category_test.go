package tests

import (
	"fmt"
	"net/http"
	"testing"
)

func TestConfigCategoryList(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.GET("v1/admin/configCategory/getConfigCategoryList").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize")
	data.Value("pageSize").Number().Equal(0)
	data.Value("page").Number().Equal(0)
	data.Value("total").Number().Ge(0)

	list := data.Value("list").Array()
	list.Length().Ge(0)
	first := list.First().Object()
	first.Keys().ContainsOnly("id", "name", "icon", "sort", "info", "key", "status", "createdAt", "updatedAt")
	first.Value("id").Number().Ge(0)

}

func TestCreateConfigCategoryMap(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.GET("v1/admin/configCategory/getCreateConfigCategoryMap").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")
}

func TestConfigCategoryProcess(t *testing.T) {
	data := map[string]interface{}{
		"name":   "箱包",
		"sort":   0,
		"key":    "xiangbao",
		"icon":   "",
		"info":   "箱包",
		"status": 1,
	}
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/configCategory/createConfigCategory").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("创建成功")

	configCategory := obj.Value("data").Object()
	configCategory.Value("id").Number().Ge(0)
	configCategory.Value("name").String().Equal(data["name"].(string))
	configCategory.Value("key").String().Equal(data["key"].(string))
	configCategory.Value("icon").String().Equal(data["icon"].(string))
	configCategory.Value("sort").Number().Equal(data["sort"].(int))
	configCategory.Value("info").String().Equal(data["info"].(string))
	configCategory.Value("status").Number().Equal(data["status"].(int))
	configCategoryId := configCategory.Value("id").Number().Raw()

	// getUpdateConfigCategoryMap
	obj = auth.GET(fmt.Sprintf("v1/admin/configCategory/getUpdateConfigCategoryMap/%d", int(configCategoryId))).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	update := map[string]interface{}{
		"name":   "箱包1",
		"sort":   0,
		"key":    "xiangbao1sdfsdf",
		"icon":   "",
		"info":   "箱包1",
		"status": 2,
	}

	obj = auth.PUT(fmt.Sprintf("v1/admin/configCategory/updateConfigCategory/%d", int(configCategoryId))).
		WithJSON(update).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("更新成功")
	configCategory = obj.Value("data").Object()

	configCategory.Value("name").String().Equal(update["name"].(string))
	configCategory.Value("key").String().Equal(update["key"].(string))
	configCategory.Value("icon").String().Equal(update["icon"].(string))
	configCategory.Value("sort").Number().Equal(update["sort"].(int))
	configCategory.Value("info").String().Equal(update["info"].(string))
	configCategory.Value("status").Number().Equal(update["status"].(int))

	obj = auth.GET(fmt.Sprintf("v1/admin/configCategory/getConfigCategoryById/%d", int(configCategoryId))).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
	configCategory = obj.Value("data").Object()

	configCategory.Value("id").Number().Ge(0)
	configCategory.Value("name").String().Equal(update["name"].(string))
	configCategory.Value("key").String().Equal(update["key"].(string))
	configCategory.Value("icon").String().Equal(update["icon"].(string))
	configCategory.Value("sort").Number().Equal(update["sort"].(int))
	configCategory.Value("info").String().Equal(update["info"].(string))
	configCategory.Value("status").Number().Equal(update["status"].(int))

	// changeConfigCategoryStatus
	obj = auth.POST("v1/admin/configCategory/changeConfigCategoryStatus").
		WithJSON(map[string]interface{}{
			"id":     configCategoryId,
			"status": 2,
		}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("设置成功")

	// deleteConfigCategory
	obj = auth.DELETE(fmt.Sprintf("v1/admin/configCategory/deleteConfigCategory/%d", int(configCategoryId))).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("删除成功")

}

func TestConfigCategoryRegisterError(t *testing.T) {
	data := map[string]interface{}{
		"name":   "箱包",
		"sort":   0,
		"key":    "shop",
		"icon":   "",
		"info":   "箱包",
		"status": 2,
	}
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/configCategory/createConfigCategory").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("添加失败:KEY shop 已被使用")

}
