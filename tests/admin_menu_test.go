package tests

import (
	"fmt"
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

func TestClientMenuList(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/menu/getClientMenuList").
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
	obj := auth.GET("v1/admin/menu/getMenuList").
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
		"path":      "/100",
		"hidden":    g.StatusFalse,
		"menu_name": "test_menu_process",
		"pid":       0,
		"route":     "test_menu_process",
		"sort":      111,
	}
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.GET("v1/admin/menu/getAddMenuMap").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	obj = auth.POST("v1/admin/menu/addBaseMenu").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("添加成功")

	menu := obj.Value("data").Object()
	menu.Value("id").Number().Ge(0)
	menu.Value("pid").Number().Equal(data["pid"].(int))
	menu.Value("menu_name").String().Equal(data["menu_name"].(string))
	menu.Value("path").String().Equal(data["path"].(string))
	menu.Value("hidden").Number().Equal(data["hidden"].(int))
	menu.Value("route").String().Equal(data["route"].(string))
	menu.Value("sort").Number().Equal(data["sort"].(int))
	menu.Value("authoritys").Null()
	menu.Value("children").Null()
	menuId := menu.Value("id").Number().Raw()

	if menuId > 0 {
		obj = auth.GET("v1/admin/menu/getAddMenuMap").
			WithJSON(data).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		obj = auth.GET("v1/admin/menu/getAddMenuMap").
			WithJSON(data).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		obj = auth.GET(fmt.Sprintf("v1/admin/menu/getEditMenuMap/%d", int(menuId))).
			WithJSON(data).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		update := map[string]interface{}{
			"path":      "/",
			"hidden":    g.StatusTrue,
			"menu_name": "test_update_menu_process",
			"pid":       0,
			"route":     "test_update_menu_process",
			"sort":      111,
		}
		obj = auth.POST(fmt.Sprintf("v1/admin/menu/updateBaseMenu/%d", int(menuId))).
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

		// deleteBaseMenu
		obj = auth.DELETE(fmt.Sprintf("v1/admin/menu/deleteBaseMenu/%d", int(menuId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("删除成功")
	}

}

func TestMenuAddError(t *testing.T) {
	data := map[string]interface{}{
		"path":      "/",
		"hidden":    g.StatusFalse,
		"menu_name": "dashboard",
		"pid":       0,
		"route":     "/admin/dashboard",
		"sort":      111,
	}
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/menu/addBaseMenu").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("添加失败:存在重复route，请修改route")

}
