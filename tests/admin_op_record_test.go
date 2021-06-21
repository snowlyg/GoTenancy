package tests

import (
	"net/http"
	"testing"
)

func TestOperationRecode(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/sysOperationRecord/getSysOperationRecordList").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10}).
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
	first.Keys().ContainsOnly("id", "ip", "method", "path", "status", "latency", "agent", "errorMessage", "body", "resp", "userId", "tenancyName", "userName", "nickName", "createdAt", "updatedAt")
	first.Value("id").Number().Ge(0)

}
