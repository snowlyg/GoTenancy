package tests

import (
	"fmt"
	"net/http"
	"testing"
)

var cuserClientId = 3

func TestClientGeneralUserList(t *testing.T) {
	params := []param{
		{args: map[string]interface{}{"page": 1, "pageSize": 10}, length: 2},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "labelId": "2"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "labelId": "3"}, length: 0},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "labelId": "1", "nickName": "C端用户2"}, length: 0},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "labelId": "2", "nickName": "C端用户1"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "labelId": "1", "nickName": "C端用户"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "labelId": "1", "nickName": "C端用户", "sex": "2"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "labelId": "2", "nickName": "C端用户", "sex": "1"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "labelId": "1", "nickName": "C端用户", "sex": "0"}, length: 0},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "labelId": "1", "nickName": "C端用户", "payCount": "0"}, length: 0},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "labelId": "1", "nickName": "C端用户", "payCount": "5"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "labelId": "1", "nickName": "C端用户", "payCount": "2"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "labelId": "2", "nickName": "C端用户", "userType": "wechat"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "labelId": "1", "nickName": "C端用户", "userType": "routine"}, length: 1},
	}
	for _, param := range params {
		userClientGeneralTest(t, param.args, param.length)
	}
}

func userClientGeneralTest(t *testing.T, params map[string]interface{}, length int) {
	auth := tenancyWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/merchant/cuser/getGeneralList").
		WithJSON(params).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize")
	data.Value("pageSize").Number().Equal(10)
	data.Value("page").Number().Equal(1)
	data.Value("total").Number().Equal(length)
	if length > 0 {
		data.Value("list").Array().Length().Ge(0)
		list := data.Value("list").Array()
		list.Length().Ge(0)
		first := list.First().Object()
		first.Keys().ContainsOnly(
			"id",
			"uid",
			"createdAt",
			"updatedAt",
			"userName",
			"authorityName",
			"authorityType",
			"authorityId",
			"groupName",
			"label",
			"email",
			"phone",
			"nickName",
			"avatarUrl",
			"sex",
			"subscribe",
			"openId",
			"unionId",
			"country",
			"province",
			"city",
			"idCard",
			"isAuth",
			"realName",
			"birthday",
			"mark",
			"address",
			"lastTime",
			"lastIp",
			"nowMoney",
			"userType",
			"mainUid",
			"payCount",
			"firstPayTime",
			"lastPayTime",
			"payPrice",
		)
		first.Value("id").Number().Ge(0)
	}
}

func TestClientUserGetOrderList(t *testing.T) {
	auth := tenancyWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST(fmt.Sprintf("v1/merchant/cuser/getOrderList/%d", cuserClientId)).
		WithJSON(map[string]interface{}{
			"page":     1,
			"pageSize": 10,
		}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	res := obj.Value("data").Object()
	res.Keys().ContainsOnly("list", "total", "page", "pageSize", "stat")
	res.Value("pageSize").Number().Equal(10)
	res.Value("page").Number().Equal(1)
	res.Value("total").Number().Equal(5)
}

func TestClientUserSetUserLabel(t *testing.T) {
	auth := tenancyWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.GET(fmt.Sprintf("v1/merchant/cuser/setUserLabelMap/%d", cuserClientId)).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := map[string]interface{}{
		"ids":      []int{3},
		"label_id": []int{1, 2, 3, 4},
	}
	obj = auth.POST(fmt.Sprintf("v1/merchant/cuser/setUserLabel/%d", cuserClientId)).
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
}
