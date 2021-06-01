package tests

import (
	"net/http"
	"testing"
)

func TestParentRegion(t *testing.T) {
	e := baseTester(t)
	obj := e.GET("/v1/public/region/0").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("获取成功")
	obj.Value("data").Array().Length().Equal(31)
}

func TestSubRegion1(t *testing.T) {
	e := baseTester(t)
	obj := e.GET("/v1/public/region/19").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("获取成功")
	obj.Value("data").Array().Length().Equal(1)
}

func TestSubRegion2(t *testing.T) {
	e := baseTester(t)
	obj := e.GET("/v1/public/region/20").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("获取成功")
	obj.Value("data").Array().Length().Equal(16)
}

func TestSubRegionList(t *testing.T) {
	e := baseTester(t)
	obj := e.GET("/v1/public/getRegionList").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("获取成功")
	obj.Value("data").Array().Length().Equal(31)
}
