package tests

import (
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/g"
)

func TestMenu(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.GET("v1/admin/menu/getMenu").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Array()
	data.Length().Ge(0)
	first := data.First().Object()
	first.Keys().ContainsOnly(
		"authoritys",
		"children",
		"createdAt",
		"hidden",
		"icon",
		"id",
		"is_menu",
		"is_tenancy",
		"menu_id",
		"menu_name",
		"params",
		"path",
		"pid",
		"route",
		"sort",
		"updatedAt",
	)
	first.Value("id").Number().Ge(0)

}
func TestBaseMenu(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/menu/getBaseMenuTree").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Array()
	data.Length().Ge(0)
	first := data.First().Object()
	first.Keys().ContainsOnly(
		"authoritys",
		"id",
		"createdAt",
		"icon",
		"menu_name",
		"route",
		"is_tenancy",
		"is_menu",
		"path",
		"hidden",
		"updatedAt",
		"pid",
		"params",
		"sort",
		"children",
	)
	first.Value("id").Number().Ge(0)

}

func TestMenuList(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/menu/getMenuList").
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
	first.Keys().ContainsOnly(
		"is_menu",
		"updatedAt",
		"pid",
		"menu_name",
		"is_tenancy",
		"id",
		"createdAt",
		"path",
		"children",
		"route",
		"params",
		"hidden",
		"icon",
		"sort",
		"authoritys",
	)
	first.Value("id").Number().Ge(0)

}

func TestMenuProcess(t *testing.T) {
	data := map[string]interface{}{
		"route":     "view/routerHolder.vue",
		"hidden":    g.StatusFalse,
		"menu_name": "test_menu_process",
		"parameters": []map[string]interface{}{
			{
				"type":  "query",
				"key":   "21321",
				"value": "1",
			},
			{
				"type":  "params",
				"key":   "12321",
				"value": "1",
			},
		},
		"pid":  0,
		"path": "21312331",
		"sort": 111,
	}
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/menu/addBaseMenu").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("添加成功")

	menu := obj.Value("data").Object()
	menu.Value("id").Number().Ge(0)
	menu.Value("pid").String().Equal(data["pid"].(string))
	menu.Value("menu_name").String().Equal(data["menu_name"].(string))
	menu.Value("path").String().Equal(data["path"].(string))
	menu.Value("hidden").Boolean().Equal(data["hidden"].(bool))
	menu.Value("route").String().Equal(data["route"].(string))
	menu.Value("sort").Number().Equal(data["sort"].(int))
	menu.Value("authoritys").Null()
	menu.Value("children").Null()
	menu.Value("parameters").Array().NotEmpty()
	meta := menu.Value("meta").Object()
	meta.Value("keepAlive").Boolean().Equal(data["meta"].(map[string]interface{})["keepAlive"].(bool))
	meta.Value("defaultMenu").Boolean().Equal(data["meta"].(map[string]interface{})["defaultMenu"].(bool))
	meta.Value("closeTab").Boolean().Equal(data["meta"].(map[string]interface{})["closeTab"].(bool))
	meta.Value("title").String().Equal(data["meta"].(map[string]interface{})["title"].(string))
	meta.Value("icon").String().Equal(data["meta"].(map[string]interface{})["icon"].(string))
	parameter := menu.Value("parameters").Array().First().Object()
	parameter.Value("id").Number().Ge(0)
	parameter.Value("SysBaseMenuID").Number().Ge(0)
	parameter.Value("type").String().Equal(data["parameters"].([]map[string]interface{})[0]["type"].(string))
	parameter.Value("key").String().Equal(data["parameters"].([]map[string]interface{})[0]["key"].(string))
	parameter.Value("value").String().Equal(data["parameters"].([]map[string]interface{})[0]["value"].(string))
	menuId := menu.Value("id").Number().Raw()

	update := map[string]interface{}{
		"id":     menuId,
		"route":  "view/routerHolder.vue",
		"hidden": g.StatusTrue,
		"meta": map[string]interface{}{
			"title":       "132131",
			"icon":        "delete-solid",
			"defaultMenu": false,
			"closeTab":    true,
			"keepAlive":   true,
		},
		"menu_name": "test_update_menu_process",
		"parameters": []map[string]interface{}{
			{
				"type":  "query",
				"key":   "21321",
				"value": "1",
			},
			{
				"type":  "params",
				"key":   "12321",
				"value": "1",
			},
		},
		"pid":  0,
		"path": "21312331",
		"sort": 111,
	}

	obj = auth.POST("v1/admin/menu/updateBaseMenu").
		WithJSON(update).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("更新成功")

	obj = auth.POST("v1/admin/menu/addMenuAuthority").
		WithJSON(map[string]interface{}{
			"authorityId": "9528",
			"menus": []map[string]interface{}{
				{"id": 1},
				{"id": 2},
				{"id": 6},
			},
		}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("添加成功")

	// setUserAuthority
	obj = auth.DELETE("v1/admin/menu/deleteBaseMenu").
		WithJSON(map[string]interface{}{"id": menuId}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("删除成功")

}

func TestMenuAddError(t *testing.T) {
	data := map[string]interface{}{
		"route":     "view/routerHolder.vue",
		"hidden":    g.StatusFalse,
		"menu_name": "dashboard213123",
		"pid":       0,
		"path":      "21312331",
		"sort":      111,
	}
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/menu/addBaseMenu").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("添加失败:存在重复menu_name，请修改name")

}
