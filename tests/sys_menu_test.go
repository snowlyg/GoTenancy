package tests

import (
	"net/http"
	"testing"
)

func TestMenu(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/menu/getMenu").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("获取成功")

	data := obj.Value("data").Object().Value("menus").Array()
	data.Length().Ge(0)
	first := data.First().Object()
	first.Keys().ContainsOnly(
		"id",
		"parentId",
		"name",
		"path",
		"hidden",
		"component",
		"sort",
		"meta",
		"authoritys",
		"menuId",
		"children",
		"parameters",
		"createdAt",
		"updatedAt",
	)
	first.Value("id").Number().Ge(0)

}
func TestBaseMenu(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/menu/getBaseMenuTree").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("获取成功")

	data := obj.Value("data").Object().Value("menus").Array()
	data.Length().Ge(0)
	first := data.First().Object()
	first.Keys().ContainsOnly(
		"id",
		"parentId",
		"name",
		"path",
		"hidden",
		"component",
		"sort",
		"meta",
		"authoritys",
		"children",
		"parameters",
		"createdAt",
		"updatedAt",
	)
	first.Value("id").Number().Ge(0)

}

func TestMenuList(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/menu/getMenuList").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10}).
		Expect().Status(http.StatusOK).JSON().Object()
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
	first.Keys().ContainsOnly(
		"id",
		"parentId",
		"name",
		"path",
		"hidden",
		"component",
		"sort",
		"meta",
		"authoritys",
		"children",
		"parameters",
		"createdAt",
		"updatedAt",
	)
	first.Value("id").Number().Ge(0)

}

func TestMenuProcess(t *testing.T) {
	data := map[string]interface{}{
		"component": "view/routerHolder.vue",
		"hidden":    true,
		"meta": map[string]interface{}{
			"title":       "132131",
			"icon":        "delete-solid",
			"defaultMenu": false,
			"closeTab":    true,
			"keepAlive":   true,
		},
		"name": "test_menu_process",
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
		"parentId": "0",
		"path":     "21312331",
		"sort":     111,
	}
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/menu/addBaseMenu").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("添加成功")

	menu := obj.Value("data").Object()
	menu.Value("id").Number().Ge(0)
	menu.Value("parentId").String().Equal(data["parentId"].(string))
	menu.Value("name").String().Equal(data["name"].(string))
	menu.Value("path").String().Equal(data["path"].(string))
	menu.Value("hidden").Boolean().Equal(data["hidden"].(bool))
	menu.Value("component").String().Equal(data["component"].(string))
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
		"id":        menuId,
		"component": "view/routerHolder.vue",
		"hidden":    true,
		"meta": map[string]interface{}{
			"title":       "132131",
			"icon":        "delete-solid",
			"defaultMenu": false,
			"closeTab":    true,
			"keepAlive":   true,
		},
		"name": "test_update_menu_process",
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
		"parentId": "0",
		"path":     "21312331",
		"sort":     111,
	}

	obj = auth.POST("/v1/admin/menu/updateBaseMenu").
		WithJSON(update).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("更新成功")

	obj = auth.POST("/v1/admin/menu/addMenuAuthority").
		WithJSON(map[string]interface{}{
			"authorityId": "9528",
			"menus": []map[string]interface{}{
				{"id": 1},
				{"id": 2},
				{"id": 6},
			},
		}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("添加成功")

	// setUserAuthority
	obj = auth.DELETE("/v1/admin/menu/deleteBaseMenu").
		WithJSON(map[string]interface{}{"id": menuId}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("删除成功")

}

func TestMenuAddError(t *testing.T) {
	data := map[string]interface{}{
		"component": "view/routerHolder.vue",
		"hidden":    true,
		"meta": map[string]interface{}{
			"title":       "132131",
			"icon":        "delete-solid",
			"defaultMenu": false,
			"closeTab":    true,
			"keepAlive":   true,
		},
		"name": "dashboard",
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
		"parentId": "0",
		"path":     "21312331",
		"sort":     111,
	}
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/menu/addBaseMenu").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(4000)
	obj.Value("msg").String().Equal("添加失败:存在重复name，请修改name")

}
