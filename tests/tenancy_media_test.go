package tests

import (
	"net/http"
	"os"
	"testing"
)

func TestMediaList(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/media/getFileList").
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
	data.Value("list").Array().Empty()

}

func TestMediaProcess(t *testing.T) {
	name := "test_img.jpg"
	path := "/api"
	fh, _ := os.Open("./" + name)
	defer fh.Close()
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("/v1/admin/media/upload").
		WithMultipart().
		WithFile("file", name, fh).
		WithForm(map[string]interface{}{"path": path}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("上传成功")

	media := obj.Value("data").Object().Value("file").Object()
	media.Value("id").Number().Ge(0)
	media.Value("name").String().Equal(name)
	media.Value("url").String().NotEmpty()
	media.Value("tag").String().Equal("jpg")
	media.Value("key").String().NotEmpty()
	mediaId := media.Value("id").Number().Raw()

	// setUserAuthority
	obj = auth.DELETE("/v1/admin/media/deleteFile").
		WithJSON(map[string]interface{}{"id": mediaId}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("删除成功")

}