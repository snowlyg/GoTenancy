package tests

import (
	"fmt"
	"net/http"
	"testing"
)

func TestClinetProductList(t *testing.T) {
	params := []param{
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "1"}, length: 3},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "1", "keyword": "领立"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "1", "isGiftBag": "1"}, length: 0},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "1", "cateId": 185}, length: 0},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "1", "tenancyCategoryId": 173}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "2"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "3"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "4"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "5"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "6"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "7"}, length: 1},
	}
	for _, param := range params {
		clinetProductList(t, param.args, param.length)
	}
}

func clinetProductList(t *testing.T, params map[string]interface{}, length int) {
	auth := tenancyWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/merchant/product/getProductList").
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
		list := data.Value("list").Array()
		list.Length().Ge(0)
		first := list.First().Object()
		first.Keys().ContainsOnly(
			"id",
			"storeName",
			"storeInfo",
			"keyword",
			"barCode",
			"isShow",
			"status",
			"unitName",
			"sort",
			"rank",
			"sales",
			"price",
			"cost",
			"otPrice",
			"stock",
			"isHot",
			"isBenefit",
			"isBest",
			"isNew",
			"isGood",
			"productType",
			"ficti",
			"browse",
			"codePath",
			"videoLink",
			"specType",
			"extensionType",
			"refusal",
			"rate",
			"replyCount",
			"isGiftBag",
			"careCount",
			"image",
			"oldId",
			"tempId",
			"sysTenancyId",
			"sysTenancyName",
			"cateName",
			"productCates",
			"brandName",
			"sysBrandId",
			"productCategoryId",
			"createdAt",
			"updatedAt",
		)
		first.Value("id").Number().Ge(0)
	}

}

func TestGetClientProductFilter(t *testing.T) {
	auth := tenancyWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.GET("v1/merchant/product/getProductFilter").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

}

func TestClinetProductProcess(t *testing.T) {
	data := map[string]interface{}{
		"attr": []map[string]interface{}{
			{
				"detail": []string{"S",
					"L",
					"XL",
					"XXL",
				},
				"value": "尺寸",
			},
		},
		"attrValue": []map[string]interface{}{
			{
				"image":        "http://127.0.0.1:8089/uploads/file/b39024efbc6de61976f585c8421c6bba_20210702150027.png",
				"barCode":      "",
				"brokerage":    1,
				"brokerageTwo": 1,
				"cost":         1,
				"detail": map[string]interface{}{
					"尺寸": "S",
				},
				"otPrice": 1,
				"price":   1,
				"stock":   1,
				"value0":  "S",
				"volume":  1,
				"weight":  1,
			},
		},
		"cateId":        183,
		"content":       "<p>是的发生的发sad</p>",
		"extensionType": 2,
		"image":         "http://127.0.0.1:8089/uploads/file/b39024efbc6de61976f585c8421c6bba_20210702150027.png",
		"isGiftBag":     2,
		"isGood":        1,
		"keyword":       "sdfdsfsdfsdf",
		"sliderImages": []string{
			"http://127.0.0.1:8089/uploads/file/b39024efbc6de61976f585c8421c6bba_20210702150027.png",
			"http://127.0.0.1:8089/uploads/file/0701aa317da5a004fbf6111545678a6c_20210702150036.png",
		},
		"sort":              1,
		"specType":          1,
		"storeInfo":         "的是否是否",
		"storeName":         "是防守打法发",
		"sysBrandId":        3,
		"tempId":            2,
		"tenancyCategoryId": []int{174},
		"unitName":          "放松的方式",
		"videoLink":         "sdfsdfsd",
		"barCode":           "sdfsdfsd",
	}
	auth := tenancyWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/merchant/product/createProduct").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("创建成功")

	product := obj.Value("data").Object()
	product.Value("id").Number().Ge(0)
	product.Value("storeName").String().Equal(data["storeName"].(string))
	product.Value("storeInfo").String().Equal(data["storeInfo"].(string))
	product.Value("keyword").String().Equal(data["keyword"].(string))
	product.Value("barCode").String().Equal(data["barCode"].(string))
	product.Value("isShow").Number().Equal(2)
	product.Value("status").Number().Equal(2)
	product.Value("unitName").String().Equal(data["unitName"].(string))
	product.Value("sort").Number().Equal(data["sort"].(int))
	product.Value("rank").Number().Equal(0)
	product.Value("sales").Number().Equal(0)
	product.Value("price").Number().Equal(0)
	product.Value("cost").Number().Equal(0)
	product.Value("otPrice").Number().Equal(0)
	product.Value("stock").Number().Equal(0)
	product.Value("isHot").Number().Equal(2)
	product.Value("isBenefit").Number().Equal(2)
	product.Value("isBest").Number().Equal(2)
	product.Value("isNew").Number().Equal(2)
	product.Value("isGood").Number().Equal(data["isGood"].(int))
	product.Value("productType").Number().Equal(1)
	product.Value("ficti").Number().Equal(0)
	product.Value("browse").Number().Equal(0)
	product.Value("codePath").String().Equal("")
	product.Value("videoLink").String().Equal(data["videoLink"].(string))
	product.Value("specType").Number().Equal(data["specType"].(int))
	product.Value("extensionType").Number().Equal(data["extensionType"].(int))
	product.Value("refusal").String().Equal("")
	product.Value("rate").Number().Equal(5)
	product.Value("replyCount").Number().Equal(0)
	product.Value("isGiftBag").Number().Equal(2)
	product.Value("careCount").Number().Equal(0)
	product.Value("image").String().NotEmpty()
	product.Value("sliderImage").String().NotEmpty()
	product.Value("oldId").Number().Equal(0)
	product.Value("tempId").Number().Equal(data["tempId"].(int))
	product.Value("sysBrandId").Number().Equal(data["sysBrandId"].(int))
	product.Value("productCategoryId").Number().Equal(data["cateId"].(int))
	productId := product.Value("id").Number().Raw()
	if productId > 0 {

		update := map[string]interface{}{
			"attr": []map[string]interface{}{
				{
					"detail": []string{"S",
						"L",
						"XL",
						"XXL",
					},
					"value": "尺寸",
				},
			},
			"attrValue": []map[string]interface{}{
				{
					"image":        "http://127.0.0.1:8089/uploads/file/b39024efbc6de61976f585c8421c6bba_20210702150027.png",
					"barCode":      "",
					"brokerage":    1,
					"brokerageTwo": 1,
					"cost":         1,
					"detail": map[string]interface{}{
						"尺寸": "S",
					},
					"otPrice": 1,
					"price":   1,
					"stock":   1,
					"value0":  "S",
					"volume":  1,
					"weight":  1,
				},
			},
			"cateId":        183,
			"content":       "<p>是的发生的发sadsdfsdfsdf</p>",
			"extensionType": 1,
			"image":         "http://127.0.0.1:8089/uploads/file/b39024efbc6de61976f585c8421c6bba_20210702150027.png",
			"isGiftBag":     1,
			"isGood":        2,
			"keyword":       "sdfdsfsdfsdf",
			"sliderImages": []string{
				"http://127.0.0.1:8089/uploads/file/b39024efbc6de61976f585c8421c6bba_20210702150027.png",
				"http://127.0.0.1:8089/uploads/file/0701aa317da5a004fbf6111545678a6c_20210702150036.png",
			},
			"sort":              21321,
			"specType":          2,
			"storeInfo":         "的是否是否",
			"storeName":         "是防守打法发",
			"sysBrandId":        3,
			"tempId":            2,
			"tenancyCategoryId": []int{174},
			"unitName":          "放松的方式213123",
			"videoLink":         "sdfsdfsd11",
			"barCode":           "sdfsdfsd11",
		}

		obj = auth.PUT(fmt.Sprintf("v1/merchant/product/updateProduct/%d", int(productId))).
			WithJSON(update).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("更新成功")

		obj = auth.GET(fmt.Sprintf("v1/merchant/product/getProductById/%d", int(productId))).
			WithJSON(map[string]interface{}{"id": productId}).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("操作成功")
		product = obj.Value("data").Object()

		product.Value("id").Number().Ge(0)
		product.Value("storeName").String().Equal(data["storeName"].(string))
		product.Value("storeInfo").String().Equal(data["storeInfo"].(string))
		product.Value("keyword").String().Equal(data["keyword"].(string))
		product.Value("barCode").String().Equal(data["barCode"].(string))
		product.Value("isShow").Number().Equal(2)
		product.Value("status").Number().Equal(2)
		product.Value("unitName").String().Equal(data["unitName"].(string))
		product.Value("sort").Number().Equal(data["sort"].(int))
		product.Value("rank").Number().Equal(0)
		product.Value("sales").Number().Equal(0)
		product.Value("price").Number().Equal(0)
		product.Value("cost").Number().Equal(0)
		product.Value("otPrice").Number().Equal(0)
		product.Value("stock").Number().Equal(0)
		product.Value("isHot").Number().Equal(2)
		product.Value("isBenefit").Number().Equal(2)
		product.Value("isBest").Number().Equal(2)
		product.Value("isNew").Number().Equal(2)
		product.Value("isGood").Number().Equal(data["isGood"].(int))
		product.Value("productType").Number().Equal(1)
		product.Value("ficti").Number().Equal(0)
		product.Value("browse").Number().Equal(0)
		product.Value("codePath").String().Equal("")
		product.Value("videoLink").String().Equal(data["videoLink"].(string))
		product.Value("specType").Number().Equal(data["specType"].(int))
		product.Value("extensionType").Number().Equal(data["extensionType"].(int))
		product.Value("refusal").String().Equal("")
		product.Value("rate").Number().Equal(5)
		product.Value("replyCount").Number().Equal(0)
		product.Value("isGiftBag").Number().Equal(2)
		product.Value("careCount").Number().Equal(0)
		product.Value("image").String().NotEmpty()
		product.Value("sliderImage").String().NotEmpty()
		product.Value("oldId").Number().Equal(0)
		product.Value("tempId").Number().Equal(data["tempId"].(int))
		product.Value("sysBrandId").Number().Equal(data["sysBrandId"].(int))
		product.Value("productCategoryId").Number().Equal(data["cateId"].(int))

		obj = auth.POST("v1/merchant/product/changeProductIsShow").
			WithJSON(map[string]interface{}{"id": productId, "isShow": 1}).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("设置成功")

		obj = auth.DELETE(fmt.Sprintf("v1/merchant/product/deleteProduct/%d", int(productId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("操作成功")

		obj = auth.GET(fmt.Sprintf("v1/merchant/product/restoreProduct/%d", int(productId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("操作成功")

		// setUserAuthority
		obj = auth.DELETE(fmt.Sprintf("v1/merchant/product/destoryProduct/%d", int(productId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("删除成功")
	}

}
