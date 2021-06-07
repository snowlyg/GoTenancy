package tests

import (
	"net/http"
	"testing"

	"github.com/snowlyg/multi"
)

func TestAuthorityList(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/authority/getAuthorityList").
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
		"authorityId",
		"authorityName",
		"authorityType",
		"parentId",
		"dataAuthorityId",
		"children",
		"menus",
		"defaultRouter",
		"createdAt",
		"updatedAt",
	)

}
func TestAdminAuthority(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/authority/getAdminAuthorityList").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object()
	list := data.Value("list").Array()
	list.Length().Ge(0)
	first := list.First().Object()
	first.Keys().ContainsOnly(
		"authorityId",
		"authorityName",
		"authorityType",
		"parentId",
		"dataAuthorityId",
		"children",
		"menus",
		"defaultRouter",
		"createdAt",
		"updatedAt",
	)

}
func TestTenancyAuthority(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/authority/getTenancyAuthorityList").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object()
	list := data.Value("list").Array()
	list.Length().Ge(0)
	first := list.First().Object()
	first.Keys().ContainsOnly(
		"authorityId",
		"authorityName",
		"authorityType",
		"parentId",
		"dataAuthorityId",
		"children",
		"menus",
		"defaultRouter",
		"createdAt",
		"updatedAt",
	)

}
func TestGeneralAuthority(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/authority/getGeneralAuthorityList").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object()
	list := data.Value("list").Array()
	list.Length().Ge(0)
	first := list.First().Object()
	first.Keys().ContainsOnly(
		"authorityId",
		"authorityName",
		"authorityType",
		"parentId",
		"dataAuthorityId",
		"children",
		"menus",
		"defaultRouter",
		"createdAt",
		"updatedAt",
	)

}

func TestAuthorityProcess(t *testing.T) {
	data := map[string]interface{}{
		"authorityId":   "9523",
		"authorityName": "测试角色",
		"parentId":      "0",
		"authorityType": multi.AdminAuthority,
		"dataAuthorityId": []map[string]interface{}{
			{"authorityId": "8881"},
		},
	}
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/authority/createAuthority").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("创建成功")

	authority := obj.Value("data").Object().Value("authority").Object()
	authority.Value("parentId").String().Equal(data["parentId"].(string))
	authority.Value("authorityName").String().Equal(data["authorityName"].(string))
	authority.Value("authorityId").String().Equal(data["authorityId"].(string))
	authority.Value("authorityType").Number().Equal(data["authorityType"].(int))
	authority.Value("children").Null()
	authority.Value("menus").Null()
	dataAuthorityId := authority.Value("dataAuthorityId").Array().First().Object()
	dataAuthorityId.Value("authorityId").String().Equal(data["dataAuthorityId"].([]map[string]interface{})[0]["authorityId"].(string))
	authorityId := authority.Value("authorityId").String().Raw()

	update := map[string]interface{}{
		"authorityId":   "9523",
		"authorityName": "测试角色",
		"parentId":      "0",
		"authorityType": multi.AdminAuthority,
		"dataAuthorityId": []map[string]interface{}{
			{"authorityId": "8881"},
		},
	}

	obj = auth.PUT("/v1/admin/authority/updateAuthority").
		WithJSON(update).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("更新成功")

	authority = obj.Value("data").Object().Value("authority").Object()
	authority.Value("parentId").String().Equal(update["parentId"].(string))
	authority.Value("authorityName").String().Equal(update["authorityName"].(string))
	authority.Value("authorityId").String().Equal(update["authorityId"].(string))
	authority.Value("authorityType").Number().Equal(data["authorityType"].(int))
	authority.Value("children").Null()
	authority.Value("menus").Null()
	dataAuthorityId = authority.Value("dataAuthorityId").Array().First().Object()
	dataAuthorityId.Value("authorityId").String().Equal(update["dataAuthorityId"].([]map[string]interface{})[0]["authorityId"].(string))

	copy := map[string]interface{}{
		"authority": map[string]interface{}{
			"authorityId":   "9511",
			"authorityName": "测试角色",
			"authorityType": multi.AdminAuthority,
			"parentId":      "0",
			"dataAuthorityId": []map[string]interface{}{
				{"authorityId": "8881"},
			},
		},
		"authorityId":   "9523",
		"authorityName": "测试角色",
		"parentId":      "0",
		"authorityType": multi.AdminAuthority,
		"dataAuthorityId": []map[string]interface{}{
			{"authorityId": "8881"},
		},
		"oldAuthorityId": "9528",
	}

	obj = auth.POST("/v1/admin/authority/copyAuthority").
		WithJSON(copy).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("拷贝成功")

	authority = obj.Value("data").Object().Value("authority").Object()
	authority.Value("parentId").String().Equal(copy["authority"].(map[string]interface{})["parentId"].(string))
	authority.Value("authorityName").String().Equal(copy["authority"].(map[string]interface{})["authorityName"].(string))
	authority.Value("authorityId").String().Equal(copy["authority"].(map[string]interface{})["authorityId"].(string))
	authority.Value("authorityType").Number().Equal(copy["authority"].(map[string]interface{})["authorityType"].(int))
	authority.Value("children").Array().Empty()

	dataAuthorityId = authority.Value("dataAuthorityId").Array().First().Object()
	dataAuthorityId.Value("authorityId").String().Equal(copy["dataAuthorityId"].([]map[string]interface{})[0]["authorityId"].(string))

	setdata := map[string]interface{}{
		"authorityId": "8881",
		"dataAuthorityId": []map[string]interface{}{
			{"authorityId": "8881"},
		},
	}
	obj = auth.POST("/v1/admin/authority/setDataAuthority").
		WithJSON(setdata).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("设置成功")

	// delete
	obj = auth.DELETE("/v1/admin/authority/deleteAuthority").
		WithJSON(map[string]interface{}{
			"authorityId": authorityId,
		}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("删除成功")

	// delete
	obj = auth.DELETE("/v1/admin/authority/deleteAuthority").
		WithJSON(map[string]interface{}{
			"authorityId": "9511",
		}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("删除成功")

}

func TestAuthorityRegisterError(t *testing.T) {
	data := map[string]interface{}{
		"authorityId":     "999",
		"authorityName":   "测试角色",
		"parentId":        "0",
		"authorityType":   multi.AdminAuthority,
		"dataAuthorityId": nil,
	}
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/authority/createAuthority").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("创建失败存在相同角色id")

}
