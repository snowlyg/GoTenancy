package tests

import (
	"fmt"
	"net/http"
	"testing"
)

var cuserId = 3

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
	obj := auth.POST("v1/admin/cuser/getGeneralList").
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

func TestUserGetOrderList(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST(fmt.Sprintf("v1/admin/cuser/getOrderList/%d", cuserId)).
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

func TestUserGetBillList(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST(fmt.Sprintf("v1/admin/cuser/getBillList/%d", cuserId)).
		WithJSON(map[string]interface{}{
			"page":     1,
			"pageSize": 10,
		}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	res := obj.Value("data").Object()
	res.Keys().ContainsOnly("list", "total", "page", "pageSize")
	res.Value("pageSize").Number().Equal(10)
	res.Value("page").Number().Equal(1)
	res.Value("total").Number().Equal(4)
}

func TestUserGetGeneralDetail(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.GET(fmt.Sprintf("v1/admin/cuser/getGeneralDetail/%d", cuserId)).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")
	res := obj.Value("data").Object()
	res.Value("uid").Number().Equal(cuserId)
	res.Value("nowMoney").Number().Equal(0)
	res.Value("payCount").Number().Equal(5)
	res.Value("payPrice").Number().Equal(20)
	res.Value("totalPayCount").Number().Equal(5)
	res.Value("totalPayPrice").Number().Equal(673)
	res.Value("groupId").Number().Equal(2)
	res.Value("labelId").Array().First().Equal(1)
	res.Value("avatarUrl").String().Equal("https://thirdwx.qlogo.cn/mmopen/vi_32/PEyYoZmTJtaJdeYWWibrnDUadmXKVYyTtyRq2nxtWbBic5jJTLTT4KHmox1tNvOicgIXxspgmxicghpCFob1icAIWFw/132")
	res.Value("nickName").String().Equal("C端用户")
	res.Value("idCard").String().Equal("445281199411285863")
	res.Value("realName").String().Equal("余思琳1")
	res.Value("birthday").String().Equal("2021-07-16")
	res.Value("mark").String().Equal("mark1")
	res.Value("address").String().Equal("address1")
	res.Value("phone").String().Equal("13800138001")
}

func TestUserGroupProcess(t *testing.T) {
	data := map[string]interface{}{
		"groupName": "sdfsdfs34234",
	}
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/userGroup/createUserGroup").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("创建成功")

	userGroup := obj.Value("data").Object()
	userGroup.Value("id").Number().Ge(0)
	userGroup.Value("groupName").String().Equal(data["groupName"].(string))
	userGroupId := userGroup.Value("id").Number().Raw()
	if userGroupId > 0 {
		update := map[string]interface{}{
			"groupName": "sdfsdfs213213",
		}
		obj = auth.PUT(fmt.Sprintf("v1/admin/userGroup/updateUserGroup/%d", int(userGroupId))).
			WithJSON(update).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("更新成功")
		userGroup = obj.Value("data").Object()

		userGroup.Value("id").Number().Ge(0)
		userGroup.Value("groupName").String().Equal(update["groupName"].(string))

		obj = auth.GET("v1/admin/userGroup/getCreateUserGroupMap").
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		obj = auth.GET(fmt.Sprintf("v1/admin/userGroup/getUpdateUserGroupMap/%d", int(userGroupId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		// deleteUserGroup
		obj = auth.DELETE(fmt.Sprintf("v1/admin/userGroup/deleteUserGroup/%d", int(userGroupId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("删除成功")
	}
}
