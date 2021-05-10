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
	data.Value("total").Number().Equal(1)
	data.Value("list").Array().Length().Ge(0)

	user := data.Value("list").Array().First().Object()
	// user.Keys().ContainsOnly("ID")
	user.Value("ID").Number().Ge(0)

	baseLogOut(auth)
}

func TestTenancyProcess(t *testing.T) {
	data := map[string]interface{}{
		"name":            "宝安妇女儿童医院",
		"tele":            "0755-23568911",
		"address":         "xxx街道666号",
		"business_hours":  "08:30-17:30",
		"sys_region_code": 0}
	auth := baseWithLoginTester(t)
	obj := auth.POST("/v1/admin/tenancy/createTenancy").
		WithJSON(data).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("创建成功")

	tenancy := obj.Value("data").Object()
	// tenancy.Keys().ContainsOnly("ID")
	tenancy.Value("ID").Number().Ge(0)
	tenancy.Value("name").String().Equal(data["name"].(string))
	tenancy.Value("tele").String().Equal(data["tele"].(string))
	tenancy.Value("address").String().Equal(data["address"].(string))
	tenancy.Value("business_hours").String().Equal(data["business_hours"].(string))
	tenancy.Value("sys_region_code").Number().Equal(data["sys_region_code"].(int))
	tenancyId := tenancy.Value("ID").Number().Raw()

	update := map[string]interface{}{
		"id":              tenancyId,
		"name":            "宝安妇女儿童附属医院",
		"tele":            "0755-235689111",
		"address":         "xxx街道667号",
		"business_hours":  "08:30-17:40",
		"sys_region_code": 0}

	// setAdminInfo
	obj = auth.PUT("/v1/admin/tenancy/updateTenancy").
		WithJSON(update).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("更新成功")
	tenancy = obj.Value("data").Object()
	// tenancy.Keys().ContainsOnly("ID")
	tenancy.Value("ID").Number().Ge(0)
	tenancy.Value("name").String().Equal(update["name"].(string))
	tenancy.Value("tele").String().Equal(update["tele"].(string))
	tenancy.Value("address").String().Equal(update["address"].(string))
	tenancy.Value("business_hours").String().Equal(update["business_hours"].(string))
	tenancy.Value("sys_region_code").Number().Equal(update["sys_region_code"].(int))

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
		"name":            "宝安中心人民医院",
		"tele":            "0755-23568911",
		"address":         "xxx街道666号",
		"business_hours":  "08:30-17:30",
		"sys_region_code": 0}
	auth := baseWithLoginTester(t)
	obj := auth.POST("/v1/admin/tenancy/createTenancy").
		WithJSON(data).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(7)
	obj.Value("msg").String().Equal("创建失败")

	baseLogOut(auth)
}
