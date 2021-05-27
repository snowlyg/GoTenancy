package router

import (
	"github.com/kataras/iris/v12"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitMediaRouter(Router iris.Party) {
	MediaGroup := Router.Party("/media")
	{
		MediaGroup.Post("/upload", v1.UploadFile)       // 上传文件
		MediaGroup.Post("/getFileList", v1.GetFileList) // 获取上传文件列表
		MediaGroup.Post("/deleteFile", v1.DeleteFile)   // 删除指定文件
		// MediaGroup.Post("/breakpointContinue", v1.BreakpointContinue)             // 断点续传
		// MediaGroup.Get("/findFile", v1.FindFile)                                  // 查询当前文件成功的切片
		// MediaGroup.Post("/breakpointContinueFinish", v1.BreakpointContinueFinish) // 查询当前文件成功的切片
		// MediaGroup.Post("/removeChunk", v1.RemoveChunk)                           // 查询当前文件成功的切片
	}
}
