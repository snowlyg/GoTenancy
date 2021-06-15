package tests

import (
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/g"
)

func TestCategoryList(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/category/getCategoryList").
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
	first.Keys().ContainsOnly("id", "pid", "cateName", "isShow", "path", "sort", "level", "children", "pic", "createdAt", "updatedAt")
	first.Value("id").Number().Ge(0)

}

func TestCategoryProcess(t *testing.T) {
	data := map[string]interface{}{
		"cateName": "数码产品",
		"isShow":   g.StatusFalse,
		"path":     "http://qmplusimg.henrongyi.top/head.png",
		"sort":     1,
		"level":    1,
		"pid":      "1",
		"pic":      "http://qmplusimg.henrongyi.top/head.png",
	}
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/category/createCategory").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("创建成功")

	category := obj.Value("data").Object()
	category.Value("id").Number().Ge(0)
	category.Value("cateName").String().Equal(data["cateName"].(string))
	category.Value("isShow").Boolean().Equal(data["isShow"].(bool))
	category.Value("path").String().Equal(data["path"].(string))
	category.Value("sort").Number().Equal(data["sort"].(int))
	category.Value("pid").String().Equal(data["pid"].(string))
	category.Value("pic").String().Equal(data["pic"].(string))
	category.Value("level").Number().Equal(data["level"].(int))
	categoryId := category.Value("id").Number().Raw()

	update := map[string]interface{}{
		"id":       categoryId,
		"cateName": "家电",
		"isShow":   g.StatusTrue,
		"path":     "http://qmplusimg.henrongyi.top/head.png",
		"sort":     2,
		"level":    1,
		"pid":      "1",
		"pic":      "http://qmplusimg.henrongyi.top/head.png",
	}

	obj = auth.PUT("v1/admin/category/updateCategory").
		WithJSON(update).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("更新成功")
	category = obj.Value("data").Object()

	category.Value("id").Number().Ge(0)
	category.Value("cateName").String().Equal(update["cateName"].(string))
	category.Value("isShow").Boolean().Equal(update["isShow"].(bool))
	category.Value("path").String().Equal(update["path"].(string))
	category.Value("sort").Number().Equal(update["sort"].(int))
	category.Value("pid").String().Equal(update["pid"].(string))
	category.Value("pic").String().Equal(update["pic"].(string))
	category.Value("level").Number().Equal(update["level"].(int))

	obj = auth.POST("v1/admin/category/getCategoryById").
		WithJSON(map[string]interface{}{"id": categoryId}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
	category = obj.Value("data").Object()

	category.Value("id").Number().Ge(0)
	category.Value("cateName").String().Equal(update["cateName"].(string))
	category.Value("isShow").Boolean().Equal(update["isShow"].(bool))
	category.Value("path").String().Equal(update["path"].(string))
	category.Value("sort").Number().Equal(update["sort"].(int))
	category.Value("pid").String().Equal(update["pid"].(string))
	category.Value("pic").String().Equal(update["pic"].(string))
	category.Value("level").Number().Equal(update["level"].(int))

	// setUserAuthority
	obj = auth.DELETE("v1/admin/category/deleteCategory").
		WithJSON(map[string]interface{}{"id": categoryId}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("删除成功")

}

func TestCategoryRegisterError(t *testing.T) {
	data := map[string]interface{}{
		"cateName": "",
		"isShow":   g.StatusTrue,
		"path":     "http://qmplusimg.henrongyi.top/head.png",
		"sort":     2,
		"level":    1,
		"pid":      "1",
		"pic":      "http://qmplusimg.henrongyi.top/head.png",
	}
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/category/createCategory").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("Key: 'CreateTenancyCategory.CateName' Error:Field validation for 'CateName' failed on the 'required' tag")

}
