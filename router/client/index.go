package client

import (
	"github.com/gin-gonic/gin"
	client "github.com/snowlyg/go-tenancy/api/v1/client"
)

// 规格参数
func InitAttrTemplateRouter(Router *gin.RouterGroup) {
	AttrTemplateRouter := Router.Group("/attrTemplate")
	{
		AttrTemplateRouter.POST("/getAttrTemplateList", client.GetAttrTemplateList)
		AttrTemplateRouter.POST("/createAttrTemplate", client.CreateAttrTemplate)
		AttrTemplateRouter.GET("/getAttrTemplateById/:id", client.GetAttrTemplateById)
		AttrTemplateRouter.PUT("/updateAttrTemplate/:id", client.UpdateAttrTemplate)
		AttrTemplateRouter.DELETE("/deleteAttrTemplate/:id", client.DeleteAttrTemplate)
	}
}

func InitBrandRouter(Router *gin.RouterGroup) {
	BrandRouter := Router.Group("/brand")
	{
		BrandRouter.GET("/getBrandList", client.GetBrandList)
	}
}

// 系统配置
func InitConfigRouter(Router *gin.RouterGroup) {
	ConfigRouter := Router.Group("/config")
	{
		ConfigRouter.GET("/getConfigMap/:category", client.GetConfigMap)
	}
}

// 系统配置数值
func InitConfigValueRouter(Router *gin.RouterGroup) {
	ConfigValueRouter := Router.Group("/configValue")
	{
		ConfigValueRouter.POST("/saveConfigValue/:category", client.SaveConfigValue)
	}
}

// 商品分类
func InitCategoryRouter(Router *gin.RouterGroup) {
	CategoryRouter := Router.Group("/productCategory")
	{
		CategoryRouter.GET("/getCreateProductCategoryMap", client.GetCreateProductCategoryMap)
		CategoryRouter.GET("/getUpdateProductCategoryMap/:id", client.GetUpdateProductCategoryMap)
		CategoryRouter.GET("/getProductCategorySelect", client.GetProductCategorySelect)
		CategoryRouter.GET("/getAdminProductCategorySelect", client.GetAdminProductCategorySelect)
		CategoryRouter.POST("/createProductCategory", client.CreateProductCategory)
		CategoryRouter.POST("/getProductCategoryList", client.GetProductCategoryList)
		CategoryRouter.GET("/getProductCategoryById/:id", client.GetProductCategoryById)
		CategoryRouter.POST("/changeProductCategoryStatus", client.ChangeProductCategoryStatus)
		CategoryRouter.PUT("/updateProductCategory/:id", client.UpdateProductCategory)
		CategoryRouter.DELETE("/deleteProductCategory/:id", client.DeleteProductCategory)
	}
}

// 多媒体
func InitMediaRouter(Router *gin.RouterGroup) {
	MediaGroup := Router.Group("/media")
	{
		MediaGroup.GET("/getUpdateMediaMap/:id", client.GetUpdateMediaMap) // 修改名称表单
		MediaGroup.POST("/upload", client.UploadFile)                      // 上传文件
		MediaGroup.POST("/getFileList", client.GetFileList)                // 获取上传文件列表
		MediaGroup.POST("/updateMediaName/:id", client.UpdateMediaName)    // 修改名称
		MediaGroup.DELETE("/deleteFile", client.DeleteFile)                // 删除指定文件
	}
}

// 商品
func InitProductRouter(Router *gin.RouterGroup) {
	ProductRouter := Router.Group("/product")
	{
		ProductRouter.GET("/getProductFilter", client.GetProductFilter)
		ProductRouter.POST("/createProduct", client.CreateProduct)
		ProductRouter.POST("/changeProductIsShow", client.ChangeProductIsShow)
		ProductRouter.POST("/getProductList", client.GetProductList)
		ProductRouter.GET("/getProductById/:id", client.GetProductById)
		ProductRouter.PUT("/updateProduct/:id", client.UpdateProduct)
		ProductRouter.GET("/restoreProduct/:id", client.RestoreProduct)
		ProductRouter.DELETE("/deleteProduct/:id", client.DeleteProduct)
		ProductRouter.DELETE("/destoryProduct/:id", client.DestoryProduct)
	}
}

// 商户
func InitTenancyRouter(Router *gin.RouterGroup) {
	TenancyRouter := Router.Group("/tenancy")
	{
		TenancyRouter.GET("/getTenancyInfo", client.GetTenancyInfo)           // 登录商户信息
		TenancyRouter.GET("/getUpdateTenancyMap", client.GetUpdateTenancyMap) // 获取商户编辑表单
		TenancyRouter.PUT("/updateTenancy/:id", client.UpdateClientTenancy)   // 获取商户编辑表单
		TenancyRouter.GET("/getTenancyCopyCount", client.GetTenancyCopyCount) // 获取商户复制商品次数
	}
}

func InitMenuRouter(Router *gin.RouterGroup) (R *gin.RouterGroup) {
	MenuRouter := Router.Group("/menu")
	{
		MenuRouter.GET("/getMenu", client.GetMenu) // 获取菜单树
	}
	return MenuRouter
}

// 运费模板
func InitShippingTemplateRouter(Router *gin.RouterGroup) {
	ShippingTemplateRouter := Router.Group("/shippingTemplate")
	{
		ShippingTemplateRouter.GET("/getShippingTemplateSelect", client.GetShippingTemplateSelect)
		ShippingTemplateRouter.POST("/getShippingTemplateList", client.GetShippingTemplateList)
		ShippingTemplateRouter.POST("/createShippingTemplate", client.CreateShippingTemplate)
		ShippingTemplateRouter.GET("/getShippingTemplateById/:id", client.GetShippingTemplateById)
		ShippingTemplateRouter.PUT("/updateShippingTemplate/:id", client.UpdateShippingTemplate)
		ShippingTemplateRouter.DELETE("/deleteShippingTemplate/:id", client.DeleteShippingTemplate)
	}
}

func InitSysOperationRecordRouter(Router *gin.RouterGroup) {
	SysOperationRecordRouter := Router.Group("/sysOperationRecord")
	{
		SysOperationRecordRouter.POST("/getSysOperationRecordList", client.GetSysOperationRecordList) // 获取SysOperationRecord列表
	}
}

func InitOrderRouter(Router *gin.RouterGroup) {
	OrderRouter := Router.Group("/order")
	{
		OrderRouter.GET("/getOrderRemarkMap/:id", client.GetOrderRemarkMap)
		OrderRouter.POST("/getOrderList", client.GetOrderList)
		OrderRouter.GET("/getOrderChart", client.GetOrderChart)
		OrderRouter.GET("/getOrderFilter", client.GetOrderFilter)
		OrderRouter.GET("/getOrderById/:id", client.GetOrderById)
		OrderRouter.POST("/getOrderRecord/:id", client.GetOrderRecord)
		OrderRouter.POST("/remarkOrder/:id", client.RemarkOrder)
	}
}
