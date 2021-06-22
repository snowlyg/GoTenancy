package tests

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAttrTemplateList(t *testing.T) {
	auth := tenancyWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/merchant/attrTemplate/getAttrTemplateList").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object().Value("list").Array()

	data.Length().Ge(0)
	first := data.First().Object()
	first.Keys().ContainsOnly(
		"id",
		"createdAt",
		"updatedAt",
		"templateName",
		"templateValue",
		"sysTenancyId",
	)
	first.Value("id").Number().Ge(0)
}

func TestAttrTemplateProcess(t *testing.T) {
	detail := "1"
	value := "sfsdf"
	data := map[string]interface{}{
		"templateName": "fsdaf ",
		"templateValue": []map[string]interface{}{
			{"value": value, "detail": []string{detail}},
		},
	}
	auth := tenancyWithLoginTester(t)
	defer baseLogOut(auth)

	obj := auth.POST("v1/merchant/attrTemplate/createAttrTemplate").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("创建成功")

	attrTemplate := obj.Value("data").Object()
	attrTemplate.Value("id").Number().Ge(0)
	attrTemplate.Value("templateName").String().Equal(data["templateName"].(string))
	attrTemplate.Value("createdAt").String().NotEmpty()
	attrTemplate.Value("updatedAt").String().NotEmpty()
	attrTemplate.Value("sysTenancyId").Number().Equal(1)
	attrTemplate.Value("templateValue").Array().First().Object().Value("value").Equal(value)
	attrTemplate.Value("templateValue").Array().First().Object().Value("detail").Array().First().Equal(detail)

	attrTemplateId := attrTemplate.Value("id").Number().Raw()

	data = map[string]interface{}{
		"templateName": "fsdaf ",
		"templateValue": []map[string]interface{}{
			{"value": value, "detail": []string{detail}},
		},
	}

	obj = auth.PUT(fmt.Sprintf("v1/merchant/attrTemplate/updateAttrTemplate/%d", int(attrTemplateId))).
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("更新成功")

	obj = auth.GET(fmt.Sprintf("v1/merchant/attrTemplate/getAttrTemplateById/%d", int(attrTemplateId))).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
	attrTemplate = obj.Value("data").Object()

	attrTemplate.Value("id").Number().Ge(0)
	attrTemplate.Value("templateName").String().Equal(data["templateName"].(string))
	attrTemplate.Value("createdAt").String().NotEmpty()
	attrTemplate.Value("updatedAt").String().NotEmpty()
	attrTemplate.Value("sysTenancyId").Number().Equal(1)
	attrTemplate.Value("templateValue").Array().First().Object().Value("value").Equal(value)
	attrTemplate.Value("templateValue").Array().First().Object().Value("detail").Array().First().Equal(detail)

	// deleteCategory
	obj = auth.DELETE(fmt.Sprintf("v1/merchant/attrTemplate/deleteAttrTemplate/%d", int(attrTemplateId))).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("删除成功")

}
