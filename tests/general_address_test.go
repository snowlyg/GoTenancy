package tests

import (
	"testing"

	"github.com/kataras/iris/v12/httptest"
)

func TestAddressList(t *testing.T) {
	auth := generalWithLoginTester(t)
	obj := auth.POST("/v1/general/address/getAddressList").
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
	first.Keys().ContainsOnly("id", "name", "phone", "sex", "country", "province", "city", "district", "detail", "isDefault", "postcode", "age", "hospitalName", "locName", "bedNum", "hospitalNo", "disease", "createdAt", "updatedAt")
	first.Value("id").Number().Ge(0)

	baseLogOut(auth)
}

func TestAddressProcess(t *testing.T) {
	data := map[string]interface{}{
		"name":         "八两金",
		"phone":        "13845687419",
		"sex":          2,
		"country":      "中国",
		"province":     "广东省",
		"city":         "东莞市",
		"district":     "寮步镇",
		"isDefault":    1,
		"detail":       "松山湖阿里产业园",
		"postcode":     "413514",
		"age":          32,
		"hospitalName": "深圳宝安中心人民医院",
		"locName":      "泌尿科一区",
		"bedNum":       "15",
		"hospitalNo":   "88956655",
		"disease":      "不孕不育",
	}
	auth := generalWithLoginTester(t)
	obj := auth.POST("/v1/general/address/createAddress").
		WithJSON(data).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("创建成功")

	address := obj.Value("data").Object()
	address.Value("id").Number().Ge(0)
	address.Value("name").String().Equal(data["name"].(string))
	address.Value("phone").String().Equal(data["phone"].(string))
	address.Value("sex").Number().Equal(data["sex"].(int))
	address.Value("country").String().Equal(data["country"].(string))
	address.Value("province").String().Equal(data["province"].(string))
	address.Value("city").String().Equal(data["city"].(string))
	address.Value("district").String().Equal(data["district"].(string))
	address.Value("detail").String().Equal(data["detail"].(string))
	address.Value("postcode").String().Equal(data["postcode"].(string))
	address.Value("isDefault").Number().Equal(data["isDefault"].(int))
	address.Value("age").Number().Equal(data["age"].(int))
	address.Value("hospitalName").String().Equal(data["hospitalName"].(string))
	address.Value("locName").String().Equal(data["locName"].(string))
	address.Value("bedNum").String().Equal(data["bedNum"].(string))
	address.Value("hospitalNo").String().Equal(data["hospitalNo"].(string))
	address.Value("disease").String().Equal(data["disease"].(string))
	addressId := address.Value("id").Number().Raw()

	update := map[string]interface{}{
		"id":           addressId,
		"name":         "八两金1",
		"phone":        "138456874191",
		"sex":          2,
		"country":      "中国1",
		"province":     "广东省1",
		"city":         "东莞市1",
		"district":     "寮步镇1",
		"isDefault":    1,
		"detail":       "松山湖阿里产业园1",
		"postcode":     "4135141",
		"age":          32,
		"hospitalName": "深圳宝安中心人民医院1",
		"locName":      "泌尿科一区1",
		"bedNum":       "151",
		"hospitalNo":   "889566551",
		"disease":      "不孕不育1",
	}

	obj = auth.PUT("/v1/general/address/updateAddress").
		WithJSON(update).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("更新成功")
	address = obj.Value("data").Object()

	address.Value("id").Number().Ge(0)
	address.Value("name").String().Equal(data["name"].(string))
	address.Value("phone").String().Equal(data["phone"].(string))
	address.Value("sex").Number().Equal(data["sex"].(int))
	address.Value("country").String().Equal(data["country"].(string))
	address.Value("province").String().Equal(data["province"].(string))
	address.Value("city").String().Equal(data["city"].(string))
	address.Value("district").String().Equal(data["district"].(string))
	address.Value("detail").String().Equal(data["detail"].(string))
	address.Value("postcode").String().Equal(data["postcode"].(string))
	address.Value("isDefault").Number().Equal(data["isDefault"].(int))
	address.Value("age").Number().Equal(data["age"].(int))
	address.Value("hospitalName").String().Equal(data["hospitalName"].(string))
	address.Value("locName").String().Equal(data["locName"].(string))
	address.Value("bedNum").String().Equal(data["bedNum"].(string))
	address.Value("hospitalNo").String().Equal(data["hospitalNo"].(string))
	address.Value("disease").String().Equal(data["disease"].(string))

	obj = auth.POST("/v1/general/address/getAddressById").
		WithJSON(map[string]interface{}{"id": addressId}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("操作成功")
	address = obj.Value("data").Object()

	address.Value("id").Number().Ge(0)
	address.Value("uuid").String().NotEmpty()
	address.Value("name").String().Equal(update["name"].(string))
	address.Value("appId").String().Equal(update["appId"].(string))
	address.Value("appSecret").String().Equal(update["appSecret"].(string))
	address.Value("remark").String().Equal(update["remark"].(string))

	// setUserAuthority
	obj = auth.DELETE("/v1/general/address/deleteAddress").
		WithJSON(map[string]interface{}{"id": addressId}).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("删除成功")

	baseLogOut(auth)
}

func TestAddressRegisterError(t *testing.T) {
	data := map[string]interface{}{
		"name":         "",
		"phone":        "13845687419",
		"sex":          2,
		"country":      "中国",
		"province":     "广东省",
		"city":         "东莞市",
		"district":     "寮步镇",
		"isDefault":    1,
		"detail":       "松山湖阿里产业园",
		"postcode":     "413514",
		"age":          32,
		"hospitalName": "深圳宝安中心人民医院",
		"locName":      "泌尿科一区",
		"bedNum":       "15",
		"hospitalNo":   "88956655",
		"disease":      "不孕不育",
	}
	auth := generalWithLoginTester(t)
	obj := auth.POST("/v1/general/address/createAddress").
		WithJSON(data).
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(4000)
	obj.Value("msg").String().Equal("Key: 'CreateAddress.Name' Error:Field validation for 'Name' failed on the 'required' tag")

	baseLogOut(auth)
}
