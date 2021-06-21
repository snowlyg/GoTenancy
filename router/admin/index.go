package admin

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitApiRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("/api")
	{
		ApiRouter.POST("/createApi", v1.CreateApi)               // 创建Api
		ApiRouter.DELETE("/deleteApi", v1.DeleteApi)             // 删除Api
		ApiRouter.POST("/getApiList", v1.GetApiList)             // 获取Api列表
		ApiRouter.POST("/getApiById", v1.GetApiById)             // 获取单条Api消息
		ApiRouter.POST("/updateApi", v1.UpdateApi)               // 更新api
		ApiRouter.POST("/getAllApis", v1.GetAllApis)             // 获取所有api
		ApiRouter.DELETE("/deleteApisByIds", v1.DeleteApisByIds) // 删除选中api
	}
}

func InitAuthorityRouter(Router *gin.RouterGroup) {
	AuthorityRouter := Router.Group("/authority")
	{
		AuthorityRouter.POST("/createAuthority", v1.CreateAuthority)                 // 创建角色
		AuthorityRouter.DELETE("/deleteAuthority", v1.DeleteAuthority)               // 删除角色
		AuthorityRouter.PUT("/updateAuthority", v1.UpdateAuthority)                  // 更新角色
		AuthorityRouter.POST("/copyAuthority", v1.CopyAuthority)                     // 更新角色
		AuthorityRouter.POST("/getAuthorityList", v1.GetAuthorityList)               // 获取角色列表
		AuthorityRouter.POST("/getAdminAuthorityList", v1.GetAdminAuthorityList)     // 获取员工角色列表
		AuthorityRouter.POST("/getTenancyAuthorityList", v1.GetTenancyAuthorityList) // 获取商户角色列表
		AuthorityRouter.POST("/getGeneralAuthorityList", v1.GetGeneralAuthorityList) // 获取普通用户角色列表
		AuthorityRouter.POST("/setDataAuthority", v1.SetDataAuthority)               // 设置角色资源权限
	}
}

func InitBrandCategoryRouter(Router *gin.RouterGroup) {
	BrandCategoryRouter := Router.Group("/brandCategory")
	{
		BrandCategoryRouter.GET("/getCreateBrandCategoryMap", v1.GetCreateBrandCategoryMap)
		BrandCategoryRouter.GET("/getUpdateBrandCategoryMap/:id", v1.GetUpdateBrandCategoryMap)
		BrandCategoryRouter.POST("/createBrandCategory", v1.CreateBrandCategory)
		BrandCategoryRouter.POST("/getBrandCategoryList", v1.GetBrandCategoryList)
		BrandCategoryRouter.GET("/getBrandCategoryById/:id", v1.GetBrandCategoryById)
		BrandCategoryRouter.POST("/changeBrandCategoryStatus", v1.ChangeBrandCategoryStatus)
		BrandCategoryRouter.PUT("/updateBrandCategory/:id", v1.UpdateBrandCategory)
		BrandCategoryRouter.DELETE("/deleteBrandCategory/:id", v1.DeleteBrandCategory)
	}
}

func InitBrandRouter(Router *gin.RouterGroup) {
	BrandRouter := Router.Group("/brand")
	{
		BrandRouter.GET("/getCreateBrandMap", v1.GetCreateBrandMap)
		BrandRouter.GET("/getUpdateBrandMap/:id", v1.GetUpdateBrandMap)
		BrandRouter.POST("/createBrand", v1.CreateBrand)
		BrandRouter.POST("/getBrandList", v1.GetBrandList)
		BrandRouter.GET("/getBrandById/:id", v1.GetBrandById)
		BrandRouter.POST("/changeBrandStatus", v1.ChangeBrandStatus)
		BrandRouter.PUT("/updateBrand/:id", v1.UpdateBrand)
		BrandRouter.DELETE("/deleteBrand/:id", v1.DeleteBrand)
	}
}

func InitCasbinRouter(Router *gin.RouterGroup) {
	CasbinRouter := Router.Group("/casbin")
	{
		CasbinRouter.POST("/updateCasbin", v1.UpdateCasbin)
		CasbinRouter.POST("/getPolicyPathByAuthorityId", v1.GetPolicyPathByAuthorityId)
	}
}

func InitCategoryRouter(Router *gin.RouterGroup) {
	CategoryRouter := Router.Group("/category")
	{
		CategoryRouter.GET("/getCreateTenancyCategoryMap", v1.GetCreateTenancyCategoryMap)
		CategoryRouter.GET("/getUpdateTenancyCategoryMap/:id", v1.GetUpdateTenancyCategoryMap)
		CategoryRouter.GET("/getCategorySelect", v1.GetCategorySelect)
		CategoryRouter.POST("/createCategory", v1.CreateCategory)
		CategoryRouter.POST("/getCategoryList", v1.GetCategoryList)
		CategoryRouter.GET("/getCategoryById/:id", v1.GetCategoryById)
		CategoryRouter.POST("/changeCategoryStatus", v1.ChangeCategoryStatus)
		CategoryRouter.PUT("/updateCategory/:id", v1.UpdateCategory)
		CategoryRouter.DELETE("/deleteCategory/:id", v1.DeleteCategory)
	}
}

func InitConfigCategoryRouter(Router *gin.RouterGroup) {
	ConfigCategoryRouter := Router.Group("/configCategory")
	{
		ConfigCategoryRouter.GET("/getCreateConfigCategoryMap", v1.GetCreateConfigCategoryMap)
		ConfigCategoryRouter.GET("/getUpdateConfigCategoryMap/:id", v1.GetUpdateConfigCategoryMap)
		ConfigCategoryRouter.GET("/getConfigCategoryList", v1.GetConfigCategoryList)
		ConfigCategoryRouter.POST("/createConfigCategory", v1.CreateConfigCategory)
		ConfigCategoryRouter.GET("/getConfigCategoryById/:id", v1.GetConfigCategoryById)
		ConfigCategoryRouter.PUT("/updateConfigCategory/:id", v1.UpdateConfigCategory)
		ConfigCategoryRouter.POST("/changeConfigCategoryStatus", v1.ChangeConfigCategoryStatus)
		ConfigCategoryRouter.DELETE("/deleteConfigCategory/:id", v1.DeleteConfigCategory)
	}
}

func InitConfigValueRouter(Router *gin.RouterGroup) {
	ConfigValueRouter := Router.Group("/configValue")
	{
		ConfigValueRouter.POST("/saveConfigValue/:category", v1.SaveConfigValue)
	}
}

func InitConfigRouter(Router *gin.RouterGroup) {
	ConfigRouter := Router.Group("/config")
	{
		ConfigRouter.GET("/getConfigMap/:category", v1.GetConfigMap)
		ConfigRouter.GET("/getCreateConfigMap", v1.GetCreateConfigMap)
		ConfigRouter.GET("/getUpdateConfigMap/:id", v1.GetUpdateConfigMap)
		ConfigRouter.POST("/getConfigList", v1.GetConfigList)
		ConfigRouter.POST("/createConfig", v1.CreateConfig)
		ConfigRouter.GET("/getConfigByKey/:key", v1.GetConfigByKey)
		ConfigRouter.GET("/getConfigByID/:id", v1.GetConfigByID)
		ConfigRouter.POST("/changeConfigStatus", v1.ChangeConfigStatus)
		ConfigRouter.PUT("/updateConfig/:id", v1.UpdateConfig)
		ConfigRouter.DELETE("/deleteConfig/:id", v1.DeleteConfig)
	}
}

func InitEmailRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("/email")
	{
		UserRouter.POST("/emailTest", v1.EmailTest) // 发送测试邮件
	}
}

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

func InitMenuRouter(Router *gin.RouterGroup) (R *gin.RouterGroup) {
	MenuRouter := Router.Group("/menu")
	{
		MenuRouter.GET("/getMenu", v1.GetMenu)                    // 获取菜单树
		MenuRouter.GET("/getMenuList", v1.GetMenuList)            // 分页获取基础menu列表
		MenuRouter.POST("/addBaseMenu", v1.AddBaseMenu)           // 新增菜单
		MenuRouter.POST("/getBaseMenuTree", v1.GetBaseMenuTree)   // 获取用户动态路由
		MenuRouter.POST("/addMenuAuthority", v1.AddMenuAuthority) // 增加menu和角色关联关系
		MenuRouter.POST("/getMenuAuthority", v1.GetMenuAuthority) // 获取指定角色menu
		MenuRouter.DELETE("/deleteBaseMenu", v1.DeleteBaseMenu)   // 删除菜单
		MenuRouter.POST("/updateBaseMenu", v1.UpdateBaseMenu)     // 更新菜单
		MenuRouter.POST("/getBaseMenuById", v1.GetBaseMenuById)   // 根据id获取菜单
		ClientMenuRouter := MenuRouter.Group("/merchant")
		{
			ClientMenuRouter.GET("/getClientMenuList", v1.GetClientMenuList) // 分页获取基础menu列表
			ClientMenuRouter.POST("/addBaseMenu", v1.AddBaseMenu)            // 新增菜单
			ClientMenuRouter.POST("/getBaseMenuTree", v1.GetBaseMenuTree)    // 获取用户动态路由
			ClientMenuRouter.POST("/addMenuAuthority", v1.AddMenuAuthority)  // 增加menu和角色关联关系
			ClientMenuRouter.POST("/getMenuAuthority", v1.GetMenuAuthority)  // 获取指定角色menu
			ClientMenuRouter.DELETE("/deleteBaseMenu", v1.DeleteBaseMenu)    // 删除菜单
			ClientMenuRouter.POST("/updateBaseMenu", v1.UpdateBaseMenu)      // 更新菜单
			ClientMenuRouter.POST("/getBaseMenuById", v1.GetBaseMenuById)    // 根据id获取菜单
		}
	}

	return MenuRouter
}

func InitMiniRouter(Router *gin.RouterGroup) {
	MiniRouter := Router.Group("/mini")
	{
		MiniRouter.POST("/createMini", v1.CreateMini)
		MiniRouter.POST("/getMiniList", v1.GetMiniList)
		MiniRouter.POST("/getMiniById", v1.GetMiniById)
		MiniRouter.PUT("/updateMini", v1.UpdateMini)
		MiniRouter.DELETE("/deleteMini", v1.DeleteMini)
	}
}

func InitSysOperationRecordRouter(Router *gin.RouterGroup) {
	SysOperationRecordRouter := Router.Group("/sysOperationRecord")
	{
		SysOperationRecordRouter.POST("/getSysOperationRecordList", v1.GetSysOperationRecordList)           // 获取SysOperationRecord列表
		SysOperationRecordRouter.POST("/createSysOperationRecord", v1.CreateSysOperationRecord)             // 新建SysOperationRecord
		SysOperationRecordRouter.DELETE("/deleteSysOperationRecord", v1.DeleteSysOperationRecord)           // 删除SysOperationRecord
		SysOperationRecordRouter.DELETE("/deleteSysOperationRecordByIds", v1.DeleteSysOperationRecordByIds) // 批量删除SysOperationRecord
		SysOperationRecordRouter.GET("/findSysOperationRecord", v1.FindSysOperationRecord)                  // 根据ID获取SysOperationRecord

	}
}

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

func InitSystemRouter(Router *gin.RouterGroup) {
	SystemRouter := Router.Group("/system")
	{
		SystemRouter.POST("/getSystemConfig", v1.GetSystemConfig) // 获取配置文件内容
		SystemRouter.POST("/setSystemConfig", v1.SetSystemConfig) // 设置配置文件内容
		SystemRouter.POST("/getServerInfo", v1.GetServerInfo)     // 获取服务器信息
		SystemRouter.POST("/reloadSystem", v1.ReloadSystem)       // 重启服务
	}
}

func InitTenancyRouter(Router *gin.RouterGroup) {
	TenancyRouter := Router.Group("/tenancy")
	{
		TenancyRouter.GET("/getTenancyInfo", v1.GetTenancyInfo)            // 登录商户信息
		TenancyRouter.GET("/getTenancies/:code", v1.GetTenanciesByRegion)  // 获取Tenancy列表(不分页)
		TenancyRouter.GET("/getTenancyCount", v1.GetTenancyCount)          // 获取Tenancy对应状态数量
		TenancyRouter.POST("/createTenancy", v1.CreateTenancy)             // 创建Tenancy
		TenancyRouter.POST("/loginTenancy/:id", v1.LoginTenancy)           // 登录商户
		TenancyRouter.POST("/getTenancyList", v1.GetTenanciesList)         // 获取Tenancy列表
		TenancyRouter.GET("/getTenancyById/:id", v1.GetTenancyById)        // 获取单条Tenancy消息
		TenancyRouter.POST("/setTenancyRegion", v1.SetTenancyRegion)       // 设置商户地区
		TenancyRouter.POST("/changeTenancyStatus", v1.ChangeTenancyStatus) // 设置商户显示/隐藏
		TenancyRouter.PUT("/updateTenancy/:id", v1.UpdateTenancy)          // 更新Tenancy
		TenancyRouter.DELETE("/deleteTenancy/:id", v1.DeleteTenancy)       // 删除Tenancy
	}
}

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("/user")
	{
		UserRouter.POST("/registerAdmin", v1.RegisterAdmin)       // 注册
		UserRouter.POST("/registerTenancy", v1.RegisterTenancy)   // 注册
		UserRouter.POST("/changePassword", v1.ChangePassword)     // 修改密码
		UserRouter.POST("/getAdminList", v1.GetAdminList)         // 分页获取管理员列表
		UserRouter.POST("/getTenancyList", v1.GetTenancyList)     // 分页获取商户列表
		UserRouter.POST("/getGeneralList", v1.GetGeneralList)     // 分页获取普通用户列表
		UserRouter.POST("/setUserAuthority", v1.SetUserAuthority) // 设置用户权限
		UserRouter.DELETE("/deleteUser", v1.DeleteUser)           // 删除用户
		UserRouter.PUT("/setUserInfo/:user_id", v1.SetUserInfo)   // 设置用户信息
	}
}
