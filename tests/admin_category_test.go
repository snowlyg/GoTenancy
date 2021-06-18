package tests

import (
	"fmt"
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
	first.Keys().ContainsOnly("id", "pid", "cateName", "status", "path", "sort", "level", "children", "pic", "createdAt", "updatedAt")
	first.Value("id").Number().Ge(0)
}

func TestCategorySelect(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.GET("v1/admin/category/getCategorySelect").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")
}

func TestCategoryProcess(t *testing.T) {
	data := map[string]interface{}{
		"cateName": "数码产品",
		"status":   g.StatusTrue,
		"path":     "http://qmplusimg.henrongyi.top/head.png",
		"sort":     1,
		"level":    1,
		"pid":      1,
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
	category.Value("status").Number().Equal(data["status"].(int))
	category.Value("path").String().Equal(data["path"].(string))
	category.Value("sort").Number().Equal(data["sort"].(int))
	category.Value("pid").Number().Equal(data["pid"].(int))
	category.Value("pic").String().Equal(data["pic"].(string))
	category.Value("level").Number().Equal(data["level"].(int))
	categoryId := category.Value("id").Number().Raw()

	update := map[string]interface{}{
		"cateName": "家电",
		"status":   g.StatusTrue,
		"path":     "http://qmplusimg.henrongyi.top/head.png",
		"sort":     2,
		"level":    1,
		"pid":      1,
		"pic":      "http://qmplusimg.henrongyi.top/head.png",
	}

	obj = auth.PUT(fmt.Sprintf("v1/admin/category/updateCategory/%d", int(categoryId))).
		WithJSON(update).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("更新成功")
	category = obj.Value("data").Object()

	category.Value("cateName").String().Equal(update["cateName"].(string))
	category.Value("status").Number().Equal(update["status"].(int))
	category.Value("path").String().Equal(update["path"].(string))
	category.Value("sort").Number().Equal(update["sort"].(int))
	category.Value("pid").Number().Equal(update["pid"].(int))
	category.Value("pic").String().Equal(update["pic"].(string))
	category.Value("level").Number().Equal(update["level"].(int))

	obj = auth.GET(fmt.Sprintf("v1/admin/category/getCategoryById/%d", int(categoryId))).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
	category = obj.Value("data").Object()

	category.Value("id").Number().Ge(0)
	category.Value("cateName").String().Equal(update["cateName"].(string))
	category.Value("status").Number().Equal(update["status"].(int))
	category.Value("path").String().Equal(update["path"].(string))
	category.Value("sort").Number().Equal(update["sort"].(int))
	category.Value("pid").Number().Equal(update["pid"].(int))
	category.Value("pic").String().Equal(update["pic"].(string))
	category.Value("level").Number().Equal(update["level"].(int))

	obj = auth.POST("v1/admin/category/changeCategoryStatus").
		WithJSON(map[string]interface{}{
			"id":     categoryId,
			"status": g.StatusTrue,
		}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("设置成功")

	obj = auth.GET("v1/admin/category/getCreateTenancyCategoryMap").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	obj = auth.GET(fmt.Sprintf("v1/admin/category/getUpdateTenancyCategoryMap/%d", int(categoryId))).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	// deleteCategory
	obj = auth.DELETE(fmt.Sprintf("v1/admin/category/deleteCategory/%d", int(categoryId))).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("删除成功")

}

func TestCategoryRegisterError(t *testing.T) {
	data := map[string]interface{}{
		"cateName": "",
		"status":   g.StatusTrue,
		"path":     "http://qmplusimg.henrongyi.top/head.png",
		"sort":     2,
		"level":    1,
		"pid":      1,
		"pic":      "http://qmplusimg.henrongyi.top/head.png",
	}
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/category/createCategory").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("Key: 'TenancyCategory.BaseTenancyCategory.CateName' Error:Field validation for 'CateName' failed on the 'required' tag")

}
