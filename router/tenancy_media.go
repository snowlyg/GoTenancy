package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitMediaRouter(Router *gin.RouterGroup) {
	MediaGroup := Router.Group("/media")
	{
		MediaGroup.GET("/getUpdateMediaMap/:id", v1.GetUpdateMediaMap) // 修改名称表单
		MediaGroup.POST("/upload", v1.UploadFile)                      // 上传文件
		MediaGroup.POST("/getFileList", v1.GetFileList)                // 获取上传文件列表
		MediaGroup.POST("/updateMediaName", v1.UpdateMediaName)        // 修改名称
		MediaGroup.DELETE("/deleteFile", v1.DeleteFile)                // 删除指定文件
	}
}
