package tests

import (
	"fmt"
	"net/http"
	"testing"
)

func TestProductList(t *testing.T) {
	params := []param{
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "1"}, length: 3},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "2"}, length: 1},
		// {args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "3"}, length: 1},
		// {args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "4"}, length: 1},
		// {args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "5"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "6"}, length: 1},
		{args: map[string]interface{}{"page": 1, "pageSize": 10, "type": "7"}, length: 1},
	}
	for _, param := range params {
		productlist(t, param.args, param.length)
	}
}

func productlist(t *testing.T, params map[string]interface{}, length int) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/product/getProductList").
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
		"sliderImage",
		"oldId",
		"tempId",
		"sysTenancyId",
		"sysTenancyName",
		"cateName",
		"brandName",
		"sysBrandId",
		"productCategoryId",
		"createdAt",
		"updatedAt",
	)
	first.Value("id").Number().Ge(0)

}

func TestGetProductFilter(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.GET("v1/admin/product/getProductFilter").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

}

func TestProductProcess(t *testing.T) {
	productId := 1
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)

	update := map[string]interface{}{
		"storeName": "领立裁腰带短袖连衣裙",
		"isHot":     0,
		"isBenefit": 0,
		"isBest":    0,
		"isNew":     0,
		"content":   "dsfsafasfasfas",
	}

	obj := auth.PUT(fmt.Sprintf("v1/admin/product/updateProduct/%d", int(productId))).
		WithJSON(update).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("更新成功")

	obj = auth.POST("v1/admin/product/changeProductStatus").
		WithJSON(map[string]interface{}{"id": productId, "status": 3}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")

	obj = auth.GET(fmt.Sprintf("v1/admin/product/getProductById/%d", int(productId))).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
	product := obj.Value("data").Object()

	product.Value("id").Number().Ge(0)
	product.Value("storeName").String().Equal(update["storeName"].(string))
	product.Value("storeInfo").String().Equal(update["storeInfo"].(string))
	product.Value("keyword").String().Equal(update["keyword"].(string))
	product.Value("barCode").String().Equal(update["barCode"].(string))
	product.Value("isShow").Number().Equal(update["isShow"].(int))
	product.Value("status").Number().Equal(update["status"].(int))
	product.Value("unitName").String().Equal(update["unitName"].(string))
	product.Value("sort").Number().Equal(update["sort"].(int))
	product.Value("rank").Number().Equal(update["rank"].(int))
	product.Value("sales").Number().Equal(update["sales"].(int))
	product.Value("price").Number().Equal(update["price"].(int))
	product.Value("cost").Number().Equal(update["cost"].(int))
	product.Value("otPrice").Number().Equal(update["otPrice"].(int))
	product.Value("stock").Number().Equal(update["stock"].(int))
	product.Value("isHot").Number().Equal(update["isHot"].(int))
	product.Value("isBenefit").Number().Equal(update["isBenefit"].(int))
	product.Value("isBest").Number().Equal(update["isBest"].(int))
	product.Value("isNew").Number().Equal(update["isNew"].(int))
	product.Value("isGood").Number().Equal(update["isGood"].(int))
	product.Value("productType").Number().Equal(update["productType"].(int))
	product.Value("ficti").Number().Equal(update["ficti"].(int))
	product.Value("browse").Number().Equal(update["browse"].(int))
	product.Value("codePath").String().Equal(update["codePath"].(string))
	product.Value("videoLink").String().Equal(update["videoLink"].(string))
	product.Value("specType").Number().Equal(update["specType"].(int))
	product.Value("extensionType").Number().Equal(update["extensionType"].(int))
	product.Value("refusal").String().Equal(update["refusal"].(string))
	product.Value("rate").Number().Equal(update["rate"].(int))
	product.Value("replyCount").Number().Equal(update["replyCount"].(int))
	product.Value("isGiftBag").Number().Equal(update["isGiftBag"].(int))
	product.Value("careCount").Number().Equal(update["careCount"].(int))
	product.Value("image").String().Equal(update["image"].(string))
	product.Value("sliderImage").String().Equal(update["sliderImage"].(string))
	product.Value("oldId").Number().Equal(update["oldId"].(int))
	product.Value("tempId").Number().Equal(update["tempId"].(int))
	product.Value("sysBrandId").Number().Equal(update["sysBrandId"].(int))
	product.Value("productCategoryId").Number().Equal(update["productCategoryId"].(int))

}
