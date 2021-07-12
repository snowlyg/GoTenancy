package tests

import (
	"fmt"
	"net/http"
	"testing"
)

func TestClientOrderList(t *testing.T) {
	params := []param{
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "0"}, length: 5},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "1"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "2"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "3"}, length: 0},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "4"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "5"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "6"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "7"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "0", "date": "today"}, length: 5},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "0", "date": "yesterday"}, length: 0},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "0", "date": "lately7"}, length: 5},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "0", "date": "lately30"}, length: 5},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "0", "date": "month"}, length: 5},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "0", "date": "year"}, length: 5},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "0", "date": "year", "isTrader": "1"}, length: 0},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "0", "date": "year", "isTrader": "2"}, length: 5},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "0", "date": "year", "isTrader": "2", "sysTenancyId": 1}, length: 5},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "0", "date": "year", "isTrader": "2", "sysTenancyId": 2}, length: 0},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "0", "date": "year", "isTrader": "2", "sysTenancyId": 1, "keywords": "real_name"}, length: 1},
	}
	for _, param := range params {
		orderClientlist(t, param.args, param.length)
	}
}

func orderClientlist(t *testing.T, params map[string]interface{}, length int) {
	auth := tenancyWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/merchant/order/getOrderList").
		WithJSON(params).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize", "stat")
	data.Value("pageSize").Number().Equal(10)
	data.Value("page").Number().Equal(1)
	data.Value("total").Number().Equal(length)

	if length > 0 {
		list := data.Value("list").Array()
		list.Length().Ge(0)
		first := list.First().Object()
		first.Keys().ContainsOnly(
			"updatedAt",
			"userAddress",
			"totalNum",
			"commissionRate",
			"paid",
			"realName",
			"orderType",
			"payTime",
			"payType",
			"deliveryId",
			"isSystemDel",
			"sysTenancyId",
			"id",
			"orderSn",
			"payPostage",
			"status",
			"deliveryType",
			"groupOrderId",
			"remark",
			"userPhone",
			"totalPostage",
			"deliveryName",
			"adminMark",
			"activityType",
			"tenancyName",
			"mark",
			"verifyCode",
			"cost",
			"isTrader",
			"verifyTime",
			"isDel",
			"groupOrderSn",
			"reconciliationId",
			"createdAt",
			"totalPrice",
			"payPrice",
			"sysUserId",
			"cartId",
			"orderProduct",
		)
		first.Value("id").Number().Ge(0)
	}
}

func TestGetClientOrderChart(t *testing.T) {
	auth := tenancyWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.GET("v1/merchant/order/getOrderChart").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
}

func TestGetClientOrderFilter(t *testing.T) {
	auth := tenancyWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.GET("v1/merchant/order/getOrderFilter").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
}

func TestClientOrderDetail(t *testing.T) {
	orderId := 1
	auth := tenancyWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.GET(fmt.Sprintf("v1/merchant/order/getOrderById/%d", orderId)).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")

	obj = auth.POST(fmt.Sprintf("v1/merchant/order/getOrderRecord/%d", orderId)).
		WithJSON(map[string]interface{}{
			"page":     1,
			"pageSize": 10,
		}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")
	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize", "stat")
	data.Value("pageSize").Number().Equal(10)
	data.Value("page").Number().Equal(1)

	obj = auth.GET(fmt.Sprintf("v1/merchant/order/deliveryOrderMap/%d", orderId)).
		WithJSON(map[string]interface{}{"remark": "remark"}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")

	obj = auth.POST(fmt.Sprintf("v1/merchant/order/deliveryOrder/%d", orderId)).
		WithJSON(map[string]interface{}{
			"deliveryId":   "13412081338",
			"deliveryName": 34,
			"deliveryType": 1,
		}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")

	obj = auth.GET(fmt.Sprintf("v1/merchant/order/getOrderRemarkMap/%d", orderId)).
		WithJSON(map[string]interface{}{"remark": "remark"}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")

	obj = auth.POST(fmt.Sprintf("v1/merchant/order/remarkOrder/%d", orderId)).
		WithJSON(map[string]interface{}{"remark": "remark"}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
}
