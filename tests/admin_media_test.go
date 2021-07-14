package tests

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestMediaList(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/media/getFileList").
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
}

func TestMediaProcess(t *testing.T) {
	name := "test_img.jpg"
	path := "/api"
	fh, _ := os.Open("./" + name)
	defer fh.Close()
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/media/upload").
		WithMultipart().
		WithFile("file", name, fh).
		WithForm(map[string]interface{}{"path": path}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("上传成功")

	obj.Value("data").Object().Value("src").String().NotEmpty()
	mediaId := obj.Value("data").Object().Value("id").Number().Raw()
	if mediaId > 0 {
		// getUpdateMediaMap
		obj = auth.GET(fmt.Sprintf("v1/admin/media/getUpdateMediaMap/%d", int(mediaId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		// changeTenancyStatus
		obj = auth.POST(fmt.Sprintf("v1/admin/media/updateMediaName/%d", int(mediaId))).
			WithJSON(map[string]interface{}{
				"id":   mediaId,
				"name": "name_jpg",
			}).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("修改成功")

		// deleteFile
		obj = auth.DELETE("v1/admin/media/deleteFile").
			WithJSON(map[string]interface{}{"ids": []float64{mediaId}}).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("删除成功")

	}

}
