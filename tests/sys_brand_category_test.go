package tests

import (
	"net/http"
	"testing"
)

func TestBrandCategoryList(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/brandCategory/getBrandCategoryList").
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
	first.Keys().ContainsOnly("id", "pid", "cateName", "isShow", "path", "sort", "level", "children", "createdAt", "updatedAt")
	first.Value("id").Number().Ge(0)

}

func TestBrandCategoryProcess(t *testing.T) {
	data := map[string]interface{}{
		"cateName": "数码产品",
		"isShow":   false,
		"path":     "http://qmplusimg.henrongyi.top/head.png",
		"sort":     1,
		"level":    1,
		"pid":      "1",
	}
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/brandCategory/createBrandCategory").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("创建成功")

	brandCategory := obj.Value("data").Object()
	brandCategory.Value("id").Number().Ge(0)
	brandCategory.Value("cateName").String().Equal(data["cateName"].(string))
	brandCategory.Value("isShow").Boolean().Equal(data["isShow"].(bool))
	brandCategory.Value("path").String().Equal(data["path"].(string))
	brandCategory.Value("sort").Number().Equal(data["sort"].(int))
	brandCategory.Value("pid").String().Equal(data["pid"].(string))
	brandCategory.Value("level").Number().Equal(data["level"].(int))
	brandCategoryId := brandCategory.Value("id").Number().Raw()

	update := map[string]interface{}{
		"id":       brandCategoryId,
		"cateName": "家电",
		"isShow":   true,
		"path":     "http://qmplusimg.henrongyi.top/head.png",
		"sort":     2,
		"level":    1,
		"pid":      "1",
	}

	obj = auth.PUT("/v1/admin/brandCategory/updateBrandCategory").
		WithJSON(update).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("更新成功")
	brandCategory = obj.Value("data").Object()

	brandCategory.Value("id").Number().Ge(0)
	brandCategory.Value("cateName").String().Equal(update["cateName"].(string))
	brandCategory.Value("isShow").Boolean().Equal(update["isShow"].(bool))
	brandCategory.Value("path").String().Equal(update["path"].(string))
	brandCategory.Value("sort").Number().Equal(update["sort"].(int))
	brandCategory.Value("pid").String().Equal(update["pid"].(string))
	brandCategory.Value("level").Number().Equal(update["level"].(int))

	obj = auth.POST("/v1/admin/brandCategory/getBrandCategoryById").
		WithJSON(map[string]interface{}{"id": brandCategoryId}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
	brandCategory = obj.Value("data").Object()

	brandCategory.Value("id").Number().Ge(0)
	brandCategory.Value("cateName").String().Equal(update["cateName"].(string))
	brandCategory.Value("isShow").Boolean().Equal(update["isShow"].(bool))
	brandCategory.Value("path").String().Equal(update["path"].(string))
	brandCategory.Value("sort").Number().Equal(update["sort"].(int))
	brandCategory.Value("pid").String().Equal(update["pid"].(string))
	brandCategory.Value("level").Number().Equal(update["level"].(int))

	// setUserAuthority
	obj = auth.DELETE("/v1/admin/brandCategory/deleteBrandCategory").
		WithJSON(map[string]interface{}{"id": brandCategoryId}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("删除成功")

}

func TestBrandCategoryRegisterError(t *testing.T) {
	data := map[string]interface{}{
		"cateName": "",
		"isShow":   true,
		"path":     "http://qmplusimg.henrongyi.top/head.png",
		"sort":     2,
		"level":    1,
		"pid":      "1",
	}
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/brandCategory/createBrandCategory").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("Key: 'CreateSysBrandCategory.CateName' Error:Field validation for 'CateName' failed on the 'required' tag")

}
