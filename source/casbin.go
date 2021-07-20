package source

import (
	"github.com/snowlyg/go-tenancy/g"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gookit/color"
	"gorm.io/gorm"
)

var Casbin = new(casbin)

type casbin struct{}

var carbines = []gormadapter.CasbinRule{
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/user/registerTenancy", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/user/registerAdmin", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/auth/logout", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/auth/clean", V2: "GET"},

	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/api/createApi", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/api/getApiList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/api/getApiById", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/api/deleteApi", V2: "DELETE"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/api/updateApi", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/api/getAllApis", V2: "POST"},

	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/authority/createAuthority", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/authority/deleteAuthority", V2: "DELETE"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/authority/getAuthorityList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/authority/getAdminAuthorityList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/authority/getTenancyAuthorityList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/authority/getGeneralAuthorityList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/authority/setDataAuthority", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/authority/updateAuthority", V2: "PUT"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/authority/copyAuthority", V2: "POST"},
	// 菜单
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/menu/getMenu", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/menu/getAddMenuMap", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/menu/getAddTenancyMenuMap", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/menu/getEditMenuMap/:id", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/menu/getMenuList", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/menu/addBaseMenu", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/menu/addTenancyBaseMenu", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/menu/getBaseMenuTree", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/menu/addMenuAuthority", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/menu/getMenuAuthority", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/menu/deleteBaseMenu/:id", V2: "DELETE"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/menu/updateBaseMenu/:id", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/menu/getBaseMenuById/:id", V2: "POST"},

	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/menu/merchant/getClientMenuList", V2: "GET"},

	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/user/changePassword", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/user/changeProfile", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/user/getAdminList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/user/getTenancyList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/user/setUserAuthority", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/user/deleteUser", V2: "DELETE"},

	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/cuser/getGeneralList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/cuser/getGeneralDetail/:id", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/cuser/getOrderList/:id", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/cuser/getBillList/:id", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/cuser/setNowMoney/:id", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/cuser/setNowMoneyMap/:id", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/cuser/setUserGroup/:id", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/cuser/setUserGroupMap/:id", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/cuser/setUserLabel/:id", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/cuser/setUserLabelMap/:id", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/cuser/editUser/:id", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/cuser/editUserMap/:id", V2: "GET"},

	// 批量用户分组和标签设置
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/cuser/batchSetUserGroupMap", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/cuser/batchSetUserGroup", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/cuser/batchSetUserLabelMap", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/cuser/batchSetUserLabel", V2: "POST"},

	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/casbin/updateCasbin", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/casbin/getPolicyPathByAuthorityId", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/casbin/casbinTest/:pathParam", V2: "GET"},

	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/system/getSystemConfig", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/system/setSystemConfig", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/system/getServerInfo", V2: "POST"},

	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/sysOperationRecord/createSysOperationRecord", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/sysOperationRecord/deleteSysOperationRecord", V2: "DELETE"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/sysOperationRecord/updateSysOperationRecord", V2: "PUT"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/sysOperationRecord/findSysOperationRecord", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/sysOperationRecord/getSysOperationRecordList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/sysOperationRecord/deleteSysOperationRecordByIds", V2: "DELETE"},

	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/user/setUserInfo/:user_id", V2: "PUT"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/email/emailTest", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/api/deleteApisByIds", V2: "DELETE"},

	// 商户
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/tenancy/getTenancySelect", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/tenancy/changeCopyMap/:id", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/tenancy/getTenancies/:code", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/tenancy/getTenancyCount", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/tenancy/getTenancyList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/tenancy/loginTenancy/:id", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/tenancy/setCopyProductNum/:id", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/tenancy/createTenancy", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/tenancy/setTenancyRegion", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/tenancy/changeTenancyStatus", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/tenancy/getTenancyById/:id", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/tenancy/updateTenancy/:id", V2: "PUT"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/tenancy/deleteTenancy/:id", V2: "DELETE"},

	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/mini/getMiniList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/mini/createMini", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/mini/getMiniById", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/mini/updateMini", V2: "PUT"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/mini/deleteMini", V2: "DELETE"},

	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/configValue/saveConfigValue/:category", V2: "POST"},

	// 配置
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/config/getConfigMap/:category", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/config/getCreateConfigMap", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/config/getUpdateConfigMap/:id", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/config/getConfigList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/config/createConfig", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/config/getConfigByKey/:key", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/config/getConfigByID/:id", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/config/changeConfigStatus", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/config/updateConfig/:id", V2: "PUT"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/config/deleteConfig/:id", V2: "DELETE"},

	// 配置分类
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/configCategory/getCreateConfigCategoryMap", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/configCategory/getUpdateConfigCategoryMap/:id", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/configCategory/getConfigCategoryList", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/configCategory/createConfigCategory", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/configCategory/changeConfigCategoryStatus", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/configCategory/getConfigCategoryById/:id", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/configCategory/updateConfigCategory/:id", V2: "PUT"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/configCategory/deleteConfigCategory/:id", V2: "DELETE"},

	// 物流公司
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/express/getCreateExpressMap", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/express/getUpdateExpressMap/:id", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/express/getExpressList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/express/createExpress", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/express/getExpressById/:id", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/express/changeExpressStatus", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/express/updateExpress/:id", V2: "PUT"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/express/deleteExpress/:id", V2: "DELETE"},

	// 品牌
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/brand/getCreateBrandMap", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/brand/getUpdateBrandMap/:id", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/brand/getBrandList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/brand/createBrand", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/brand/getBrandById/:id", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/brand/changeBrandStatus", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/brand/updateBrand/:id", V2: "PUT"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/brand/deleteBrand/:id", V2: "DELETE"},

	// 品牌分类
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/brandCategory/getCreateBrandCategoryMap", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/brandCategory/getUpdateBrandCategoryMap/:id", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/brandCategory/getBrandCategoryList", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/brandCategory/createBrandCategory", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/brandCategory/getBrandCategoryById/:id", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/brandCategory/changeBrandCategoryStatus", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/brandCategory/updateBrandCategory/:id", V2: "PUT"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/brandCategory/deleteBrandCategory/:id", V2: "DELETE"},

	// 媒体库
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/media/getUpdateMediaMap/:id", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/media/upload", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/media/getFileList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/media/updateMediaName/:id", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/media/deleteFile", V2: "DELETE"},

	//商品分类
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/productCategory/getCreateProductCategoryMap", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/productCategory/getUpdateProductCategoryMap/:id", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/productCategory/getProductCategorySelect", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/productCategory/getProductCategoryList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/productCategory/createProductCategory", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/productCategory/getProductCategoryById/:id", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/productCategory/changeProductCategoryStatus", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/productCategory/updateProductCategory/:id", V2: "PUT"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/productCategory/deleteProductCategory/:id", V2: "DELETE"},

	//商品
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/product/getEditProductFictiMap/:id", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/product/setProductFicti/:id", V2: "PUT"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/product/getProductFilter", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/product/getProductList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/product/changeProductStatus", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/product/changeMutilProductStatus", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/product/getProductById/:id", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/product/updateProduct/:id", V2: "PUT"},
	//订单
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/order/getOrderList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/order/getOrderChart", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/order/getOrderById/:id", V2: "GET"},
	//退款订单
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/refundOrder/getRefundOrderList", V2: "POST"},

	// 用户分组
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/userGroup/getCreateUserGroupMap", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/userGroup/getUpdateUserGroupMap/:id", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/userGroup/getUserGroupList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/userGroup/createUserGroup", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/userGroup/changeUserGroupStatus", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/userGroup/updateUserGroup/:id", V2: "PUT"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/userGroup/deleteUserGroup/:id", V2: "DELETE"},

	// 用户标签
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/userLabel/getCreateUserLabelMap", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/userLabel/getUpdateUserLabelMap/:id", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/userLabel/getUserLabelList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/userLabel/createUserLabel", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/userLabel/changeUserLabelStatus", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/userLabel/updateUserLabel/:id", V2: "PUT"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/userLabel/deleteUserLabel/:id", V2: "DELETE"},

	// TODO:商户用户权限
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/auth/logout", V2: "GET"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/auth/clean", V2: "GET"},
	// 品牌
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/brand/getBrandList", V2: "GET"},
	// 配置
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/config/getConfigMap/:category", V2: "GET"},
	// 配置值保存
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/configValue/saveConfigValue/:category", V2: "POST"},
	//菜单
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/menu/getMenu", V2: "GET"},
	// 商户
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/tenancy/getTenancyInfo", V2: "GET"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/tenancy/getUpdateTenancyMap", V2: "GET"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/tenancy/updateTenancy/:id", V2: "PUT"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/tenancy/getTenancyCopyCount", V2: "GET"},
	// 媒体库
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/media/getUpdateMediaMap/:id", V2: "GET"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/media/upload", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/media/getFileList", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/media/updateMediaName/:id", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/media/deleteFile", V2: "DELETE"},

	//商品分类
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/productCategory/getCreateProductCategoryMap", V2: "GET"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/productCategory/getUpdateProductCategoryMap/:id", V2: "GET"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/productCategory/getProductCategorySelect", V2: "GET"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/productCategory/getAdminProductCategorySelect", V2: "GET"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/productCategory/getProductCategoryList", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/productCategory/createProductCategory", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/productCategory/getProductCategoryById/:id", V2: "GET"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/productCategory/changeProductCategoryStatus", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/productCategory/updateProductCategory/:id", V2: "PUT"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/productCategory/deleteProductCategory/:id", V2: "DELETE"},

	//规格参数
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/attrTemplate/getAttrTemplateList", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/attrTemplate/createAttrTemplate", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/attrTemplate/getAttrTemplateById/:id", V2: "GET"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/attrTemplate/updateAttrTemplate/:id", V2: "PUT"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/attrTemplate/deleteAttrTemplate/:id", V2: "DELETE"},

	//运费模板
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/shippingTemplate/getShippingTemplateSelect", V2: "GET"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/shippingTemplate/getShippingTemplateList", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/shippingTemplate/createShippingTemplate", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/shippingTemplate/getShippingTemplateById/:id", V2: "GET"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/shippingTemplate/updateShippingTemplate/:id", V2: "PUT"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/shippingTemplate/deleteShippingTemplate/:id", V2: "DELETE"},

	//商品
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/product/getEditProductFictiMap/:id", V2: "GET"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/product/setProductFicti/:id", V2: "PUT"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/product/getProductFilter", V2: "GET"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/product/changeProductIsShow", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/product/getProductList", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/product/createProduct", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/product/getProductById/:id", V2: "GET"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/product/updateProduct/:id", V2: "PUT"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/product/restoreProduct/:id", V2: "GET"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/product/deleteProduct/:id", V2: "DELETE"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/product/destoryProduct/:id", V2: "DELETE"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/sysOperationRecord/getSysOperationRecordList", V2: "POST"},
	//订单
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/order/deliveryOrderMap/:id", V2: "GET"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/order/getOrderRemarkMap/:id", V2: "GET"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/order/getEditOrderMap/:id", V2: "GET"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/order/getOrderList", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/order/getOrderChart", V2: "GET"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/order/getOrderFilter", V2: "GET"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/order/getOrderById/:id", V2: "GET"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/order/getOrderRecord/:id", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/order/deliveryOrder/:id", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/order/remarkOrder/:id", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/order/updateOrder/:id", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/order/deleteOrder/:id", V2: "DELETE"},
	//退款订单
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/refundOrder/getRefundOrderRemarkMap/:id", V2: "GET"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/refundOrder/getRefundOrderMap/:id", V2: "GET"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/refundOrder/remarkRefundOrder/:id", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/refundOrder/getRefundOrderList", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/refundOrder/auditRefundOrder/:id", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/refundOrder/getRefundOrderRecord/:id", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/refundOrder/deleteRefundOrder/:id", V2: "DELETE"},
	// 用户手动标签
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/userLabel/getLabelList", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/userLabel/getCreateUserLabelMap", V2: "GET"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/userLabel/getUpdateUserLabelMap/:id", V2: "GET"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/userLabel/createUserLabel", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/userLabel/updateUserLabel/:id", V2: "PUT"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/userLabel/deleteUserLabel/:id", V2: "DELETE"},
	// 用户自动标签
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/userLabel/auto/getLabelList", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/userLabel/auto/createUserLabel", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/userLabel/auto/updateUserLabel/:id", V2: "PUT"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/userLabel/auto/deleteUserLabel/:id", V2: "DELETE"},

	// 物流信息
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/merchant/express/getExpressByCode/:code", V2: "GET"},

	// TODO:普通用户权限
	{Ptype: "p", V0: GeneralAuthorityId, V1: "/v1/auth/logout", V2: "GET"},
	{Ptype: "p", V0: GeneralAuthorityId, V1: "/v1/auth/clean", V2: "GET"},

	// 我的地址
	{Ptype: "p", V0: GeneralAuthorityId, V1: "/v1/user/address/getAddressList", V2: "POST"},
	{Ptype: "p", V0: GeneralAuthorityId, V1: "/v1/user/address/createAddress", V2: "POST"},
	{Ptype: "p", V0: GeneralAuthorityId, V1: "/v1/user/address/getAddressById/:id", V2: "GET"},
	{Ptype: "p", V0: GeneralAuthorityId, V1: "/v1/user/address/updateAddress/:id", V2: "PUT"},
	{Ptype: "p", V0: GeneralAuthorityId, V1: "/v1/user/address/deleteAddress/:id", V2: "DELETE"},
	// 我的发票
	{Ptype: "p", V0: GeneralAuthorityId, V1: "/v1/user/receipt/getReceiptList", V2: "POST"},
	{Ptype: "p", V0: GeneralAuthorityId, V1: "/v1/user/receipt/createReceipt", V2: "POST"},
	{Ptype: "p", V0: GeneralAuthorityId, V1: "/v1/user/receipt/getReceiptById/:id", V2: "GET"},
	{Ptype: "p", V0: GeneralAuthorityId, V1: "/v1/user/receipt/updateReceipt/:id", V2: "PUT"},
	{Ptype: "p", V0: GeneralAuthorityId, V1: "/v1/user/receipt/deleteReceipt/:id", V2: "DELETE"},
}

//Init casbin_rule 表数据初始化
func (c *casbin) Init() error {
	g.TENANCY_DB.AutoMigrate(gormadapter.CasbinRule{})
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Find(&[]gormadapter.CasbinRule{}).RowsAffected == 154 {
			color.Danger.Println("\n[Mysql] --> casbin_rule 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&carbines).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> casbin_rule 表初始数据成功!")
		return nil
	})
}
