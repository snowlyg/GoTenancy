package tests

import (
	"net/http"
	"testing"
)

func TestReceiptList(t *testing.T) {
	auth := generalWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/general/receipt/getReceiptList").
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
	first.Keys().ContainsOnly("id", "receiptType", "receiptTitle", "receiptTitleType", "dutyGaragraph", "email", "bankName", "bankCode", "address", "tel", "isDefault", "sysUserId", "createdAt", "updatedAt")
	first.Value("id").Number().Ge(0)
}

func TestReceiptProcess(t *testing.T) {
	data := map[string]interface{}{
		"receiptType":      1,
		"receiptTitle":     "深圳宝安中心人民医院",
		"receiptTitleType": 1,
		"dutyGaragraph":    "深圳宝安中心人民医院",
		"email":            "8485747@qq.com",
		"bankName":         "中国农业银行",
		"bankCode":         "4564462555641555",
		"address":          "松山湖阿里产业园",
		"tel":              "13845687419",
		"isDefault":        true,
	}
	auth := generalWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/general/receipt/createReceipt").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("创建成功")

	receipt := obj.Value("data").Object()
	receipt.Value("id").Number().Ge(0)
	receipt.Value("receiptTitle").String().Equal(data["receiptTitle"].(string))
	receipt.Value("dutyGaragraph").String().Equal(data["dutyGaragraph"].(string))
	receipt.Value("email").String().Equal(data["email"].(string))
	receipt.Value("bankName").String().Equal(data["bankName"].(string))
	receipt.Value("bankCode").String().Equal(data["bankCode"].(string))
	receipt.Value("address").String().Equal(data["address"].(string))
	receipt.Value("tel").String().Equal(data["tel"].(string))
	receipt.Value("receiptType").Number().Equal(data["receiptType"].(int))
	receipt.Value("receiptTitleType").Number().Equal(data["receiptTitleType"].(int))
	receipt.Value("isDefault").Boolean().Equal(data["isDefault"].(bool))
	receiptId := receipt.Value("id").Number().Raw()
	if receiptId > 0 {
		update := map[string]interface{}{
			"id":               receiptId,
			"receiptType":      2,
			"receiptTitle":     "深圳宝安中心人民医院",
			"receiptTitleType": 2,
			"dutyGaragraph":    "深圳宝安中心人民医院",
			"email":            "8485747@qq.com",
			"bankName":         "中国农业银行",
			"bankCode":         "4564462555641555",
			"address":          "松山湖阿里产业园",
			"tel":              "13845687419",
			"isDefault":        true,
		}

		obj = auth.PUT("v1/general/receipt/updateReceipt").
			WithJSON(update).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("更新成功")
		receipt = obj.Value("data").Object()

		receipt.Value("id").Number().Ge(0)
		receipt.Value("receiptTitle").String().Equal(update["receiptTitle"].(string))
		receipt.Value("dutyGaragraph").String().Equal(update["dutyGaragraph"].(string))
		receipt.Value("email").String().Equal(update["email"].(string))
		receipt.Value("bankName").String().Equal(update["bankName"].(string))
		receipt.Value("bankCode").String().Equal(update["bankCode"].(string))
		receipt.Value("address").String().Equal(update["address"].(string))
		receipt.Value("tel").String().Equal(update["tel"].(string))
		receipt.Value("receiptType").Number().Equal(update["receiptType"].(int))
		receipt.Value("receiptTitleType").Number().Equal(update["receiptTitleType"].(int))
		receipt.Value("isDefault").Boolean().Equal(update["isDefault"].(bool))

		obj = auth.POST("v1/general/receipt/getReceiptById").
			WithJSON(map[string]interface{}{"id": receiptId}).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("操作成功")
		receipt = obj.Value("data").Object()

		receipt.Value("id").Number().Ge(0)
		receipt.Value("receiptTitle").String().Equal(update["receiptTitle"].(string))
		receipt.Value("dutyGaragraph").String().Equal(update["dutyGaragraph"].(string))
		receipt.Value("email").String().Equal(update["email"].(string))
		receipt.Value("bankName").String().Equal(update["bankName"].(string))
		receipt.Value("bankCode").String().Equal(update["bankCode"].(string))
		receipt.Value("address").String().Equal(update["address"].(string))
		receipt.Value("tel").String().Equal(update["tel"].(string))
		receipt.Value("receiptType").Number().Equal(update["receiptType"].(int))
		receipt.Value("receiptTitleType").Number().Equal(update["receiptTitleType"].(int))
		receipt.Value("isDefault").Boolean().Equal(update["isDefault"].(bool))

		// setUserAuthority
		obj = auth.DELETE("v1/general/receipt/deleteReceipt").
			WithJSON(map[string]interface{}{"id": receiptId}).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("删除成功")
	}

}
func TestReceiptRegisterReceiptTitleError(t *testing.T) {
	data := map[string]interface{}{
		"receiptType":      1,
		"receiptTitle":     "",
		"receiptTitleType": 1,
		"dutyGaragraph":    "深圳宝安中心人民医院",
		"email":            "8485747@qq.com",
		"bankName":         "中国农业银行",
		"bankCode":         "4564462555641555",
		"address":          "松山湖阿里产业园",
		"tel":              "13845687419",
		"isDefault":        true,
	}
	auth := generalWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/general/receipt/createReceipt").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("Key: 'CreateReceipt.ReceiptTitle' Error:Field validation for 'ReceiptTitle' failed on the 'required' tag")
}
func TestReceiptRegisterReceiptTypeError(t *testing.T) {
	data := map[string]interface{}{
		"receiptType":      3,
		"receiptTitle":     "深圳宝安中心人民医院",
		"receiptTitleType": 1,
		"dutyGaragraph":    "深圳宝安中心人民医院",
		"email":            "8485747@qq.com",
		"bankName":         "中国农业银行",
		"bankCode":         "4564462555641555",
		"address":          "松山湖阿里产业园",
		"tel":              "13845687419",
		"isDefault":        true,
	}
	auth := generalWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/general/receipt/createReceipt").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("Key: 'CreateReceipt.ReceiptType' Error:Field validation for 'ReceiptType' failed on the 'oneof' tag")

}
func TestReceiptRegisterEmaileError(t *testing.T) {
	data := map[string]interface{}{
		"receiptType":      2,
		"receiptTitle":     "深圳宝安中心人民医院",
		"receiptTitleType": 1,
		"dutyGaragraph":    "深圳宝安中心人民医院",
		"email":            "8485747qq.com",
		"bankName":         "中国农业银行",
		"bankCode":         "4564462555641555",
		"address":          "松山湖阿里产业园",
		"tel":              "13845687419",
		"isDefault":        true,
	}
	auth := generalWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/general/receipt/createReceipt").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(4000)
	obj.Value("message").String().Equal("Key: 'CreateReceipt.Email' Error:Field validation for 'Email' failed on the 'email' tag")

}
