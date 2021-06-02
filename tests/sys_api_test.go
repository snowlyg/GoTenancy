package tests

import (
	"net/http"
	"testing"
)

func TestApiList(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/api/getApiList").
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
	first.Keys().ContainsOnly("id", "path", "description", "apiGroup", "method", "createdAt", "updatedAt")
	first.Value("id").Number().Ge(0)

}
func TestAllApi(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/api/getAllApis").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("获取成功")

	data := obj.Value("data").Object().Value("apis").Array()
	first := data.First().Object()
	first.Keys().ContainsOnly(
		"id",
		"path",
		"description",
		"apiGroup",
		"method",
		"createdAt",
		"updatedAt",
	)
	first.Value("id").Number().Ge(0)

}

func TestApiProcess(t *testing.T) {
	data := map[string]interface{}{
		"apiGroup":    "test_api_process",
		"description": "test_api_process",
		"method":      "POST",
		"path":        "test_api_process",
	}
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/api/createApi").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("创建成功")

	api := obj.Value("data").Object()
	api.Value("id").Number().Ge(0)
	api.Value("path").String().Equal(data["path"].(string))
	api.Value("description").String().Equal(data["description"].(string))
	api.Value("apiGroup").String().Equal(data["apiGroup"].(string))
	api.Value("method").String().Equal(data["method"].(string))
	apiId := api.Value("id").Number().Raw()
	apiPath := api.Value("path").String().Raw()
	apiMethod := api.Value("method").String().Raw()

	update := map[string]interface{}{
		"id":          apiId,
		"apiGroup":    "update_test_api_process",
		"description": "update_test_api_process",
		"method":      "POST",
		"path":        "update_test_api_process",
	}

	obj = auth.POST("/v1/admin/api/updateApi").
		WithJSON(update).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("修改成功")

	obj = auth.POST("/v1/admin/api/getApiById").
		WithJSON(map[string]interface{}{"id": apiId}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("操作成功")
	api = obj.Value("data").Object().Value("api").Object()

	api.Value("id").Number().Ge(0)
	api.Value("path").String().Equal(update["path"].(string))
	api.Value("description").String().Equal(update["description"].(string))
	api.Value("apiGroup").String().Equal(update["apiGroup"].(string))
	api.Value("method").String().Equal(update["method"].(string))

	// setUserAuthority
	obj = auth.DELETE("/v1/admin/api/deleteApi").
		WithJSON(map[string]interface{}{
			"id":     apiId,
			"path":   apiPath,
			"method": apiMethod,
		}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("删除成功")

}

func TestApiRegisterError(t *testing.T) {
	data := map[string]interface{}{
		"apiGroup":    "user",
		"description": "用户注册",
		"method":      "POST",
		"path":        "/v1/admin/user/register",
	}
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/api/createApi").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(4000)
	obj.Value("msg").String().Equal("添加失败:存在相同api")

}
