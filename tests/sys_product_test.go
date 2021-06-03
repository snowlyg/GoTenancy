package tests

import (
	"net/http"
	"testing"
)

func TestProductList(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/product/getProductList").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("获取成功")

	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize")
	data.Value("pageSize").Number().Equal(10)
	data.Value("page").Number().Equal(1)
	data.Value("total").Number().Ge(0)

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
		"tenancyStatus",
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
		"giveCouponIds",
		"isGiftBag",
		"careCount",
		"image",
		"sliderImage",
		"oldId",
		"tempId",
		"sysTenancyId",
		"sysBrandId",
		"tenancyCategoryId",
		"createdAt",
		"updatedAt",
	)
	first.Value("id").Number().Ge(0)

}

func TestProductProcess(t *testing.T) {
	data := map[string]interface{}{
		"storeName":         "领立裁腰带短袖连衣裙",
		"storeInfo":         "短袖连衣裙",
		"keyword":           "短袖连衣裙",
		"barCode":           "",
		"isShow":            1,
		"status":            1,
		"tenancyStatus":     1,
		"unitName":          "件",
		"sort":              40,
		"rank":              0,
		"sales":             1,
		"price":             80,
		"cost":              50,
		"otPrice":           100,
		"stock":             399,
		"isHot":             0,
		"isBenefit":         0,
		"isBest":            0,
		"isNew":             0,
		"isGood":            1,
		"productType":       1,
		"ficti":             100,
		"browse":            0,
		"codePath":          "",
		"videoLink":         "",
		"specType":          1,
		"extensionType":     1,
		"refusal":           "",
		"rate":              5,
		"replyCount":        0,
		"giveCouponIds":     "",
		"isGiftBag":         2,
		"careCount":         0,
		"image":             "",
		"sliderImage":       "",
		"oldId":             0,
		"tempId":            0,
		"sysBrandId":        1,
		"tenancyCategoryId": 1,
	}
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/product/createProduct").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("创建成功")

	product := obj.Value("data").Object()
	product.Value("id").Number().Ge(0)
	product.Value("storeName").String().Equal(data["storeName"].(string))
	product.Value("storeInfo").String().Equal(data["storeInfo"].(string))
	product.Value("keyword").String().Equal(data["keyword"].(string))
	product.Value("barCode").String().Equal(data["barCode"].(string))
	product.Value("isShow").Number().Equal(data["isShow"].(int))
	product.Value("status").Number().Equal(data["status"].(int))
	product.Value("tenancyStatus").Number().Equal(data["tenancyStatus"].(int))
	product.Value("unitName").String().Equal(data["unitName"].(string))
	product.Value("sort").Number().Equal(data["sort"].(int))
	product.Value("rank").Number().Equal(data["rank"].(int))
	product.Value("sales").Number().Equal(data["sales"].(int))
	product.Value("price").Number().Equal(data["price"].(int))
	product.Value("cost").Number().Equal(data["cost"].(int))
	product.Value("otPrice").Number().Equal(data["otPrice"].(int))
	product.Value("stock").Number().Equal(data["stock"].(int))
	product.Value("isHot").Number().Equal(data["isHot"].(int))
	product.Value("isBenefit").Number().Equal(data["isBenefit"].(int))
	product.Value("isBest").Number().Equal(data["isBest"].(int))
	product.Value("isNew").Number().Equal(data["isNew"].(int))
	product.Value("isGood").Number().Equal(data["isGood"].(int))
	product.Value("productType").Number().Equal(data["productType"].(int))
	product.Value("ficti").Number().Equal(data["ficti"].(int))
	product.Value("browse").Number().Equal(data["browse"].(int))
	product.Value("codePath").String().Equal(data["codePath"].(string))
	product.Value("videoLink").String().Equal(data["videoLink"].(string))
	product.Value("specType").Number().Equal(data["specType"].(int))
	product.Value("extensionType").Number().Equal(data["extensionType"].(int))
	product.Value("refusal").String().Equal(data["refusal"].(string))
	product.Value("rate").Number().Equal(data["rate"].(int))
	product.Value("replyCount").Number().Equal(data["replyCount"].(int))
	product.Value("giveCouponIds").String().Equal(data["giveCouponIds"].(string))
	product.Value("isGiftBag").Number().Equal(data["isGiftBag"].(int))
	product.Value("careCount").Number().Equal(data["careCount"].(int))
	product.Value("image").String().Equal(data["image"].(string))
	product.Value("sliderImage").String().Equal(data["sliderImage"].(string))
	product.Value("oldId").Number().Equal(data["oldId"].(int))
	product.Value("tempId").Number().Equal(data["tempId"].(int))
	product.Value("sysBrandId").Number().Equal(data["sysBrandId"].(int))
	product.Value("tenancyCategoryId").Number().Equal(data["tenancyCategoryId"].(int))
	productId := product.Value("id").Number().Raw()

	update := map[string]interface{}{
		"id":                productId,
		"storeName":         "领立裁腰带短袖连衣裙",
		"storeInfo":         "短袖连衣裙",
		"keyword":           "短袖连衣裙",
		"barCode":           "",
		"isShow":            1,
		"status":            1,
		"tenancyStatus":     1,
		"unitName":          "件",
		"sort":              40,
		"rank":              0,
		"sales":             1,
		"price":             80,
		"cost":              50,
		"otPrice":           100,
		"stock":             399,
		"isHot":             0,
		"isBenefit":         0,
		"isBest":            0,
		"isNew":             0,
		"isGood":            1,
		"productType":       1,
		"ficti":             100,
		"browse":            0,
		"codePath":          "",
		"videoLink":         "",
		"specType":          1,
		"extensionType":     1,
		"refusal":           "",
		"rate":              5,
		"replyCount":        0,
		"giveCouponIds":     "",
		"isGiftBag":         2,
		"careCount":         0,
		"image":             "",
		"sliderImage":       "",
		"oldId":             0,
		"tempId":            0,
		"sysBrandId":        1,
		"tenancyCategoryId": 1,
	}

	obj = auth.PUT("/v1/admin/product/updateProduct").
		WithJSON(update).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("更新成功")
	product = obj.Value("data").Object()

	product.Value("id").Number().Ge(0)
	product.Value("storeName").String().Equal(update["storeName"].(string))
	product.Value("storeInfo").String().Equal(update["storeInfo"].(string))
	product.Value("keyword").String().Equal(update["keyword"].(string))
	product.Value("barCode").String().Equal(update["barCode"].(string))
	product.Value("isShow").Number().Equal(update["isShow"].(int))
	product.Value("status").Number().Equal(update["status"].(int))
	product.Value("tenancyStatus").Number().Equal(update["tenancyStatus"].(int))
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
	product.Value("giveCouponIds").String().Equal(update["giveCouponIds"].(string))
	product.Value("isGiftBag").Number().Equal(update["isGiftBag"].(int))
	product.Value("careCount").Number().Equal(update["careCount"].(int))
	product.Value("image").String().Equal(update["image"].(string))
	product.Value("sliderImage").String().Equal(update["sliderImage"].(string))
	product.Value("oldId").Number().Equal(update["oldId"].(int))
	product.Value("tempId").Number().Equal(update["tempId"].(int))
	product.Value("sysBrandId").Number().Equal(update["sysBrandId"].(int))
	product.Value("tenancyCategoryId").Number().Equal(update["tenancyCategoryId"].(int))

	obj = auth.POST("/v1/admin/product/getProductById").
		WithJSON(map[string]interface{}{"id": productId}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("操作成功")
	product = obj.Value("data").Object()

	product.Value("id").Number().Ge(0)
	product.Value("storeName").String().Equal(update["storeName"].(string))
	product.Value("storeInfo").String().Equal(update["storeInfo"].(string))
	product.Value("keyword").String().Equal(update["keyword"].(string))
	product.Value("barCode").String().Equal(update["barCode"].(string))
	product.Value("isShow").Number().Equal(update["isShow"].(int))
	product.Value("status").Number().Equal(update["status"].(int))
	product.Value("tenancyStatus").Number().Equal(update["tenancyStatus"].(int))
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
	product.Value("giveCouponIds").String().Equal(update["giveCouponIds"].(string))
	product.Value("isGiftBag").Number().Equal(update["isGiftBag"].(int))
	product.Value("careCount").Number().Equal(update["careCount"].(int))
	product.Value("image").String().Equal(update["image"].(string))
	product.Value("sliderImage").String().Equal(update["sliderImage"].(string))
	product.Value("oldId").Number().Equal(update["oldId"].(int))
	product.Value("tempId").Number().Equal(update["tempId"].(int))
	product.Value("sysBrandId").Number().Equal(update["sysBrandId"].(int))
	product.Value("tenancyCategoryId").Number().Equal(update["tenancyCategoryId"].(int))

	// setUserAuthority
	obj = auth.DELETE("/v1/admin/product/deleteProduct").
		WithJSON(map[string]interface{}{"id": productId}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("删除成功")

}

func TestProductAddError(t *testing.T) {
	data := map[string]interface{}{
		"storeName":         "",
		"storeInfo":         "短袖连衣裙",
		"keyword":           "短袖连衣裙",
		"barCode":           "",
		"isShow":            1,
		"status":            1,
		"tenancyStatus":     1,
		"unitName":          "件",
		"sort":              40,
		"rank":              0,
		"sales":             1,
		"price":             80,
		"cost":              50,
		"otPrice":           100,
		"stock":             399,
		"isHot":             0,
		"isBenefit":         0,
		"isBest":            0,
		"isNew":             0,
		"isGood":            1,
		"productType":       1,
		"ficti":             100,
		"browse":            0,
		"codePath":          "",
		"videoLink":         "",
		"specType":          1,
		"extensionType":     1,
		"refusal":           "",
		"rate":              5,
		"replyCount":        0,
		"giveCouponIds":     "",
		"isGiftBag":         2,
		"careCount":         0,
		"image":             "",
		"sliderImage":       "",
		"oldId":             0,
		"tempId":            0,
		"sysBrandId":        1,
		"tenancyCategoryId": 1,
	}
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/product/createProduct").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(4000)
	obj.Value("msg").String().Equal("Key: 'CreateTenancyProduct.StoreName' Error:Field validation for 'StoreName' failed on the 'required' tag")

}
