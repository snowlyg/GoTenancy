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
		"productCates",
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
		"isHot":     2,
		"isBenefit": 2,
		"isBest":    2,
		"isNew":     2,
		"content":   "dsfsafasfasfas",
	}

	obj := auth.PUT(fmt.Sprintf("v1/admin/product/updateProduct/%d", productId)).
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
	obj.Value("message").String().Equal("设置成功")

	obj = auth.POST("v1/admin/product/changeMutilProductStatus").
		WithJSON(map[string]interface{}{"id": []int{productId}, "status": 3}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("设置成功")

	obj = auth.GET(fmt.Sprintf("v1/admin/product/getProductById/%d", productId)).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
	product := obj.Value("data").Object()

	product.Value("id").Number().Ge(0)
	product.Value("storeName").String().Equal(update["storeName"].(string))
	product.Value("isHot").Number().Equal(update["isHot"].(int))
	product.Value("isBenefit").Number().Equal(update["isBenefit"].(int))
	product.Value("isBest").Number().Equal(update["isBest"].(int))
	product.Value("isNew").Number().Equal(update["isNew"].(int))
	product.Value("content").String().Equal(update["content"].(string))
}
