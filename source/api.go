package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"

	"gorm.io/gorm"
)

var Api = new(api)

type api struct{}

var baseApis = []model.SysApi{

	{Path: "/v1/auth/logout", Description: "退出", ApiGroup: "user", Method: "GET"},
	{Path: "/v1/auth/clean", Description: "清空", ApiGroup: "user", Method: "GET"},
	{Path: "/v1/admin/user/register", Description: "用户注册", ApiGroup: "user", Method: "POST"},
	{Path: "/v1/admin/user/changePassword", Description: "修改密码", ApiGroup: "user", Method: "POST"},
	{Path: "/v1/admin/user/getAdminList", Description: "获取管理员列表", ApiGroup: "user", Method: "POST"},
	{Path: "/v1/admin/user/getTenancyList", Description: "获取商户员工列表", ApiGroup: "user", Method: "POST"},
	{Path: "/v1/admin/user/getGeneralList", Description: "获取普通用户列表", ApiGroup: "user", Method: "POST"},
	{Path: "/v1/admin/user/setUserAuthority", Description: "修改用户角色", ApiGroup: "user", Method: "POST"},
	{Path: "/v1/admin/user/setUserInfo", Description: "设置用户信息", ApiGroup: "user", Method: "PUT"},
	{Path: "/v1/admin/user/deleteUser", Description: "删除用户", ApiGroup: "user", Method: "DELETE"},

	{Path: "/v1/admin/api/createApi", Description: "创建api", ApiGroup: "api", Method: "POST"},
	{Path: "/v1/admin/api/getApiList", Description: "获取api列表", ApiGroup: "api", Method: "POST"},
	{Path: "/v1/admin/api/getApiById", Description: "获取api详细信息", ApiGroup: "api", Method: "POST"},
	{Path: "/v1/admin/api/deleteApi", Description: "删除Api", ApiGroup: "api", Method: "POST"},
	{Path: "/v1/admin/api/updateApi", Description: "更新Api", ApiGroup: "api", Method: "POST"},
	{Path: "/v1/admin/api/getAllApis", Description: "获取所有api", ApiGroup: "api", Method: "POST"},

	{Path: "/v1/admin/authority/createAuthority", Description: "创建角色", ApiGroup: "authority", Method: "POST"},
	{Path: "/v1/admin/authority/deleteAuthority", Description: "删除角色", ApiGroup: "authority", Method: "POST"},
	{Path: "/v1/admin/authority/getAuthorityList", Description: "获取角色列表", ApiGroup: "authority", Method: "POST"},
	{Path: "/v1/admin/authority/setDataAuthority", Description: "设置角色资源权限", ApiGroup: "authority", Method: "POST"},
	{Path: "/v1/admin/authority/updateAuthority", Description: "更新角色信息", ApiGroup: "authority", Method: "PUT"},
	{Path: "/v1/admin/authority/copyAuthority", Description: "拷贝角色", ApiGroup: "authority", Method: "POST"},

	{Path: "/v1/admin/menu/getMenu", Description: "获取菜单树", ApiGroup: "menu", Method: "GET"},
	{Path: "/v1/admin/menu/getMenuList", Description: "分页获取基础menu列表", ApiGroup: "menu", Method: "GET"},
	{Path: "/v1/admin/menu/addBaseMenu", Description: "新增菜单", ApiGroup: "menu", Method: "POST"},
	{Path: "/v1/admin/menu/getBaseMenuTree", Description: "获取用户动态路由", ApiGroup: "menu", Method: "POST"},
	{Path: "/v1/admin/menu/addMenuAuthority", Description: "增加menu和角色关联关系", ApiGroup: "menu", Method: "POST"},
	{Path: "/v1/admin/menu/getMenuAuthority", Description: "获取指定角色menu", ApiGroup: "menu", Method: "POST"},
	{Path: "/v1/admin/menu/deleteBaseMenu", Description: "删除菜单", ApiGroup: "menu", Method: "POST"},
	{Path: "/v1/admin/menu/updateBaseMenu", Description: "更新菜单", ApiGroup: "menu", Method: "POST"},
	{Path: "/v1/admin/menu/getBaseMenuById", Description: "根据id获取菜单", ApiGroup: "menu", Method: "POST"},

	{Path: "/v1/admin/menu/merchant/getClientMenuList", Description: "获取商户菜单", ApiGroup: "menu", Method: "GET"},

	{Path: "/v1/admin/media/getUpdateMediaMap/:id", Description: "获取媒体文件表单", ApiGroup: "media", Method: "GET"},
	{Path: "/v1/admin/media/upload", Description: "上传媒体文件", ApiGroup: "media", Method: "POST"},
	{Path: "/v1/admin/media/getFileList", Description: "获取媒体文件列表", ApiGroup: "media", Method: "POST"},
	{Path: "/v1/admin/media/updateMediaName/:id", Description: "修改媒体文件名称", ApiGroup: "media", Method: "POST"},
	{Path: "/v1/admin/media/deleteFile", Description: "删除媒体文件", ApiGroup: "media", Method: "DELETE"},

	{Path: "/v1/admin/casbin/updateCasbin", Description: "更改角色api权限", ApiGroup: "casbin", Method: "POST"},
	{Path: "/v1/admin/casbin/getPolicyPathByAuthorityId", Description: "获取权限列表", ApiGroup: "casbin", Method: "POST"},
	{Path: "/v1/admin/casbin/casbinTest/:pathParam", Description: "RESTFUL模式测试", ApiGroup: "casbin", Method: "GET"},

	{Path: "/v1/admin/system/getSystemConfig", Description: "获取配置文件内容", ApiGroup: "system", Method: "POST"},
	{Path: "/v1/admin/system/setSystemConfig", Description: "设置配置文件内容", ApiGroup: "system", Method: "POST"},
	{Path: "/v1/admin/system/getServerInfo", Description: "获取服务器信息", ApiGroup: "system", Method: "POST"},

	{Path: "/v1/admin/configValue/saveConfigValue/:category", Description: "保持配置表单", ApiGroup: "configValue", Method: "POST"},

	// 配置
	{Path: "/v1/admin/config/getConfigMap/:category", Description: "获取配置表单", ApiGroup: "config", Method: "GET"},
	{Path: "/v1/admin/config/getCreateConfigMap", Description: "获取配置创建表单", ApiGroup: "config", Method: "GET"},
	{Path: "/v1/admin/config/getUpdateConfigMap/:id", Description: "获取配置编辑表单", ApiGroup: "config", Method: "GET"},
	{Path: "/v1/admin/config/getConfigList", Description: "获取配置项列表", ApiGroup: "config", Method: "POST"},
	{Path: "/v1/admin/config/createConfig", Description: "添加配置项", ApiGroup: "config", Method: "POST"},
	{Path: "/v1/admin/config/getConfigByKey/:key", Description: "获取根据key配置项", ApiGroup: "config", Method: "GET"},
	{Path: "/v1/admin/config/getConfigByID/:id", Description: "获取根据id配置项", ApiGroup: "config", Method: "GET"},
	{Path: "/v1/admin/config/changeConfigStatus", Description: "修改配置状态", ApiGroup: "config", Method: "POST"},
	{Path: "/v1/admin/config/updateConfig/:id", Description: "更新配置项", ApiGroup: "config", Method: "PUT"},
	{Path: "/v1/admin/config/deleteConfig/:id", Description: "删除配置项", ApiGroup: "config", Method: "DELETE"},

	// 配置分类
	{Path: "/v1/admin/configCategory/getCreateConfigCategoryMap", Description: "获取配置分类创建表单", ApiGroup: "configCategory", Method: "GET"},
	{Path: "/v1/admin/configCategory/getUpdateConfigCategoryMap/:id", Description: "获取配置分类编辑表单", ApiGroup: "configCategory", Method: "GET"},
	{Path: "/v1/admin/configCategory/getConfigCategoryList", Description: "获取配置分类列表", ApiGroup: "configCategory", Method: "GET"},
	{Path: "/v1/admin/configCategory/createConfigCategory", Description: "添加配置分类", ApiGroup: "configCategory", Method: "POST"},
	{Path: "/v1/admin/configCategory/changeConfigCategoryStatus", Description: "修改配置分类状态", ApiGroup: "configCategory", Method: "POST"},
	{Path: "/v1/admin/configCategory/getConfigCategoryById/:id", Description: "获取配置分类", ApiGroup: "configCategory", Method: "GET"},
	{Path: "/v1/admin/configCategory/updateConfigCategory/:id", Description: "更新配置分类", ApiGroup: "configCategory", Method: "PUT"},
	{Path: "/v1/admin/configCategory/deleteConfigCategory/:id", Description: "删除配置分类", ApiGroup: "configCategory", Method: "DELETE"},

	// 商户
	{Path: "/v1/admin/tenancy/changeCopyMap/:id", Description: "获取修改商品复制次数map", ApiGroup: "tenancy", Method: "GET"},
	{Path: "/v1/admin/tenancy/getTenancies/:code", Description: "根据地区获取商户", ApiGroup: "tenancy", Method: "GET"},
	{Path: "/v1/admin/tenancy/getTenancyCount", Description: "获取Tenancy对应状态数量", ApiGroup: "tenancy", Method: "GET"},
	{Path: "/v1/admin/tenancy/getTenancyList", Description: "获取商户列表", ApiGroup: "tenancy", Method: "POST"},
	{Path: "/v1/admin/tenancy/loginTenancy", Description: "登录商户", ApiGroup: "tenancy", Method: "POST"},
	{Path: "/v1/admin/tenancy/createTenancy", Description: "添加商户", ApiGroup: "tenancy", Method: "POST"},
	{Path: "/v1/admin/tenancy/setTenancyRegion", Description: "设置商户地区", ApiGroup: "tenancy", Method: "POST"},
	{Path: "/v1/admin/tenancy/setCopyProductNum/:id", Description: "设置商品复制次数", ApiGroup: "tenancy", Method: "POST"},
	{Path: "/v1/admin/tenancy/changeTenancyStatus", Description: "启用/禁用商户", ApiGroup: "tenancy", Method: "POST"},
	{Path: "/v1/admin/tenancy/getTenancyById/:id", Description: "获取商户详细信息", ApiGroup: "tenancy", Method: "GET"},
	{Path: "/v1/admin/tenancy/updateTenancy/:id", Description: "更新商户", ApiGroup: "tenancy", Method: "PUT"},
	{Path: "/v1/admin/tenancy/deleteTenancy/:id", Description: "删除商户", ApiGroup: "tenancy", Method: "DELETE"},

	//商品分类
	{Path: "/v1/admin/productCategory/getCreateProductCategoryMap", Description: "获取商品分类添加表单", ApiGroup: "category", Method: "POST"},
	{Path: "/v1/admin/productCategory/getUpdateProductCategoryMap/:id", Description: "获取商品分类编辑表单", ApiGroup: "category", Method: "POST"},
	{Path: "/v1/admin/productCategory/getProductCategorySelect", Description: "获取商品分类选项", ApiGroup: "category", Method: "GET"},
	{Path: "/v1/admin/productCategory/getProductCategoryList", Description: "获取商品分类列表", ApiGroup: "category", Method: "POST"},
	{Path: "/v1/admin/productCategory/createProductCategory", Description: "添加商品分类", ApiGroup: "category", Method: "POST"},
	{Path: "/v1/admin/productCategory/getProductCategoryById/:id", Description: "获取根据id商品分类", ApiGroup: "category", Method: "GET"},
	{Path: "/v1/admin/productCategory/changeProductCategoryStatus", Description: "修改商品分类状态", ApiGroup: "category", Method: "POST"},
	{Path: "/v1/admin/productCategory/updateProductCategory/:id", Description: "更新商品分类", ApiGroup: "category", Method: "PUT"},
	{Path: "/v1/admin/productCategory/deleteProductCategory/:id", Description: "删除商品分类", ApiGroup: "category", Method: "DELETE"},

	//商品
	{Path: "/v1/admin/product/getEditProductFictiMap/:id", Description: "获取设置虚拟销量表单", ApiGroup: "product", Method: "GET"},
	{Path: "/v1/admin/product/setProductFicti/:id", Description: "设置虚拟销量", ApiGroup: "product", Method: "PUT"},
	{Path: "/v1/admin/product/getProductFilter", Description: "获取商品过滤参数", ApiGroup: "product", Method: "GET"},
	{Path: "/v1/admin/product/changeProductStatus", Description: "强制下架，重新审核", ApiGroup: "product", Method: "POST"},
	{Path: "/v1/admin/product/changeMutilProductStatus", Description: "批量强制下架，重新审核", ApiGroup: "product", Method: "POST"},
	{Path: "/v1/admin/product/getProductList", Description: "获取商品列表", ApiGroup: "product", Method: "POST"},
	{Path: "/v1/admin/product/getProductById/:id", Description: "获取商品详情", ApiGroup: "product", Method: "GET"},
	{Path: "/v1/admin/product/updateProduct/:id", Description: "编辑商品", ApiGroup: "product", Method: "PUT"},

	//品牌分类
	{Path: "/v1/admin/category/getCreateBrandCategoryMap", Description: "获取品牌分类添加表单", ApiGroup: "brandCategory", Method: "POST"},
	{Path: "/v1/admin/category/getUpdateBrandCategoryMap/:id", Description: "获取品牌分类编辑表单", ApiGroup: "brandCategory", Method: "POST"},
	{Path: "/v1/admin/brandCategory/getBrandCategoryList", Description: "获取品牌分类列表", ApiGroup: "brandCategory", Method: "GET"},
	{Path: "/v1/admin/brandCategory/createBrandCategory", Description: "添加品牌分类", ApiGroup: "brandCategory", Method: "POST"},
	{Path: "/v1/admin/brandCategory/getBrandCategoryById/:id", Description: "获取根据id品牌分类", ApiGroup: "brandCategory", Method: "GET"},
	{Path: "/v1/admin/brandCategory/changeBrandCategoryStatus", Description: "修改品牌分类状态", ApiGroup: "brandCategory", Method: "POST"},
	{Path: "/v1/admin/brandCategory/updateBrandCategory/:id", Description: "更新品牌分类", ApiGroup: "brandCategory", Method: "PUT"},
	{Path: "/v1/admin/brandCategory/deleteBrandCategory/:id", Description: "删除品牌分类", ApiGroup: "brandCategory", Method: "DELETE"},

	//品牌
	{Path: "/v1/admin/brand/getBrandList", Description: "获取品牌列表", ApiGroup: "brand", Method: "POST"},
	{Path: "/v1/admin/brand/getCreateBrandMap", Description: "获取品牌添加表单", ApiGroup: "brand", Method: "GET"},
	{Path: "/v1/admin/brand/getUpdateBrandMap/:id", Description: "获取品牌编辑表单", ApiGroup: "brand", Method: "GET"},
	{Path: "/v1/admin/brand/createBrand", Description: "添加品牌", ApiGroup: "brand", Method: "POST"},
	{Path: "/v1/admin/brand/getBrandById/:id", Description: "获取根据id品牌", ApiGroup: "brand", Method: "GET"},
	{Path: "/v1/admin/brand/changeBrandStatus", Description: "修改品牌分类状态", ApiGroup: "brand", Method: "POST"},
	{Path: "/v1/admin/brand/updateBrand/:id", Description: "更新品牌", ApiGroup: "brand", Method: "PUT"},
	{Path: "/v1/admin/brand/deleteBrand/:id", Description: "删除品牌", ApiGroup: "brand", Method: "DELETE"},

	// 小程序
	{Path: "/v1/admin/mini/getMiniList", Description: "获取小程序列表", ApiGroup: "mini", Method: "POST"},
	{Path: "/v1/admin/mini/createMini", Description: "添加小程序", ApiGroup: "mini", Method: "POST"},
	{Path: "/v1/admin/mini/getMiniById", Description: "获取小程序详细信息", ApiGroup: "mini", Method: "POST"},
	{Path: "/v1/admin/mini/updateMini", Description: "更新小程序", ApiGroup: "mini", Method: "PUT"},
	{Path: "/v1/admin/mini/deleteMini", Description: "删除小程序", ApiGroup: "mini", Method: "DELETE"},

	{Path: "/v1/admin/authority/getAdminAuthorityList", Description: "获取员工角色列表", ApiGroup: "authority", Method: "POST"},
	{Path: "/v1/admin/authority/getTenancyAuthorityList", Description: "获取商户角色列表", ApiGroup: "authority", Method: "POST"},
	{Path: "/v1/admin/authority/getGeneralAuthorityList", Description: "获取普通用户角色列表", ApiGroup: "authority", Method: "POST"},

	{Path: "/v1/admin/sysOperationRecord/createSysOperationRecord", Description: "新增操作记录", ApiGroup: "sysOperationRecord", Method: "POST"},
	{Path: "/v1/admin/sysOperationRecord/deleteSysOperationRecord", Description: "删除操作记录", ApiGroup: "sysOperationRecord", Method: "DELETE"},
	{Path: "/v1/admin/sysOperationRecord/findSysOperationRecord", Description: "根据ID获取操作记录", ApiGroup: "sysOperationRecord", Method: "GET"},
	{Path: "/v1/admin/sysOperationRecord/getSysOperationRecordList", Description: "获取操作记录列表", ApiGroup: "sysOperationRecord", Method: "POST"},
	{Path: "/v1/admin/sysOperationRecord/deleteSysOperationRecordByIds", Description: "批量删除操作历史", ApiGroup: "sysOperationRecord", Method: "DELETE"},

	// TODO:商户用户权限
	{Path: "/v1/merchant/config/getConfigMap/:category", Description: "获取配置表单", ApiGroup: "configClient", Method: "GET"},
	// 配置值保存
	{Path: "/v1/merchant/configValue/saveConfigValue/:category", Description: "保持配置表单", ApiGroup: "configValueClient", Method: "POST"},
	//菜单
	{Path: "/v1/merchant/menu/getMenu", Description: "获取菜单树", ApiGroup: "menuClient", Method: "GET"},
	// 商户
	{Path: "/v1/merchant/tenancy/getTenancyInfo", Description: "获取登录商户信息", ApiGroup: "tenancyClient", Method: "GET"},
	{Path: "/v1/merchant/tenancy/getUpdateTenancyMap", Description: "获取登录商户信息表单", ApiGroup: "tenancyClient", Method: "GET"},
	{Path: "/v1/merchant/tenancy/getTenancyCopyCount", Description: "获取商户商品复制次数", ApiGroup: "tenancyClient", Method: "GET"},
	{Path: "/v1/merchant/tenancy/updateTenancy/:id", Description: "保存登录商户信息", ApiGroup: "tenancyClient", Method: "PUT"},
	// 媒体库
	{Path: "/v1/merchant/media/getUpdateMediaMap/:id", Description: "获取媒体文件表单", ApiGroup: "mediaClient", Method: "GET"},
	{Path: "/v1/merchant/media/upload", Description: "上传文件", ApiGroup: "mediaClient", Method: "POST"},
	{Path: "/v1/merchant/media/getFileList", Description: "getFileList", ApiGroup: "mediaClient", Method: "POST"},
	{Path: "/v1/merchant/media/updateMediaName/:id", Description: "修改媒体文件名称", ApiGroup: "mediaClient", Method: "POST"},
	{Path: "/v1/merchant/media/deleteFile", Description: "删除媒体文件", ApiGroup: "mediaClient", Method: "DELETE"},
	//品牌
	{Path: "/v1/merchant/brand/getBrandList", Description: "获取品牌列表", ApiGroup: "brandClient", Method: "GET"},
	//商品分类
	{Path: "/v1/merchant/productCategory/getCreateProductCategoryMap", Description: "获取商品分类添加表单", ApiGroup: "categoryClient", Method: "GET"},
	{Path: "/v1/merchant/productCategory/getUpdateProductCategoryMap/:id", Description: "获取商品分类编辑表单", ApiGroup: "categoryClient", Method: "GET"},
	{Path: "/v1/merchant/productCategory/getProductCategorySelect", Description: "获取商品分类选项", ApiGroup: "categoryClient", Method: "GET"},
	{Path: "/v1/merchant/productCategory/getAdminProductCategorySelect", Description: "获取平台商品分类选项", ApiGroup: "categoryClient", Method: "GET"},
	{Path: "/v1/merchant/productCategory/getProductCategoryList", Description: "获取商品分类列表", ApiGroup: "categoryClient", Method: "POST"},
	{Path: "/v1/merchant/productCategory/createProductCategory", Description: "添加商品分类", ApiGroup: "categoryClient", Method: "POST"},
	{Path: "/v1/merchant/productCategory/getProductCategoryById/:id", Description: "获取根据id商品分类", ApiGroup: "categoryClient", Method: "GET"},
	{Path: "/v1/merchant/productCategory/changeProductCategoryStatus", Description: "修改商品分类状态", ApiGroup: "categoryClient", Method: "POST"},
	{Path: "/v1/merchant/productCategory/updateProductCategory/:id", Description: "更新商品分类", ApiGroup: "categoryClient", Method: "PUT"},
	{Path: "/v1/merchant/productCategory/deleteProductCategory/:id", Description: "删除商品分类", ApiGroup: "categoryClient", Method: "DELETE"},

	//规格参数
	{Path: "/v1/merchant/attrTemplate/getAttrTemplateList", Description: "获取规格参数列表", ApiGroup: "attrTemplateClient", Method: "POST"},
	{Path: "/v1/merchant/attrTemplate/createAttrTemplate", Description: "添加规格参数", ApiGroup: "attrTemplateClient", Method: "POST"},
	{Path: "/v1/merchant/attrTemplate/getAttrTemplateById/:id", Description: "获取规格参数详情", ApiGroup: "attrTemplateClient", Method: "GET"},
	{Path: "/v1/merchant/attrTemplate/updateAttrTemplate/:id", Description: "更新规格参数", ApiGroup: "attrTemplateClient", Method: "PUT"},
	{Path: "/v1/merchant/attrTemplate/deleteAttrTemplate/:id", Description: "删除规格参数", ApiGroup: "attrTemplateClient", Method: "DELETE"},

	//运费模板
	{Path: "/v1/merchant/shippingTemplate/getShippingTemplateSelect", Description: "获取运费模板下拉", ApiGroup: "shippingTemplateClient", Method: "GET"},
	{Path: "/v1/merchant/shippingTemplate/getShippingTemplateList", Description: "获取运费模板列表", ApiGroup: "shippingTemplateClient", Method: "POST"},
	{Path: "/v1/merchant/shippingTemplate/createShippingTemplate", Description: "添加运费模板", ApiGroup: "shippingTemplateClient", Method: "POST"},
	{Path: "/v1/merchant/shippingTemplate/getShippingTemplateById/:id", Description: "获取运费模板详情", ApiGroup: "shippingTemplateClient", Method: "GET"},
	{Path: "/v1/merchant/shippingTemplate/updateShippingTemplate/:id", Description: "更新运费模板", ApiGroup: "shippingTemplateClient", Method: "PUT"},
	{Path: "/v1/merchant/shippingTemplate/deleteShippingTemplate/:id", Description: "删除运费模板", ApiGroup: "attrTemplateClient", Method: "DELETE"},

	//商品
	{Path: "/v1/merchant/product/getEditProductFictiMap/:id", Description: "获取设置虚拟销量表单", ApiGroup: "productClient", Method: "GET"},
	{Path: "/v1/merchant/product/setProductFicti/:id", Description: "设置虚拟销量", ApiGroup: "productClient", Method: "PUT"},
	{Path: "/v1/merchant/product/getProductFilter", Description: "获取商品过滤参数", ApiGroup: "productClient", Method: "GET"},
	{Path: "/v1/merchant/product/changeProductIsShow", Description: "上下架商品", ApiGroup: "productClient", Method: "POST"},
	{Path: "/v1/merchant/product/getProductList", Description: "获取商品列表", ApiGroup: "productClient", Method: "POST"},
	{Path: "/v1/merchant/product/createProduct", Description: "添加商品", ApiGroup: "productClient", Method: "POST"},
	{Path: "/v1/merchant/product/getProductById/:id", Description: "获取商品详情", ApiGroup: "productClient", Method: "GET"},
	{Path: "/v1/merchant/product/updateProduct/:id", Description: "编辑商品", ApiGroup: "productClient", Method: "PUT"},
	{Path: "/v1/merchant/product/restoreProduct/:id", Description: "还原商品", ApiGroup: "productClient", Method: "GET"},
	{Path: "/v1/merchant/product/deleteProduct/:id", Description: "加入回收站", ApiGroup: "productClient", Method: "DELETE"},
	{Path: "/v1/merchant/product/destoryProduct/:id", Description: "删除商品", ApiGroup: "productClient", Method: "DELETE"},

	{Path: "/v1/merchant/sysOperationRecord/getSysOperationRecordList", Description: "获取操作记录列表", ApiGroup: "sysOperationRecordClient", Method: "POST"},

	{Path: "/v1/admin/email/emailTest", Description: "发送测试邮件", ApiGroup: "email", Method: "POST"},
	{Path: "/v1/admin/api/deleteApisByIds", Description: "批量删除api", ApiGroup: "api", Method: "DELETE"},
}

//@description: sys_apis 表数据初始化
func (a *api) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, len(baseApis)}).Find(&[]model.SysApi{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_apis 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&baseApis).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_apis 表初始数据成功!")
		return nil
	})
}
