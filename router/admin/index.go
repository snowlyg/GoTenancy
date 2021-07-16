package admin

import (
	"github.com/gin-gonic/gin"
	admin "github.com/snowlyg/go-tenancy/api/v1/admin"
)

// api
func InitApiRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("/api")
	{
		ApiRouter.POST("/createApi", admin.CreateApi)               // 创建Api
		ApiRouter.DELETE("/deleteApi", admin.DeleteApi)             // 删除Api
		ApiRouter.POST("/getApiList", admin.GetApiList)             // 获取Api列表
		ApiRouter.POST("/getApiById", admin.GetApiById)             // 获取单条Api消息
		ApiRouter.POST("/updateApi", admin.UpdateApi)               // 更新api
		ApiRouter.POST("/getAllApis", admin.GetAllApis)             // 获取所有api
		ApiRouter.DELETE("/deleteApisByIds", admin.DeleteApisByIds) // 删除选中api
	}
}

// 角色管理
func InitAuthorityRouter(Router *gin.RouterGroup) {
	AuthorityRouter := Router.Group("/authority")
	{
		AuthorityRouter.POST("/createAuthority", admin.CreateAuthority)                 // 创建角色
		AuthorityRouter.DELETE("/deleteAuthority", admin.DeleteAuthority)               // 删除角色
		AuthorityRouter.PUT("/updateAuthority", admin.UpdateAuthority)                  // 更新角色
		AuthorityRouter.POST("/copyAuthority", admin.CopyAuthority)                     // 更新角色
		AuthorityRouter.POST("/getAuthorityList", admin.GetAuthorityList)               // 获取角色列表
		AuthorityRouter.POST("/getAdminAuthorityList", admin.GetAdminAuthorityList)     // 获取员工角色列表
		AuthorityRouter.POST("/getTenancyAuthorityList", admin.GetTenancyAuthorityList) // 获取商户角色列表
		AuthorityRouter.POST("/getGeneralAuthorityList", admin.GetGeneralAuthorityList) // 获取普通用户角色列表
		AuthorityRouter.POST("/setDataAuthority", admin.SetDataAuthority)               // 设置角色资源权限
	}
}

// 品牌分类
func InitBrandCategoryRouter(Router *gin.RouterGroup) {
	BrandCategoryRouter := Router.Group("/brandCategory")
	{
		BrandCategoryRouter.GET("/getBrandCategoryList", admin.GetBrandCategoryList)
		BrandCategoryRouter.GET("/getCreateBrandCategoryMap", admin.GetCreateBrandCategoryMap)
		BrandCategoryRouter.GET("/getUpdateBrandCategoryMap/:id", admin.GetUpdateBrandCategoryMap)
		BrandCategoryRouter.POST("/createBrandCategory", admin.CreateBrandCategory)
		BrandCategoryRouter.GET("/getBrandCategoryById/:id", admin.GetBrandCategoryById)
		BrandCategoryRouter.POST("/changeBrandCategoryStatus", admin.ChangeBrandCategoryStatus)
		BrandCategoryRouter.PUT("/updateBrandCategory/:id", admin.UpdateBrandCategory)
		BrandCategoryRouter.DELETE("/deleteBrandCategory/:id", admin.DeleteBrandCategory)
	}
}

// 品牌
func InitBrandRouter(Router *gin.RouterGroup) {
	BrandRouter := Router.Group("/brand")
	{
		BrandRouter.POST("/getBrandList", admin.GetBrandList)
		BrandRouter.GET("/getCreateBrandMap", admin.GetCreateBrandMap)
		BrandRouter.GET("/getUpdateBrandMap/:id", admin.GetUpdateBrandMap)
		BrandRouter.POST("/createBrand", admin.CreateBrand)
		BrandRouter.GET("/getBrandById/:id", admin.GetBrandById)
		BrandRouter.POST("/changeBrandStatus", admin.ChangeBrandStatus)
		BrandRouter.PUT("/updateBrand/:id", admin.UpdateBrand)
		BrandRouter.DELETE("/deleteBrand/:id", admin.DeleteBrand)
	}
}

// casbin
func InitCasbinRouter(Router *gin.RouterGroup) {
	CasbinRouter := Router.Group("/casbin")
	{
		CasbinRouter.POST("/updateCasbin", admin.UpdateCasbin)
		CasbinRouter.POST("/getPolicyPathByAuthorityId", admin.GetPolicyPathByAuthorityId)
	}
}

// 分类
func InitCategoryRouter(Router *gin.RouterGroup) {
	CategoryRouter := Router.Group("/productCategory")
	{
		CategoryRouter.GET("/getCreateProductCategoryMap", admin.GetCreateProductCategoryMap)
		CategoryRouter.GET("/getUpdateProductCategoryMap/:id", admin.GetUpdateProductCategoryMap)
		CategoryRouter.GET("/getProductCategorySelect", admin.GetProductCategorySelect)
		CategoryRouter.POST("/createProductCategory", admin.CreateProductCategory)
		CategoryRouter.POST("/getProductCategoryList", admin.GetProductCategoryList)
		CategoryRouter.GET("/getProductCategoryById/:id", admin.GetProductCategoryById)
		CategoryRouter.POST("/changeProductCategoryStatus", admin.ChangeProductCategoryStatus)
		CategoryRouter.PUT("/updateProductCategory/:id", admin.UpdateProductCategory)
		CategoryRouter.DELETE("/deleteProductCategory/:id", admin.DeleteProductCategory)
	}
}

// 配置分类
func InitConfigCategoryRouter(Router *gin.RouterGroup) {
	ConfigCategoryRouter := Router.Group("/configCategory")
	{
		ConfigCategoryRouter.GET("/getCreateConfigCategoryMap", admin.GetCreateConfigCategoryMap)
		ConfigCategoryRouter.GET("/getUpdateConfigCategoryMap/:id", admin.GetUpdateConfigCategoryMap)
		ConfigCategoryRouter.GET("/getConfigCategoryList", admin.GetConfigCategoryList)
		ConfigCategoryRouter.POST("/createConfigCategory", admin.CreateConfigCategory)
		ConfigCategoryRouter.GET("/getConfigCategoryById/:id", admin.GetConfigCategoryById)
		ConfigCategoryRouter.PUT("/updateConfigCategory/:id", admin.UpdateConfigCategory)
		ConfigCategoryRouter.POST("/changeConfigCategoryStatus", admin.ChangeConfigCategoryStatus)
		ConfigCategoryRouter.DELETE("/deleteConfigCategory/:id", admin.DeleteConfigCategory)
	}
}

// 配置值
func InitConfigValueRouter(Router *gin.RouterGroup) {
	ConfigValueRouter := Router.Group("/configValue")
	{
		ConfigValueRouter.POST("/saveConfigValue/:category", admin.SaveConfigValue)
	}
}

// 系统配置
func InitConfigRouter(Router *gin.RouterGroup) {
	ConfigRouter := Router.Group("/config")
	{
		ConfigRouter.GET("/getConfigMap/:category", admin.GetConfigMap)
		ConfigRouter.GET("/getCreateConfigMap", admin.GetCreateConfigMap)
		ConfigRouter.GET("/getUpdateConfigMap/:id", admin.GetUpdateConfigMap)
		ConfigRouter.POST("/getConfigList", admin.GetConfigList)
		ConfigRouter.POST("/createConfig", admin.CreateConfig)
		ConfigRouter.GET("/getConfigByKey/:key", admin.GetConfigByKey)
		ConfigRouter.GET("/getConfigByID/:id", admin.GetConfigByID)
		ConfigRouter.POST("/changeConfigStatus", admin.ChangeConfigStatus)
		ConfigRouter.PUT("/updateConfig/:id", admin.UpdateConfig)
		ConfigRouter.DELETE("/deleteConfig/:id", admin.DeleteConfig)
	}
}

// 邮件测试
func InitEmailRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("/email")
	{
		UserRouter.POST("/emailTest", admin.EmailTest) // 发送测试邮件
	}
}

// 多媒体
func InitMediaRouter(Router *gin.RouterGroup) {
	MediaGroup := Router.Group("/media")
	{
		MediaGroup.GET("/getUpdateMediaMap/:id", admin.GetUpdateMediaMap) // 修改名称表单
		MediaGroup.POST("/upload", admin.UploadFile)                      // 上传文件
		MediaGroup.POST("/getFileList", admin.GetFileList)                // 获取上传文件列表
		MediaGroup.POST("/updateMediaName/:id", admin.UpdateMediaName)    // 修改名称
		MediaGroup.DELETE("/deleteFile", admin.DeleteFile)                // 删除指定文件
	}
}

// 菜单管理
func InitMenuRouter(Router *gin.RouterGroup) {
	MenuRouter := Router.Group("/menu")
	{
		MenuRouter.GET("/getMenu", admin.GetMenu)                           // 获取菜单树
		MenuRouter.GET("/getAddMenuMap", admin.GetAddMenuMap)               // 添加表单
		MenuRouter.GET("/getAddTenancyMenuMap", admin.GetAddTenancyMenuMap) // 添加商户菜单表单
		MenuRouter.GET("/getEditMenuMap/:id", admin.GetEditMenuMap)         // 编辑表单
		MenuRouter.GET("/getMenuList", admin.GetMenuList)                   // 分页获取基础menu列表
		MenuRouter.POST("/addBaseMenu", admin.AddBaseMenu)                  // 新增菜单
		MenuRouter.POST("/addTenancyBaseMenu", admin.AddTenancyBaseMenu)    // 新增商户菜单
		MenuRouter.POST("/getBaseMenuTree", admin.GetBaseMenuTree)          // 获取用户动态路由
		MenuRouter.POST("/addMenuAuthority", admin.AddMenuAuthority)        // 增加menu和角色关联关系
		MenuRouter.POST("/getMenuAuthority", admin.GetMenuAuthority)        // 获取指定角色menu
		MenuRouter.DELETE("/deleteBaseMenu/:id", admin.DeleteBaseMenu)      // 删除菜单
		MenuRouter.POST("/updateBaseMenu/:id", admin.UpdateBaseMenu)        // 更新菜单
		MenuRouter.POST("/getBaseMenuById/:id", admin.GetBaseMenuById)      // 根据id获取菜单
		ClientMenuRouter := MenuRouter.Group("/merchant")
		{
			ClientMenuRouter.GET("/getClientMenuList", admin.GetClientMenuList) // 分页获取基础menu列表
		}
	}
}

// 小程序管理
func InitMiniRouter(Router *gin.RouterGroup) {
	MiniRouter := Router.Group("/mini")
	{
		MiniRouter.POST("/createMini", admin.CreateMini)
		MiniRouter.POST("/getMiniList", admin.GetMiniList)
		MiniRouter.POST("/getMiniById", admin.GetMiniById)
		MiniRouter.PUT("/updateMini", admin.UpdateMini)
		MiniRouter.DELETE("/deleteMini", admin.DeleteMini)
	}
}

// 操作日志
func InitSysOperationRecordRouter(Router *gin.RouterGroup) {
	SysOperationRecordRouter := Router.Group("/sysOperationRecord")
	{
		SysOperationRecordRouter.POST("/getSysOperationRecordList", admin.GetSysOperationRecordList)           // 获取SysOperationRecord列表
		SysOperationRecordRouter.POST("/createSysOperationRecord", admin.CreateSysOperationRecord)             // 新建SysOperationRecord
		SysOperationRecordRouter.DELETE("/deleteSysOperationRecord", admin.DeleteSysOperationRecord)           // 删除SysOperationRecord
		SysOperationRecordRouter.DELETE("/deleteSysOperationRecordByIds", admin.DeleteSysOperationRecordByIds) // 批量删除SysOperationRecord
		SysOperationRecordRouter.GET("/findSysOperationRecord", admin.FindSysOperationRecord)                  // 根据ID获取SysOperationRecord

	}
}

// 产品管理
func InitProductRouter(Router *gin.RouterGroup) {
	ProductRouter := Router.Group("/product")
	{
		ProductRouter.GET("/getEditProductFictiMap/:id", admin.GetEditProductFictiMap)
		ProductRouter.GET("/getProductFilter", admin.GetProductFilter)
		ProductRouter.PUT("/setProductFicti/:id", admin.SetProductFicti)                //设置虚拟销量
		ProductRouter.POST("/changeProductStatus", admin.ChangeProductStatus)           // 强制下架，重新审核
		ProductRouter.POST("/changeMutilProductStatus", admin.ChangeMutilProductStatus) // 强制下架，重新审核
		ProductRouter.POST("/getProductList", admin.GetProductList)
		ProductRouter.GET("/getProductById/:id", admin.GetProductById)
		ProductRouter.PUT("/updateProduct/:id", admin.UpdateProduct)
	}
}

// 系统设置
func InitSystemRouter(Router *gin.RouterGroup) {
	SystemRouter := Router.Group("/system")
	{
		SystemRouter.POST("/getSystemConfig", admin.GetSystemConfig) // 获取配置文件内容
		SystemRouter.POST("/setSystemConfig", admin.SetSystemConfig) // 设置配置文件内容
		SystemRouter.POST("/getServerInfo", admin.GetServerInfo)     // 获取服务器信息
		SystemRouter.POST("/reloadSystem", admin.ReloadSystem)       // 重启服务
	}
}

// 商户
func InitTenancyRouter(Router *gin.RouterGroup) {
	TenancyRouter := Router.Group("/tenancy")
	{
		TenancyRouter.GET("/getTenancySelect", admin.GetTenancySelect)        // 获取Tenancy列表(不分页)
		TenancyRouter.GET("/changeCopyMap/:id", admin.ChangeCopyMap)          // 获取修改商品复制次数map
		TenancyRouter.GET("/getTenancies/:code", admin.GetTenanciesByRegion)  // 获取Tenancy列表(不分页)
		TenancyRouter.GET("/getTenancyCount", admin.GetTenancyCount)          // 获取Tenancy对应状态数量
		TenancyRouter.POST("/createTenancy", admin.CreateTenancy)             // 创建Tenancy
		TenancyRouter.POST("/setCopyProductNum/:id", admin.SetCopyProductNum) // 设置商品复制次数
		TenancyRouter.POST("/loginTenancy/:id", admin.LoginTenancy)           // 登录商户
		TenancyRouter.POST("/getTenancyList", admin.GetTenanciesList)         // 获取Tenancy列表
		TenancyRouter.GET("/getTenancyById/:id", admin.GetTenancyById)        // 获取单条Tenancy消息
		TenancyRouter.POST("/setTenancyRegion", admin.SetTenancyRegion)       // 设置商户地区
		TenancyRouter.POST("/changeTenancyStatus", admin.ChangeTenancyStatus) // 设置商户显示/隐藏
		TenancyRouter.PUT("/updateTenancy/:id", admin.UpdateTenancy)          // 更新Tenancy
		TenancyRouter.DELETE("/deleteTenancy/:id", admin.DeleteTenancy)       // 删除Tenancy
	}
}

// 管理员
func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("/user")
	{
		UserRouter.POST("/registerAdmin", admin.RegisterAdmin)       // 注册管理员
		UserRouter.POST("/registerTenancy", admin.RegisterTenancy)   // 注册商户管理员
		UserRouter.POST("/changePassword", admin.ChangePassword)     // 修改密码
		UserRouter.POST("/changeProfile", admin.ChangeProfile)       // 修改密码
		UserRouter.POST("/getAdminList", admin.GetAdminList)         // 分页获取管理员列表
		UserRouter.POST("/getTenancyList", admin.GetTenancyList)     // 分页获取商户管理员列表
		UserRouter.POST("/setUserAuthority", admin.SetUserAuthority) // 设置用户权限
		UserRouter.DELETE("/deleteUser", admin.DeleteUser)           // 删除用户
		UserRouter.PUT("/setUserInfo/:user_id", admin.SetUserInfo)   // 设置用户信息

	}
}

// 管理员
func InitCUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("/cuser")
	{
		UserRouter.GET("/editUserMap/:id", admin.UpdateUserMap)              // 批量设置分组表单
		UserRouter.POST("/editUser/:id", admin.UpdateUser)                   // 设置分组
		UserRouter.POST("/batchSetUserGroupMap", admin.BatchSetUserGroupMap) // 批量设置分组表单
		UserRouter.POST("/batchSetUserGroup", admin.BatchSetUserGroup)       // 设置分组
		UserRouter.POST("/batchSetUserLabelMap", admin.BatchSetUserLabelMap) // 批量设置标签表单
		UserRouter.POST("/batchSetUserLabel", admin.BatchSetUserLabel)       // 设置标签
		UserRouter.GET("/setUserGroupMap/:id", admin.SetUserGroupMap)        // 设置分组表单
		UserRouter.POST("/setUserGroup/:id", admin.SetUserGroup)             // 设置分组
		UserRouter.GET("/setUserLabelMap/:id", admin.SetUserLabelMap)        // 设置标签表单
		UserRouter.POST("/setUserLabel/:id", admin.SetUserLabel)             // 设置标签
		UserRouter.GET("/setNowMoneyMap/:id", admin.SetNowMoneyMap)          // 设置余额表单
		UserRouter.POST("/setNowMoney/:id", admin.SetNowMoney)               // 设置余额
		UserRouter.GET("/getGeneralDetail/:id", admin.GetGeneralDetail)      // 用户消费详情
		UserRouter.POST("/getGeneralList", admin.GetGeneralList)             // 分页获取c用户列表
		UserRouter.POST("/getOrderList/:id", admin.GetUserOrderList)         // 用户订单列表
		UserRouter.POST("/getBillList/:id", admin.GetBillList)               // 订单列表
	}
}

// 订单管理
func InitOrderRouter(Router *gin.RouterGroup) {
	OrderRouter := Router.Group("/order")
	{
		OrderRouter.POST("/getOrderList", admin.GetOrderList)
		OrderRouter.GET("/getOrderChart", admin.GetOrderChart)
		OrderRouter.GET("/getOrderById/:id", admin.GetOrderById)
	}
}

// 物流信息
func InitExpressRouter(Router *gin.RouterGroup) {
	ExpressRouter := Router.Group("/express")
	{
		ExpressRouter.POST("/getExpressList", admin.GetExpressList)
		ExpressRouter.GET("/getCreateExpressMap", admin.GetCreateExpressMap)
		ExpressRouter.GET("/getUpdateExpressMap/:id", admin.GetUpdateExpressMap)
		ExpressRouter.POST("/createExpress", admin.CreateExpress)
		ExpressRouter.GET("/getExpressById/:id", admin.GetExpressById)
		ExpressRouter.POST("/changeExpressStatus", admin.ChangeExpressStatus)
		ExpressRouter.PUT("/updateExpress/:id", admin.UpdateExpress)
		ExpressRouter.DELETE("/deleteExpress/:id", admin.DeleteExpress)
	}
}

// 退款单
func InitRefundOrderRouter(Router *gin.RouterGroup) {
	RefundOrderRouter := Router.Group("/refundOrder")
	{
		RefundOrderRouter.POST("/getRefundOrderList", admin.GetRefundOrderList)
	}
}

// 用户分组
func InitUserGroupRouter(Router *gin.RouterGroup) {
	UserGroupRouter := Router.Group("/userGroup")
	{
		UserGroupRouter.POST("/getUserGroupList", admin.GetUserGroupList)
		UserGroupRouter.GET("/getCreateUserGroupMap", admin.GetCreateUserGroupMap)
		UserGroupRouter.GET("/getUpdateUserGroupMap/:id", admin.GetUpdateUserGroupMap)
		UserGroupRouter.POST("/createUserGroup", admin.CreateUserGroup)
		UserGroupRouter.PUT("/updateUserGroup/:id", admin.UpdateUserGroup)
		UserGroupRouter.DELETE("/deleteUserGroup/:id", admin.DeleteUserGroup)
	}
}

// 用户标签
func InitUserLabelRouter(Router *gin.RouterGroup) {
	UserLabelRouter := Router.Group("/userLabel")
	{
		UserLabelRouter.POST("/getUserLabelList", admin.GetUserLabelList)
		UserLabelRouter.GET("/getCreateUserLabelMap", admin.GetCreateUserLabelMap)
		UserLabelRouter.GET("/getUpdateUserLabelMap/:id", admin.GetUpdateUserLabelMap)
		UserLabelRouter.POST("/createUserLabel", admin.CreateUserLabel)
		UserLabelRouter.PUT("/updateUserLabel/:id", admin.UpdateUserLabel)
		UserLabelRouter.DELETE("/deleteUserLabel/:id", admin.DeleteUserLabel)
	}
}
