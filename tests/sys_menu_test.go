package tests

import (
	"net/http"
	"testing"
)

func TestMenu(t *testing.T) {
	auth := baseWithLoginTester(t)
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

	baseLogOut(auth)
}
func TestBaseMenu(t *testing.T) {
	auth := baseWithLoginTester(t)
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

	baseLogOut(auth)
}

func TestMenuList(t *testing.T) {
	auth := baseWithLoginTester(t)
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

	baseLogOut(auth)
}

func TestMenuProcess(t *testing.T) {
	data := map[string]interface{}{
		"name":      "中德澳线上点餐商城",
		"appId":     "YJ3s1abt7MAfT6gWVKoDresdfsdf",
		"appSecret": "tRE49zaf5NCm6PidFZoaFg3u4WCHDok7fxgL63yV0pF4AMsdfsdfsdfssa",
		"remark":    "中德澳线上点餐商城",
	}
	auth := baseWithLoginTester(t)
	obj := auth.POST("/v1/admin/menu/createMenu").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("创建成功")

	menu := obj.Value("data").Object()
	menu.Value("id").Number().Ge(0)
	menu.Value("uuid").String().NotEmpty()
	menu.Value("name").String().Equal(data["name"].(string))
	menu.Value("appId").String().Equal(data["appId"].(string))
	menu.Value("appSecret").String().Equal(data["appSecret"].(string))
	menu.Value("remark").String().Equal(data["remark"].(string))
	menuId := menu.Value("id").Number().Raw()

	update := map[string]interface{}{
		"id":        menuId,
		"name":      "中德澳线上点餐商城1",
		"appId":     "YJ3s1abt7MAfT6gWVKoDjnhjsfsd",
		"appSecret": "tRE49zaf5NCm6PidFZoaFg3u4WCHDok7fxgL63yV0pF4AMsdfbnfgh",
		"remark":    "中德澳线上点餐商城1",
	}

	obj = auth.PUT("/v1/admin/menu/updateMenu").
		WithJSON(update).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("更新成功")
	menu = obj.Value("data").Object()

	menu.Value("id").Number().Ge(0)
	menu.Value("uuid").String().NotEmpty()
	menu.Value("name").String().Equal(update["name"].(string))
	menu.Value("appId").String().Equal(update["appId"].(string))
	menu.Value("appSecret").String().Equal(update["appSecret"].(string))
	menu.Value("remark").String().Equal(update["remark"].(string))

	obj = auth.POST("/v1/admin/menu/getMenuById").
		WithJSON(map[string]interface{}{"id": menuId}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("操作成功")
	menu = obj.Value("data").Object()

	menu.Value("id").Number().Ge(0)
	menu.Value("uuid").String().NotEmpty()
	menu.Value("name").String().Equal(update["name"].(string))
	menu.Value("appId").String().Equal(update["appId"].(string))
	menu.Value("appSecret").String().Equal(update["appSecret"].(string))
	menu.Value("remark").String().Equal(update["remark"].(string))

	// setUserAuthority
	obj = auth.DELETE("/v1/admin/menu/deleteMenu").
		WithJSON(map[string]interface{}{"id": menuId}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("删除成功")

	baseLogOut(auth)
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
		"name": "212331",
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
	obj := auth.POST("/v1/admin/menu/addBaseMenu").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(4000)
	obj.Value("msg").String().Equal("添加失败:存在重复name，请修改name")

	baseLogOut(auth)
}
