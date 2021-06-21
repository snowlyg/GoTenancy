package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/g"
)

func TestClientCategoryList(t *testing.T) {
	auth := tenancyWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/merchant/category/getCategoryList").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Array()

	data.Length().Ge(0)
	first := data.First().Object()
	first.Keys().ContainsOnly("id", "pid", "cateName", "status", "path", "sort", "level", "children", "pic", "createdAt", "updatedAt")
	first.Value("id").Number().Ge(0)
}

func TestClientCategorySelect(t *testing.T) {
	auth := tenancyWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.GET("v1/merchant/category/getCategorySelect").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")
}

func TestGetAdminCategorySelect(t *testing.T) {
	auth := tenancyWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.GET("v1/merchant/category/getAdminCategorySelect").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")
}

func TestClientCategoryProcess(t *testing.T) {
	data := map[string]interface{}{
		"cateName": "数码产品",
		"status":   g.StatusTrue,
		"path":     "http://qmplusimg.henrongyi.top/head.png",
		"sort":     1,
		"level":    1,
		"pid":      1,
		"pic":      "http://qmplusimg.henrongyi.top/head.png",
	}
	auth := tenancyWithLoginTester(t)
	defer baseLogOut(auth)

	obj := auth.POST("v1/merchant/category/createCategory").
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

	obj = auth.PUT(fmt.Sprintf("v1/merchant/category/updateCategory/%d", int(categoryId))).
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

	obj = auth.GET(fmt.Sprintf("v1/merchant/category/getCategoryById/%d", int(categoryId))).
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

	obj = auth.POST("v1/merchant/category/changeCategoryStatus").
		WithJSON(map[string]interface{}{
			"id":     categoryId,
			"status": g.StatusTrue,
		}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("设置成功")

	obj = auth.GET("v1/merchant/category/getCreateTenancyCategoryMap").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	obj = auth.GET(fmt.Sprintf("v1/merchant/category/getUpdateTenancyCategoryMap/%d", int(categoryId))).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	// deleteCategory
	obj = auth.DELETE(fmt.Sprintf("v1/merchant/category/deleteCategory/%d", int(categoryId))).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("删除成功")

}

func TestClientCategoryRegisterError(t *testing.T) {
	data := map[string]interface{}{
		"cateName": "",
		"status":   g.StatusTrue,
		"path":     "http://qmplusimg.henrongyi.top/head.png",
		"sort":     2,
		"level":    1,
		"pid":      1,
		"pic":      "http://qmplusimg.henrongyi.top/head.png",
	}
	auth := tenancyWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/merchant/category/createCategory").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("Key: 'TenancyCategory.BaseTenancyCategory.CateName' Error:Field validation for 'CateName' failed on the 'required' tag")

}
