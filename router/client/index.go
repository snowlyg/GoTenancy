package client

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

// 规格参数
func InitAttrTemplateRouter(Router *gin.RouterGroup) {
	AttrTemplateRouter := Router.Group("/attrTemplate")
	{
		AttrTemplateRouter.POST("/createAttrTemplate", v1.CreateAttrTemplate)
		AttrTemplateRouter.POST("/getAttrTemplateList", v1.GetAttrTemplateList)
		AttrTemplateRouter.POST("/getAttrTemplateById", v1.GetAttrTemplateById)
		AttrTemplateRouter.PUT("/updateAttrTemplate", v1.UpdateAttrTemplate)
		AttrTemplateRouter.DELETE("/deleteAttrTemplate", v1.DeleteAttrTemplate)
	}
}

// 系统配置
func InitConfigRouter(Router *gin.RouterGroup) {
	ConfigRouter := Router.Group("/config")
	{
		ConfigRouter.GET("/getConfigMap/:category", v1.GetConfigMap)
	}
}

// 系统配置数值
func InitConfigValueRouter(Router *gin.RouterGroup) {
	ConfigValueRouter := Router.Group("/configValue")
	{
		ConfigValueRouter.POST("/saveConfigValue/:category", v1.SaveConfigValue)
	}
}

// 商品分类
func InitCategoryRouter(Router *gin.RouterGroup) {
	CategoryRouter := Router.Group("/category")
	{
		CategoryRouter.GET("/getCreateTenancyCategoryMap", v1.GetCreateTenancyCategoryMap)
		CategoryRouter.GET("/getUpdateTenancyCategoryMap/:id", v1.GetUpdateTenancyCategoryMap)
		CategoryRouter.GET("/getCategorySelect", v1.GetClientCategorySelect)
		CategoryRouter.GET("/getAdminCategorySelect", v1.GetCategorySelect)
		CategoryRouter.POST("/createCategory", v1.CreateCategory)
		CategoryRouter.POST("/getCategoryList", v1.GetClientCategoryList)
		CategoryRouter.GET("/getCategoryById/:id", v1.GetCategoryById)
		CategoryRouter.POST("/changeCategoryStatus", v1.ChangeCategoryStatus)
		CategoryRouter.PUT("/updateCategory/:id", v1.UpdateCategory)
		CategoryRouter.DELETE("/deleteCategory/:id", v1.DeleteCategory)
	}
}

// 多媒体
func InitMediaRouter(Router *gin.RouterGroup) {
	MediaGroup := Router.Group("/media")
	{
		MediaGroup.GET("/getUpdateMediaMap/:id", v1.GetUpdateMediaMap) // 修改名称表单
		MediaGroup.POST("/upload", v1.UploadFile)                      // 上传文件
		MediaGroup.POST("/getFileList", v1.GetFileList)                // 获取上传文件列表
		MediaGroup.POST("/updateMediaName/:id", v1.UpdateMediaName)    // 修改名称
		MediaGroup.DELETE("/deleteFile", v1.DeleteFile)                // 删除指定文件
	}
}

// 商品
func InitProductRouter(Router *gin.RouterGroup) {
	ProductRouter := Router.Group("/product")
	{
		ProductRouter.GET("/getEditProductFictiMap/:id", v1.GetEditProductFictiMap)
		ProductRouter.GET("/getProductFilter", v1.GetProductFilter)
		ProductRouter.PUT("/setProductFicti/:id", v1.SetProductFicti)
		ProductRouter.POST("/createProduct", v1.CreateProduct)
		ProductRouter.POST("/changeProductStatus", v1.ChangeProductStatus)
		ProductRouter.POST("/getProductList", v1.GetProductList)
		ProductRouter.GET("/getProductById/:id", v1.GetProductById)
		ProductRouter.PUT("/updateProduct/:id", v1.UpdateProduct)
		ProductRouter.DELETE("/deleteProduct/:id", v1.DeleteProduct)
	}
}

// 商户
func InitTenancyRouter(Router *gin.RouterGroup) {
	TenancyRouter := Router.Group("/tenancy")
	{
		TenancyRouter.GET("/getTenancyInfo", v1.GetTenancyInfo)           // 登录商户信息
		TenancyRouter.GET("/getUpdateTenancyMap", v1.GetUpdateTenancyMap) // 获取商户编辑表单
		TenancyRouter.PUT("/updateTenancy/:id", v1.UpdateClientTenancy)   // 获取商户编辑表单
	}
}

func InitMenuRouter(Router *gin.RouterGroup) (R *gin.RouterGroup) {
	MenuRouter := Router.Group("/menu")
	{
		MenuRouter.GET("/getMenu", v1.GetMenu) // 获取菜单树
	}
	return MenuRouter
}

func InitSysOperationRecordRouter(Router *gin.RouterGroup) {
	SysOperationRecordRouter := Router.Group("/sysOperationRecord")
	{
		SysOperationRecordRouter.POST("/getSysOperationRecordList", v1.GetSysOperationRecordList) // 获取SysOperationRecord列表

	}
}
