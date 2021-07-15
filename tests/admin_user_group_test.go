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
