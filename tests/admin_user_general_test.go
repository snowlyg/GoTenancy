package tests

import (
	"net/http"
	"testing"
)

func TestGeneralUserList(t *testing.T) {
	params := []param{
		{args: map[string]interface{}{"page": 1, "pageSize": 10}, length: 2},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "groupId": "1"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "groupId": "2"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "groupId": "3"}, length: 0},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "groupId": "2", "labelId": "2"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "groupId": "2", "labelId": "3"}, length: 0},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "groupId": "1", "labelId": "1", "nickName": "C端用户2"}, length: 0},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "groupId": "2", "labelId": "2", "nickName": "C端用户1"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "groupId": "1", "labelId": "1", "nickName": "C端用户"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "groupId": "1", "labelId": "1", "nickName": "C端用户", "sex": "2"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "groupId": "2", "labelId": "2", "nickName": "C端用户", "sex": "1"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "groupId": "1", "labelId": "1", "nickName": "C端用户", "sex": "0"}, length: 0},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "groupId": "1", "labelId": "1", "nickName": "C端用户", "payCount": "0"}, length: 0},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "groupId": "1", "labelId": "1", "nickName": "C端用户", "payCount": "5"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "groupId": "1", "labelId": "1", "nickName": "C端用户", "payCount": "2"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "groupId": "2", "labelId": "2", "nickName": "C端用户", "userType": "wechat"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "groupId": "1", "labelId": "1", "nickName": "C端用户", "userType": "routine"}, length: 1},
	}
	for _, param := range params {
		userGeneralTest(t, param.args, param.length)
	}
}

func userGeneralTest(t *testing.T, params map[string]interface{}, length int) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/user/getGeneralList").
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
			"createdAt",
			"updatedAt",
			"userName",
			"authorityName",
			"authorityType",
			"authorityId",
			"groupName",
			"labelId",
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
			"addres",
			"lastTime",
			"lastIp",
			"nowMoney",
			"userType",
			"mainUid",
			"payCount",
			"payPrice",
		)
		first.Value("id").Number().Ge(0)
	}
}
