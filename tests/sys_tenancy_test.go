package tests

import (
	"testing"

	"github.com/kataras/iris/v12/httptest"
)

func TestTenancyList(t *testing.T) {
	auth := baseWithLoginTester(t)
	obj := auth.POST("/v1/admin/tenancy/getTenancyList").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10}).
		Expect().Status(httptest.StatusOK).JSON().Object()
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
	first.Keys().ContainsOnly("id", "uuid", "name", "tele", "address", "businessTime", "sysRegionCode", "createdAt", "updatedAt")
	first.Value("id").Number().Ge(0)

	baseLogOut(auth)
}

func TestTenancyByRegion(t *testing.T) {
	auth := baseWithLoginTester(t)
	obj := auth.GET("/v1/admin/tenancy/getTenancies/1").
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("获取成功")
	obj.Value("data").Array().Length().Equal(1)

	baseLogOut(auth)
}

func TestTenancyProcess(t *testing.T) {
	data := map[string]interface{}{
		"name":          "宝安妇女儿童医院",
		"tele":          "0755-23568911",
		"address":       "xxx街道666号",
		"businessTime":  "08:30-17:30",
		"sysRegionCode": 1,
	}
	auth := baseWithLoginTester(t)
	obj := auth.POST("/v1/admin/tenancy/createTenancy").
		WithJSON(data).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("创建成功")

	tenancy := obj.Value("data").Object()
	tenancy.Value("id").Number().Ge(0)
	tenancy.Value("uuid").String().NotEmpty()
	tenancy.Value("name").String().Equal(data["name"].(string))
	tenancy.Value("tele").String().Equal(data["tele"].(string))
	tenancy.Value("address").String().Equal(data["address"].(string))
	tenancy.Value("businessTime").String().Equal(data["businessTime"].(string))
	tenancy.Value("sysRegionCode").Number().Equal(data["sysRegionCode"].(int))
	tenancyId := tenancy.Value("id").Number().Raw()

	update := map[string]interface{}{
		"id":            tenancyId,
		"name":          "宝安妇女儿童附属医院",
		"tele":          "0755-235689111",
		"address":       "xxx街道667号",
		"businessTime":  "08:30-17:40",
		"sysRegionCode": 1,
	}

	obj = auth.PUT("/v1/admin/tenancy/updateTenancy").
		WithJSON(update).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("更新成功")
	tenancy = obj.Value("data").Object()

	tenancy.Value("id").Number().Ge(0)
	tenancy.Value("name").String().Equal(update["name"].(string))
	tenancy.Value("tele").String().Equal(update["tele"].(string))
	tenancy.Value("address").String().Equal(update["address"].(string))
	tenancy.Value("businessTime").String().Equal(update["businessTime"].(string))
	tenancy.Value("sysRegionCode").Number().Equal(update["sysRegionCode"].(int))

	obj = auth.POST("/v1/admin/tenancy/getTenancyById").
		WithJSON(map[string]interface{}{"id": tenancyId}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("操作成功")
	tenancy = obj.Value("data").Object()

	tenancy.Value("id").Number().Ge(0)
	tenancy.Value("uuid").String().NotEmpty()
	tenancy.Value("name").String().Equal(update["name"].(string))
	tenancy.Value("tele").String().Equal(update["tele"].(string))
	tenancy.Value("address").String().Equal(update["address"].(string))
	tenancy.Value("businessTime").String().Equal(update["businessTime"].(string))
	tenancy.Value("sysRegionCode").Number().Equal(update["sysRegionCode"].(int))

	// setTenancyRegion
	obj = auth.POST("/v1/admin/tenancy/setTenancyRegion").
		WithJSON(map[string]interface{}{
			"id":            tenancyId,
			"sysRegionCode": 2,
		}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("设置成功")

	// setUserAuthority
	obj = auth.DELETE("/v1/admin/tenancy/deleteTenancy").
		WithJSON(map[string]interface{}{"id": tenancyId}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("删除成功")

	baseLogOut(auth)
}

func TestTenancyRegisterError(t *testing.T) {
	data := map[string]interface{}{
		"name":          "宝安中心人民医院",
		"tele":          "0755-23568911",
		"address":       "xxx街道666号",
		"businessTime":  "08:30-17:30",
		"sysRegionCode": 1}
	auth := baseWithLoginTester(t)
	obj := auth.POST("/v1/admin/tenancy/createTenancy").
		WithJSON(data).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(4000)
	obj.Value("msg").String().Equal("创建失败")

	baseLogOut(auth)
}
