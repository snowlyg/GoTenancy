package tests

import (
	"net/http"
	"testing"
)

func TestRefundOrderList(t *testing.T) {
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
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "", "date": "year", "isTrader": "1"}, length: 0},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "status": "", "date": "year", "isTrader": "2"}, length: 1},
	}
	for _, param := range params {
		refundOrderlist(t, param.args, param.length)
	}
}

func refundOrderlist(t *testing.T, params map[string]interface{}, length int) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/refundOrder/getRefundOrderList").
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
