package service

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/multi"
	"gorm.io/gorm"
)

func GetEditProductFictiMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	ficti, err := GetProductFictiByID(id)
	if err != nil {
		return Form{}, err
	}
	formStr = fmt.Sprintf(`{"rule":[{"type":"input","field":"number","value":"%s","title":"现有虚拟销量","props":{"type":"text","placeholder":"请输入现有虚拟销量","readonly":true}},{"type":"radio","field":"type","value":1,"title":"修改类型","props":{},"options":[{"value":1,"label":"增加"},{"value":2,"label":"减少"}]},{"type":"inputNumber","field":"ficti","value":0,"title":"修改虚拟销量数","props":{"placeholder":"请输入修改虚拟销量数"}}],"action":"","method":"PUT","title":"修改虚拟销量数","config":{}}`, strconv.FormatInt(int64(ficti), 10))

	err = json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	form.SetAction(fmt.Sprintf("%s/%d", "product/setProductFicti", id), ctx)
	return form, err
}

// 出售中 1: is_show' => 1, 'status' => 1
// 仓库中 2:'is_show' => 2, 'status' => 1
// 3,4,5 商户才有
// 已售罄 3:'is_show' => 1, 'stock' => 0, 'status' => 1
// 警戒库存 4:'stock' => $stock ? $stock : 0, 'status' => 1
// 回收站 5:'deleted_at' => not null
// 待审核 6:'status' => 2
// 审核未通过 7:'status' => 3

// GetProductFilter
func GetProductFilter(ctx *gin.Context) ([]response.TenancyProductFilter, error) {
	wheres := getProductConditions(ctx)
	var filters []response.TenancyProductFilter
	for _, where := range wheres {
		filter := response.TenancyProductFilter{Name: where.Name, Type: where.Type}
		db := g.TENANCY_DB.Model(&model.TenancyProduct{})
		// 显示软删除数据
		if where.IsDeleted {
			db = db.Unscoped()
		}

		if where.Conditions != nil && len(where.Conditions) > 0 {
			for key, cn := range where.Conditions {
				if cn == nil {
					db = db.Where(key)
				} else {
					db = db.Where(fmt.Sprintf("%s = ?", key), cn)
				}
			}
		}

		err := db.Count(&filter.Count).Error
		if err != nil {
			return filters, err
		}
		filters = append(filters, filter)
	}

	return filters, nil
}

// getProductConditions
func getProductConditions(ctx *gin.Context) []response.TenancyProductCondition {
	stock := 0
	if config, err := GetTenancyConfigValue("mer_store_stock", multi.GetTenancyId(ctx)); err == nil {
		if value, err := strconv.Atoi(config.Value); err == nil {
			stock = value
		}
	}

	conditions := []response.TenancyProductCondition{
		{Name: "出售中", Type: 1, Conditions: map[string]interface{}{"is_show": 1, "status": 1}},
		{Name: "仓库中", Type: 2, Conditions: map[string]interface{}{"is_show": 2, "status": 1}},

		{Name: "待审核", Type: 6, Conditions: map[string]interface{}{"status": 2}},
		{Name: "审核未通过", Type: 7, Conditions: map[string]interface{}{"status": 3}},
	}

	if multi.IsTenancy(ctx) {
		other := []response.TenancyProductCondition{{Name: "已售罄", Type: 3, Conditions: map[string]interface{}{"is_show": 1, "stock": stock, "status": 1}},
			{Name: "警戒库存", Type: 4, Conditions: map[string]interface{}{"stock": stock, "status": 1}},
			{Name: "回收站", Type: 5, Conditions: map[string]interface{}{"deleted_at is not null": nil}, IsDeleted: true},
		}
		conditions = append(conditions, other...)
	}
	return conditions
}

// getProductConditionByType
func getProductConditionByType(ctx *gin.Context, t int) response.TenancyProductCondition {
	conditions := getProductConditions(ctx)
	for _, condition := range conditions {
		if condition.Type == t {
			return condition
		}
	}
	return conditions[0]
}

// CreateProduct
func CreateProduct(product model.TenancyProduct, ctx *gin.Context) (model.TenancyProduct, error) {
	product.SysTenancyID = multi.GetTenancyId(ctx)
	err := g.TENANCY_DB.Create(&product).Error
	return product, err
}

// GetProductByID
func GetProductByID(id uint) (response.TenancyProductDetail, error) {
	var product response.TenancyProductDetail
	err := g.TENANCY_DB.Model(&model.TenancyProduct{}).
		Select("tenancy_products.*,sys_tenancies.name as sys_tenancy_name,sys_brands.brand_name as brand_name,tenancy_categories.cate_name as cate_name,tenancy_product_contents.content as content").
		Joins("left join sys_tenancies on tenancy_products.sys_tenancy_id = sys_tenancies.id").
		Joins("left join sys_brands on tenancy_products.sys_brand_id = sys_brands.id").
		Joins("left join tenancy_categories on tenancy_products.tenancy_category_id = tenancy_categories.id").
		Joins("left join tenancy_product_contents on tenancy_product_contents.tenancy_product_id = tenancy_products.id").
		Where("tenancy_products.id = ?", id).
		First(&product).Error
	return product, err
}

// GetProductFictiByID
func GetProductFictiByID(id uint) (int32, error) {
	var product response.TenancyProductFicti
	err := g.TENANCY_DB.Model(&model.TenancyProduct{}).
		Select("ficti").
		Where("tenancy_products.id = ?", id).
		First(&product).Error
	return product.Ficti, err
}

// UpdateProduct
func UpdateProduct(req request.UpdateTenancyProduct, id uint) error {
	content := model.TenancyProductContent{Content: req.Content, TenancyProductID: id, Type: req.ProductType}
	product := model.TenancyProduct{BaseTenancyProduct: model.BaseTenancyProduct{StoreName: req.StoreName, Ficti: req.Ficti, IsBenefit: req.IsBenefit, IsBest: req.IsBest, IsHot: req.IsHot, IsNew: req.IsNew}}
	err := g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Updates(&product).Error; err != nil {
			return err
		}
		if err := tx.Where("tenancy_product_id = ?", id).FirstOrCreate(&content).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// ChangeProductStatus
func ChangeProductStatus(changeStatus request.ChangeProductStatus) error {
	return g.TENANCY_DB.Model(&model.TenancyProduct{}).Where("id in ?", changeStatus.Id).Updates(map[string]interface{}{"status": changeStatus.Status, "refusal": changeStatus.Refusal}).Error
}

// SetProductFicti
func SetProductFicti(req request.SetProductFicti, id uint) error {
	ficti, err := GetProductFictiByID(id)
	if err != nil {
		return err
	}
	// 增加
	if req.Type == 1 {
		ficti = ficti + req.Ficti
	} else if req.Type == 2 {
		if ficti <= req.Ficti {
			ficti = 0
		} else {
			ficti = ficti - req.Ficti
		}
	}
	if err := g.TENANCY_DB.Model(&model.TenancyProduct{}).Where("id = ?", id).Updates(map[string]interface{}{"ficti": ficti}).Error; err != nil {
		return err
	}
	return err
}

// DeleteProduct
func DeleteProduct(id uint) error {
	var product model.TenancyProduct
	return g.TENANCY_DB.Where("id = ?", id).Delete(&product).Error
}

// GetProductInfoList
func GetProductInfoList(info request.TenancyProductPageInfo, ctx *gin.Context) ([]response.TenancyProductList, int64, error) {
	var tenancyList []response.TenancyProductList
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.TenancyProduct{})
	t, err := strconv.Atoi(info.Type)
	if err != nil {
		return tenancyList, total, err
	}
	cond := getProductConditionByType(ctx, t)
	if cond.IsDeleted {
		db = db.Unscoped()
	}
	for key, cn := range cond.Conditions {
		if cn == nil {
			db = db.Where(fmt.Sprintf("%s%s", "tenancy_products.", key))
		} else {
			db = db.Where(fmt.Sprintf("%s%s = ?", "tenancy_products.", key), cn)
		}
	}
	if multi.IsTenancy(ctx) {
		db = db.Where("tenancy_products.sys_tenancy_id = ?", multi.GetTenancyId(ctx))
	}
	if info.Keyword != "" {
		db = db.Where(g.TENANCY_DB.Where("tenancy_products.store_name like ?", info.Keyword+"%").Or("tenancy_products.store_info like ?", info.Keyword+"%").Or("tenancy_products.keyword like ?", info.Keyword+"%").Or("tenancy_products.bar_code like ?", info.Keyword+"%"))
	}
	if info.ProductCategoryId > 0 {
		db = db.Where("Product_products.tenancy_category_id = ?", info.ProductCategoryId)
	}

	err = db.Count(&total).Error
	if err != nil {
		return tenancyList, total, err
	}
	err = db.Select("tenancy_products.*,sys_tenancies.name as sys_tenancy_name,sys_brands.brand_name as brand_name,tenancy_categories.cate_name as cate_name").
		Joins("left join sys_tenancies on tenancy_products.sys_tenancy_id = sys_tenancies.id").
		Joins("left join sys_brands on tenancy_products.sys_brand_id = sys_brands.id").
		Joins("left join tenancy_categories on tenancy_products.tenancy_category_id = tenancy_categories.id").
		Limit(limit).Offset(offset).Find(&tenancyList).Error
	return tenancyList, total, err
}
