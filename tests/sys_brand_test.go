package tests

import (
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/g"
)

func TestBrandList(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/brand/getBrandList").
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
	first.Keys().ContainsOnly("id", "brandName", "isShow", "pic", "sort", "brandCategoryId", "createdAt", "updatedAt")
	first.Value("id").Number().Ge(0)
}

func TestBrandProcess(t *testing.T) {
	data := map[string]interface{}{
		"brandName":       "冈本",
		"isShow":          false,
		"pic":             "http://qmplusimg.henrongyi.top/head.png",
		"sort":            1,
		"brandCategoryId": 1,
	}
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/brand/createBrand").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("创建成功")

	brand := obj.Value("data").Object()
	brand.Value("id").Number().Ge(0)
	brand.Value("brandName").String().Equal(data["brandName"].(string))
	brand.Value("isShow").Number().Equal(data["isShow"].(int))
	brand.Value("pic").String().Equal(data["pic"].(string))
	brand.Value("sort").Number().Equal(data["sort"].(int))
	brandId := brand.Value("id").Number().Raw()

	update := map[string]interface{}{
		"id":              brandId,
		"brandName":       "威尔刚",
		"isShow":          true,
		"pic":             "http://qmplusimg.henrongyi.top/head.png",
		"sort":            2,
		"brandCategoryId": 1,
	}

	obj = auth.PUT("v1/admin/brand/updateBrand").
		WithJSON(update).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("更新成功")
	brand = obj.Value("data").Object()

	brand.Value("id").Number().Ge(0)
	brand.Value("brandName").String().Equal(update["brandName"].(string))
	brand.Value("isShow").Number().Equal(update["isShow"].(int))
	brand.Value("pic").String().Equal(update["pic"].(string))
	brand.Value("sort").Number().Equal(update["sort"].(int))

	obj = auth.POST("v1/admin/brand/getBrandById").
		WithJSON(map[string]interface{}{"id": brandId}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
	brand = obj.Value("data").Object()

	brand.Value("id").Number().Ge(0)
	brand.Value("brandName").String().Equal(update["brandName"].(string))
	brand.Value("isShow").Number().Equal(update["isShow"].(int))
	brand.Value("pic").String().Equal(update["pic"].(string))
	brand.Value("sort").Number().Equal(update["sort"].(int))

	obj = auth.POST("v1/admin/brand/setBrandCate").
		WithJSON(map[string]interface{}{"id": brandId, "brandCategoryId": 1}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("设置成功")

	// setUserAuthority
	obj = auth.DELETE("v1/admin/brand/deleteBrand").
		WithJSON(map[string]interface{}{"id": brandId}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("删除成功")
}

func TestBrandRegisterError(t *testing.T) {
	data := map[string]interface{}{
		"brandName":       "",
		"isShow":          g.StatusTrue,
		"pic":             "http://qmplusimg.henrongyi.top/head.png",
		"sort":            2,
		"brandCategoryId": 1,
	}
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/brand/createBrand").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("Key: 'CreateSysBrand.BrandName' Error:Field validation for 'BrandName' failed on the 'required' tag")

}
