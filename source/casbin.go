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
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/api/deleteApi", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/api/updateApi", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/api/getAllApis", V2: "POST"},

	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/authority/createAuthority", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/authority/deleteAuthority", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/authority/getAuthorityList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/authority/getAdminAuthorityList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/authority/getTenancyAuthorityList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/authority/getGeneralAuthorityList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/authority/setDataAuthority", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/authority/updateAuthority", V2: "PUT"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/authority/copyAuthority", V2: "POST"},

	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/menu/getMenu", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/menu/getMenuList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/menu/addBaseMenu", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/menu/getBaseMenuTree", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/menu/addMenuAuthority", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/menu/getMenuAuthority", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/menu/deleteBaseMenu", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/menu/updateBaseMenu", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/menu/getBaseMenuById", V2: "POST"},

	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/user/changePassword", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/user/getAdminList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/user/getTenancyList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/user/getGeneralList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/user/setUserAuthority", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/user/deleteUser", V2: "DELETE"},

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
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/sysOperationRecord/getSysOperationRecordList", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/sysOperationRecord/deleteSysOperationRecordByIds", V2: "DELETE"},

	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/user/setUserInfo/{user_id:int}", V2: "PUT"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/email/emailTest", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/api/deleteApisByIds", V2: "DELETE"},

	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/tenancy/getTenancies/{code:int}", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/tenancy/getTenancyList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/tenancy/createTenancy", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/tenancy/setTenancyRegion", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/tenancy/getTenancyById", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/tenancy/updateTenancy", V2: "PUT"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/tenancy/deleteTenancy", V2: "DELETE"},

	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/mini/getMiniList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/mini/createMini", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/mini/getMiniById", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/mini/updateMini", V2: "PUT"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/mini/deleteMini", V2: "DELETE"},

	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/config/getConfigList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/config/createConfig", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/config/getConfigByName", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/config/updateConfig", V2: "PUT"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/config/deleteConfig", V2: "DELETE"},

	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/brand/getBrandList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/brand/createBrand", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/brand/getBrandById", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/brand/setBrandCate", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/brand/updateBrand", V2: "PUT"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/brand/deleteBrand", V2: "DELETE"},

	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/brandCategory/getBrandCategoryList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/brandCategory/createBrandCategory", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/brandCategory/getBrandCategoryById", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/brandCategory/updateBrandCategory", V2: "PUT"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/brandCategory/deleteBrandCategory", V2: "DELETE"},

	// 媒体库
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/media/upload", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/media/getFileList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/media/deleteFile", V2: "POST"},

	//商品分类
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/category/getCategoryList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/category/createCategory", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/category/getCategoryById", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/category/updateCategory", V2: "PUT"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/category/deleteCategory", V2: "DELETE"},

	//商品
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/product/getProductList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/product/createProduct", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/product/getProductById", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/product/updateProduct", V2: "PUT"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/admin/product/deleteProduct", V2: "DELETE"},

	//商户用户权限
	// 媒体库
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/admin/media/upload", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/admin/media/getFileList", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/admin/media/deleteFile", V2: "POST"},

	//商品分类
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/admin/category/getCategoryList", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/admin/category/createCategory", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/admin/category/getCategoryById", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/admin/category/updateCategory", V2: "PUT"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/admin/category/deleteCategory", V2: "DELETE"},

	//商品
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/admin/product/getProductList", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/admin/product/createProduct", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/admin/product/getProductById", V2: "POST"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/admin/product/updateProduct", V2: "PUT"},
	{Ptype: "p", V0: TenancyAuthorityId, V1: "/v1/admin/product/deleteProduct", V2: "DELETE"},

	// 普通用户权限
	{Ptype: "p", V0: GeneralAuthorityId, V1: "/v1/auth/logout", V2: "GET"},
	{Ptype: "p", V0: GeneralAuthorityId, V1: "/v1/auth/clean", V2: "GET"},
	// 我的地址
	{Ptype: "p", V0: GeneralAuthorityId, V1: "/v1/general/address/getAddressList", V2: "POST"},
	{Ptype: "p", V0: GeneralAuthorityId, V1: "/v1/general/address/createAddress", V2: "POST"},
	{Ptype: "p", V0: GeneralAuthorityId, V1: "/v1/general/address/getAddressById", V2: "POST"},
	{Ptype: "p", V0: GeneralAuthorityId, V1: "/v1/general/address/updateAddress", V2: "PUT"},
	{Ptype: "p", V0: GeneralAuthorityId, V1: "/v1/general/address/deleteAddress", V2: "DELETE"},
	// 我的发票
	{Ptype: "p", V0: GeneralAuthorityId, V1: "/v1/general/receipt/getReceiptList", V2: "POST"},
	{Ptype: "p", V0: GeneralAuthorityId, V1: "/v1/general/receipt/createReceipt", V2: "POST"},
	{Ptype: "p", V0: GeneralAuthorityId, V1: "/v1/general/receipt/getReceiptById", V2: "POST"},
	{Ptype: "p", V0: GeneralAuthorityId, V1: "/v1/general/receipt/updateReceipt", V2: "PUT"},
	{Ptype: "p", V0: GeneralAuthorityId, V1: "/v1/general/receipt/deleteReceipt", V2: "DELETE"},
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
