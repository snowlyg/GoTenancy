package tests

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/snowlyg/go-tenancy/g"
)

func TestTenancyList(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/tenancy/getTenancyList").
		WithJSON(map[string]interface{}{
			"page":     1,
			"pageSize": 10,
			"status":   "1",
			"keyword":  "宝安",
			"date":     fmt.Sprintf("%d/%02d/01-%d/%02d/28", time.Now().Year(), time.Now().Month(), time.Now().Year(), time.Now().Month()),
		}).
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
	first.Keys().ContainsOnly("id", "uuid", "name", "tele", "address", "businessTime", "sysRegionCode", "status", "createdAt", "updatedAt")
	first.Value("id").Number().Ge(0)
}
func TestTenancyListNextMonth(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/tenancy/getTenancyList").
		WithJSON(map[string]interface{}{
			"page":     1,
			"pageSize": 10,
			"status":   "1",
			"keyword":  "宝安",
			"date":     fmt.Sprintf("%d/%02d/01-%d/%02d/28", time.Now().Year(), time.Now().AddDate(0, 1, 0).Month(), time.Now().Year(), time.Now().AddDate(0, 1, 0).Month()),
		}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize")
	data.Value("pageSize").Number().Equal(10)
	data.Value("page").Number().Equal(1)
	data.Value("total").Number().Equal(0)
	data.Value("list").Array().Empty()
}

func TestTenancyListNoKeyword(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/tenancy/getTenancyList").
		WithJSON(map[string]interface{}{
			"page":     1,
			"pageSize": 10,
			"status":   "1",
		}).
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
	first.Keys().ContainsOnly("id", "uuid", "name", "tele", "address", "businessTime", "sysRegionCode", "status", "createdAt", "updatedAt")
	first.Value("id").Number().Ge(0)
}

func TestTenancyListKeyword(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/tenancy/getTenancyList").
		WithJSON(map[string]interface{}{
			"page":     1,
			"pageSize": 10,
			"status":   "1",
			"keyword":  "北京",
		}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize")
	data.Value("pageSize").Number().Equal(10)
	data.Value("page").Number().Equal(1)
	data.Value("total").Number().Equal(0)
	data.Value("list").Array().Empty()
}

func TestTenancyListStatus(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/tenancy/getTenancyList").
		WithJSON(map[string]interface{}{
			"page":     1,
			"pageSize": 10,
			"status":   "2",
			"keyword":  "",
		}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize")
	data.Value("pageSize").Number().Equal(10)
	data.Value("page").Number().Equal(1)
	data.Value("total").Number().Equal(0)
	data.Value("list").Array().Empty()
}

func TestTenancyByRegion(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.GET("/v1/admin/tenancy/getTenancies/1").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")
	obj.Value("data").Array().Length().Ge(1)
}

func TestGetTenancyCount(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.GET("/v1/admin/tenancy/getTenancyCount").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")
	obj.Value("data").Object().Value("invalid").Equal(0)
	obj.Value("data").Object().Value("valid").Equal(2)
}

func TestTenancyProcess(t *testing.T) {
	data := map[string]interface{}{
		"name":          "宝安妇女儿童医院",
		"tele":          "0755-23568911",
		"address":       "xxx街道666号",
		"businessTime":  "08:30-17:30",
		"status":        g.StatusTrue,
		"sysRegionCode": 1,
	}
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/tenancy/createTenancy").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("创建成功")

	tenancy := obj.Value("data").Object()
	tenancy.Value("id").Number().Ge(0)
	tenancy.Value("uuid").String().NotEmpty()
	tenancy.Value("name").String().Equal(data["name"].(string))
	tenancy.Value("tele").String().Equal(data["tele"].(string))
	tenancy.Value("address").String().Equal(data["address"].(string))
	tenancy.Value("businessTime").String().Equal(data["businessTime"].(string))
	tenancy.Value("sysRegionCode").Number().Equal(data["sysRegionCode"].(int))
	tenancy.Value("status").Number().Equal(data["status"].(int))
	tenancyId := tenancy.Value("id").Number().Raw()

	update := map[string]interface{}{
		"name":          "宝安妇女儿童附属医院",
		"tele":          "0755-235689111",
		"address":       "xxx街道667号",
		"businessTime":  "08:30-17:40",
		"status":        g.StatusTrue,
		"sysRegionCode": 3,
	}

	obj = auth.PUT(fmt.Sprintf("/v1/admin/tenancy/updateTenancy/%f", tenancyId)).
		WithJSON(update).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("更新成功")
	tenancy = obj.Value("data").Object()

	tenancy.Value("name").String().Equal(update["name"].(string))
	tenancy.Value("tele").String().Equal(update["tele"].(string))
	tenancy.Value("address").String().Equal(update["address"].(string))
	tenancy.Value("businessTime").String().Equal(update["businessTime"].(string))
	tenancy.Value("sysRegionCode").Number().Equal(update["sysRegionCode"].(int))
	tenancy.Value("status").Number().Equal(update["status"].(int))

	obj = auth.GET(fmt.Sprintf("/v1/admin/tenancy/getTenancyById/%f", tenancyId)).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
	tenancy = obj.Value("data").Object()

	tenancy.Value("id").Number().Ge(0)
	tenancy.Value("uuid").String().NotEmpty()
	tenancy.Value("name").String().Equal(update["name"].(string))
	tenancy.Value("tele").String().Equal(update["tele"].(string))
	tenancy.Value("address").String().Equal(update["address"].(string))
	tenancy.Value("businessTime").String().Equal(update["businessTime"].(string))
	tenancy.Value("sysRegionCode").Number().Equal(update["sysRegionCode"].(int))
	tenancy.Value("status").Number().Equal(update["status"].(int))

	// setTenancyRegion
	obj = auth.POST("/v1/admin/tenancy/setTenancyRegion").
		WithJSON(map[string]interface{}{
			"id":            tenancyId,
			"sysRegionCode": 2,
		}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("设置成功")

	// changeTenancyStatus
	obj = auth.POST("/v1/admin/tenancy/changeTenancyStatus").
		WithJSON(map[string]interface{}{
			"id":     tenancyId,
			"status": 2,
		}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("设置成功")

	// setUserAuthority
	obj = auth.DELETE(fmt.Sprintf("/v1/admin/tenancy/deleteTenancy/%f", tenancyId)).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("删除成功")

}

func TestTenancyRegisterError(t *testing.T) {
	data := map[string]interface{}{
		"name":          "宝安中心人民医院",
		"tele":          "0755-23568911",
		"address":       "xxx街道666号",
		"businessTime":  "08:30-17:30",
		"status":        g.StatusTrue,
		"sysRegionCode": 1,
	}
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/tenancy/createTenancy").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("添加失败:名称已被注冊")

}
