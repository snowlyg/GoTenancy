package tests

import (
	"fmt"
	"net/http"
	"testing"
)

func TestUserGroupList(t *testing.T) {

	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/userGroup/getUserGroupList").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize")
	data.Value("pageSize").Number().Equal(10)
	data.Value("page").Number().Equal(1)
	data.Value("total").Number().Equal(2)

	list := data.Value("list").Array()
	list.Length().Ge(0)
	first := list.First().Object()
	first.Keys().ContainsOnly("id", "groupName", "createdAt", "updatedAt")
}

func TestUserEditUsers(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.GET(fmt.Sprintf("v1/admin/cuser/editUserMap/%d", cuserId)).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := map[string]interface{}{
		"address":  "address1",
		"birthday": "2021-07-16",
		"groupId":  2,
		"idCard":   "445281199411285863",
		"labelId":  []int{1},
		"mark":     "mark1",
		"phone":    "13800138001",
		"realName": "余思琳1",
		"uid":      3,
	}
	obj = auth.POST(fmt.Sprintf("v1/admin/cuser/editUser/%d", cuserId)).
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
}
func TestUserSetNowMoney(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.GET(fmt.Sprintf("v1/admin/cuser/setNowMoneyMap/%d", cuserId)).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := map[string]interface{}{
		"nowMoney": 2,
		"type":     2,
	}
	obj = auth.POST(fmt.Sprintf("v1/admin/cuser/setNowMoney/%d", cuserId)).
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
}
func TestUserSetUserGroup(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.GET(fmt.Sprintf("v1/admin/cuser/setUserGroupMap/%d", cuserId)).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := map[string]interface{}{
		"id":  2,
		"ids": []int{0},
	}
	obj = auth.POST(fmt.Sprintf("v1/admin/cuser/setUserGroup/%d", cuserId)).
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
}
func TestUserSetUserLabel(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.GET(fmt.Sprintf("v1/admin/cuser/setUserLabelMap/%d", cuserId)).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := map[string]interface{}{
		"ids":      []int{3},
		"label_id": []int{1, 2, 3, 4},
	}
	obj = auth.POST(fmt.Sprintf("v1/admin/cuser/setUserLabel/%d", cuserId)).
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
}
func TestUserBatchSetUserLabel(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/cuser/batchSetUserLabelMap").
		WithJSON(map[string]interface{}{
			"ids": "1,2",
		}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := map[string]interface{}{
		"id":  1,
		"ids": []int{1, 2},
	}
	obj = auth.POST("v1/admin/cuser/batchSetUserLabel").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("设置成功")
}
func TestUserBatchSetUserGroup(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/cuser/batchSetUserGroupMap").
		WithJSON(map[string]interface{}{
			"ids": "1,2",
		}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := map[string]interface{}{
		"id":  1,
		"ids": []int{1, 2},
	}
	obj = auth.POST("v1/admin/cuser/batchSetUserGroup").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("设置成功")
}
