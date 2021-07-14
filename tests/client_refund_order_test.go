package tests

import (
	"fmt"
	"net/http"
	"testing"
)

func TestClientRefundOrderList(t *testing.T) {
	params := []param{
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": ""}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "1"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "2"}, length: 0},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "3"}, length: 0},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "4"}, length: 0},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "5"}, length: 0},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "6"}, length: 0},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "7"}, length: 0},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "", "date": "today"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "", "date": "yesterday"}, length: 0},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "", "date": "lately7"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "", "date": "lately30"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "", "date": "month"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "", "date": "year"}, length: 1},
	}
	for _, param := range params {
		refundOrderClientlist(t, param.args, param.length)
	}
}

func refundOrderClientlist(t *testing.T, params map[string]interface{}, length int) {
	auth := tenancyWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/merchant/refundOrder/getRefundOrderList").
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
			"id",
			"createdAt",
			"updatedAt",
			"refundOrderSn",
			"deliveryType",
			"deliveryId",
			"deliveryMark",
			"deliveryPics",
			"deliveryPhone",
			"merDeliveryUser",
			"merDeliveryAddress",
			"phone",
			"mark",
			"merMark",
			"adminMark",
			"pics",
			"refundType",
			"refundMessage",
			"refundPrice",
			"refundNum",
			"failMessage",
			"status",
			"statusTime",
			"isDel",
			"isSystemDel",
			"orderSn",
			"userNickName",
			"tenancyName",
			"isTrader",
			"reconciliationId",
			"orderId",
			"sysUserId",
			"sysTenancyId",
			"refundProduct",
		)
		first.Value("id").Number().Ge(0)
	}
}

func TestClientRefundOrderRecord(t *testing.T) {
	orderId := 1
	auth := tenancyWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST(fmt.Sprintf("v1/merchant/order/getRefundOrderRecord/%d", orderId)).
		WithJSON(map[string]interface{}{
			"page":     1,
			"pageSize": 10,
		}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")
	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize")
	data.Value("pageSize").Number().Equal(10)
	data.Value("page").Number().Equal(1)
}

func TestClientRefundOrderRemark(t *testing.T) {
	orderId := 1
	auth := tenancyWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.GET(fmt.Sprintf("v1/merchant/order/getRefundOrderRemarkMap/%d", orderId)).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")

	obj = auth.POST(fmt.Sprintf("v1/merchant/order/remarkRefundOrder/%d", orderId)).
		WithJSON(map[string]interface{}{"mer_mark": "remark"}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
}

func TestClientRefundOrderAudit(t *testing.T) {
	orderId := 1
	auth := tenancyWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.GET(fmt.Sprintf("v1/merchant/order/getRefundOrderMap/%d", orderId)).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")

	obj = auth.POST(fmt.Sprintf("v1/merchant/order/auditRefundOrder/%d", orderId)).
		WithJSON(map[string]interface{}{"status": 1}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
}
