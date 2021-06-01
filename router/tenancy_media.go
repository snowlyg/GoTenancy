package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitMediaRouter(Router *gin.RouterGroup) {
	MediaGroup := Router.Group("/media")
	{
		MediaGroup.POST("/upload", v1.UploadFile)       // 上传文件
		MediaGroup.POST("/getFileList", v1.GetFileList) // 获取上传文件列表
		MediaGroup.POST("/deleteFile", v1.DeleteFile)   // 删除指定文件
		// MediaGroup.POST("/breakpointContinue", v1.BreakpointContinue)             // 断点续传
		// MediaGroup.GET("/findFile", v1.FindFile)                                  // 查询当前文件成功的切片
		// MediaGroup.POST("/breakpointContinueFinish", v1.BreakpointContinueFinish) // 查询当前文件成功的切片
		// MediaGroup.POST("/removeChunk", v1.RemoveChunk)                           // 查询当前文件成功的切片
	}
}
