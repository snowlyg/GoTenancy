package service

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

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
func GetProductFilter(ctx *gin.Context) ([]response.ProductFilter, error) {
	wheres := getProductConditions(ctx)
	var filters []response.ProductFilter
	for _, where := range wheres {
		filter := response.ProductFilter{Name: where.Name, Type: where.Type}
		db := g.TENANCY_DB.Model(&model.Product{})
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
func getProductConditions(ctx *gin.Context) []response.ProductCondition {
	stock := 0
	if config, err := GetTenancyConfigValue("mer_store_stock", multi.GetTenancyId(ctx)); err == nil {
		if value, err := strconv.Atoi(config.Value); err == nil {
			stock = value
		}
	}

	conditions := []response.ProductCondition{
		{Name: "出售中", Type: 1, Conditions: map[string]interface{}{"is_show": 1, "status": 1}},
		{Name: "仓库中", Type: 2, Conditions: map[string]interface{}{"is_show": 2, "status": 1}},

		{Name: "待审核", Type: 6, Conditions: map[string]interface{}{"status": 2}},
		{Name: "审核未通过", Type: 7, Conditions: map[string]interface{}{"status": 3}},
	}

	if multi.IsTenancy(ctx) {
		other := []response.ProductCondition{{Name: "已售罄", Type: 3, Conditions: map[string]interface{}{"is_show": 1, "stock": stock, "status": 1}},
			{Name: "警戒库存", Type: 4, Conditions: map[string]interface{}{"stock": stock, "status": 1}},
			{Name: "回收站", Type: 5, Conditions: map[string]interface{}{"deleted_at is not null": nil}, IsDeleted: true},
		}
		conditions = append(conditions, other...)
	}
	return conditions
}

// getProductConditionByType
func getProductConditionByType(ctx *gin.Context, t int) response.ProductCondition {
	conditions := getProductConditions(ctx)
	for _, condition := range conditions {
		if condition.Type == t {
			return condition
		}
	}
	return conditions[0]
}

// CreateProduct
func CreateProduct(req request.CreateProduct, ctx *gin.Context) (model.Product, error) {

	var product model.Product
	product.BaseProduct = req.BaseProduct
	product.SysTenancyID = multi.GetTenancyId(ctx)
	product.SliderImage = strings.Join(req.SliderImages, ",")
	product.ProductCategoryID = req.CateId
	product.IsHot = g.StatusFalse
	product.IsBenefit = g.StatusFalse
	product.IsBest = g.StatusFalse
	product.IsNew = g.StatusFalse
	product.ProductType = model.GeneralSale
	product.Status = model.AuditProductStatus
	err := g.TENANCY_DB.Create(&product).Error
	if err != nil {
		return model.Product{}, fmt.Errorf("create product %w", err)
	}

	var productCates []model.ProductProductCate
	for _, categoryId := range req.CategoryIds {
		productCate := model.ProductProductCate{ProductID: product.ID, ProductCategoryID: categoryId, SysTenancyID: product.SysTenancyID}
		productCates = append(productCates, productCate)
	}

	err = g.TENANCY_DB.Model(&model.ProductProductCate{}).Create(&productCates).Error
	if err != nil {
		return model.Product{}, fmt.Errorf("create product product cate %w", err)
	}

	err = g.TENANCY_DB.Model(&model.ProductContent{}).Create(map[string]interface{}{
		"Content": req.Content, "Type": product.ProductType, "ProductID": product.ID,
	}).Error
	if err != nil {
		return model.Product{}, fmt.Errorf("create product content %w", err)
	}

	var productAttrs []model.ProductAttr
	for _, attr := range req.Attr {
		productAttr := model.ProductAttr{ProductID: product.ID, AttrName: attr.Value, AttrValues: strings.Join(attr.Detail, ","), Type: product.ProductType}
		productAttrs = append(productAttrs, productAttr)
	}
	err = g.TENANCY_DB.Model(&model.ProductAttr{}).Create(&productAttrs).Error
	if err != nil {
		return model.Product{}, fmt.Errorf("create product attr %w", err)
	}

	var productAttrValues []model.ProductAttrValue
	for _, attrValue := range req.AttrValue {
		detail, err := json.Marshal(attrValue.Detail)
		if err != nil {
			return model.Product{}, fmt.Errorf("json product attr value detail marshal %w", err)
		}
		attrValue.BaseProductAttrValue.Sku = attrValue.Value0
		productAttrValue := model.ProductAttrValue{ProductID: product.ID, BaseProductAttrValue: attrValue.BaseProductAttrValue, Detail: string(detail), Type: product.ProductType}
		productAttrValues = append(productAttrValues, productAttrValue)
	}
	err = g.TENANCY_DB.Model(&model.ProductAttrValue{}).Create(&productAttrValues).Error
	if err != nil {
		return model.Product{}, fmt.Errorf("create product attr value %w", err)
	}

	return product, nil
}

// GetProductByID
func GetProductByID(id uint) (response.ProductDetail, error) {
	var product response.ProductDetail
	err := g.TENANCY_DB.Model(&model.Product{}).
		Select("products.*,sys_tenancies.name as sys_tenancy_name,sys_brands.brand_name as brand_name,product_categories.cate_name as cate_name,product_contents.content as content,shipping_templates.name as temp_name").
		Joins("left join sys_tenancies on products.sys_tenancy_id = sys_tenancies.id").
		Joins("left join sys_brands on products.sys_brand_id = sys_brands.id").
		Joins("left join product_categories on products.product_category_id = product_categories.id").
		Joins("left join product_contents on product_contents.product_id = products.id").
		Joins("left join shipping_templates on products.temp_id = shipping_templates.id").
		Where("products.id = ?", id).
		First(&product).Error
	if err != nil {
		return response.ProductDetail{}, err
	}
	product.SliderImages = strings.Split(product.SliderImage, ",")

	var attrs []model.ProductAttr
	err = g.TENANCY_DB.Model(&model.ProductAttr{}).Where("product_id = ?", id).
		Find(&attrs).Error
	if err != nil {
		return response.ProductDetail{}, err
	}
	values := []request.Value{}
	for _, attr := range attrs {
		value := request.Value{Value: attr.AttrName, Detail: strings.Split(attr.AttrValues, ",")}
		values = append(values, value)
	}
	product.Attr = values

	var attrValues []model.ProductAttrValue
	err = g.TENANCY_DB.Model(&model.ProductAttrValue{}).Where("product_id = ?", id).
		Find(&attrValues).Error
	if err != nil {
		return response.ProductDetail{}, err
	}
	var productAttrValues []request.ProductAttrValue
	for _, attrValue := range attrValues {
		productAttrValue := request.ProductAttrValue{BaseProductAttrValue: attrValue.BaseProductAttrValue, Value0: attrValue.BaseProductAttrValue.Sku}
		err := json.Unmarshal([]byte(attrValue.Detail), &productAttrValue.Detail)
		if err != nil {
			return response.ProductDetail{}, fmt.Errorf("json product attr value detail marshal %w", err)
		}
		productAttrValue.Value0 = attrValue.BaseProductAttrValue.Sku
		productAttrValues = append(productAttrValues, productAttrValue)
	}
	product.AttrValue = productAttrValues
	product.CateId = product.ProductCategoryID
	product.SliderImages = strings.Split(product.SliderImage, ",")

	productCates, err := getProductCatesByProductId(product.ID, product.SysTenancyID)
	if err != nil {
		return response.ProductDetail{}, err
	}
	product.ProductCates = productCates

	var categoryIds []uint
	for _, productCate := range productCates {
		categoryIds = append(categoryIds, productCate.ID)
	}
	product.CategoryIds = categoryIds

	return product, err
}

// GetProductFictiByID
func GetProductFictiByID(id uint) (int32, error) {
	var product response.ProductFicti
	err := g.TENANCY_DB.Model(&model.Product{}).
		Select("ficti").
		Where("products.id = ?", id).
		First(&product).Error
	return product.Ficti, err
}

// UpdateProduct
func UpdateProduct(req request.UpdateProduct, id uint) error {

	product := model.Product{
		BaseProduct: model.BaseProduct{StoreName: req.StoreName, Ficti: req.Ficti, IsBenefit: req.IsBenefit, IsBest: req.IsBest, IsHot: req.IsHot, IsNew: req.IsNew, Status: model.AuditProductStatus}}

	var content model.ProductContent
	if err := g.TENANCY_DB.Model(&model.ProductContent{}).Where("product_id = ?", id).First(&content).Error; err != nil {
		return err
	}

	err := g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Updates(&product).Error; err != nil {
			return err
		}
		if content.ProductID > 0 {
			if err := tx.Model(&model.ProductContent{}).Where("product_id = ?", content.ProductID).Updates(map[string]interface{}{"content": req.Content}).Error; err != nil {
				return err
			}
		} else {
			content = model.ProductContent{Content: req.Content, ProductID: id, Type: req.ProductType}
			if err := tx.Model(&model.ProductContent{}).Create(&content).Error; err != nil {
				return err
			}
		}

		return nil
	})
	return err
}

// ChangeProductStatus
func ChangeProductStatus(changeStatus request.ChangeProductStatus) error {
	return g.TENANCY_DB.Model(&model.Product{}).Where("id = ?", changeStatus.Id).Updates(map[string]interface{}{"status": changeStatus.Status, "refusal": changeStatus.Refusal}).Error
}

// ChangeMutilProductStatus
func ChangeMutilProductStatus(changeStatus request.ChangeMutilProductStatus) error {
	return g.TENANCY_DB.Model(&model.Product{}).Where("id in ?", changeStatus.Id).Updates(map[string]interface{}{"status": changeStatus.Status}).Error
}

// ChangeProductIsShow
func ChangeProductIsShow(changeStatus request.ChangeProductIsShow) error {
	return g.TENANCY_DB.Model(&model.Product{}).Where("id = ?", changeStatus.Id).Updates(map[string]interface{}{"is_show": changeStatus.IsShow}).Error
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
	if err := g.TENANCY_DB.Model(&model.Product{}).Where("id = ?", id).Updates(map[string]interface{}{"ficti": ficti}).Error; err != nil {
		return err
	}
	return err
}

// DeleteProduct
func DeleteProduct(id uint) error {
	var product model.Product
	return g.TENANCY_DB.Where("id = ?", id).Delete(&product).Error
}

// RestoreProduct
func RestoreProduct(id uint) error {
	return g.TENANCY_DB.Model(&model.Product{}).Unscoped().Where("id = ?", id).Updates(map[string]interface{}{"deleted_at": nil}).Error
}

// ForceDeleteProduct
func ForceDeleteProduct(id uint) error {
	var product model.Product
	return g.TENANCY_DB.Unscoped().Where("id = ?", id).Delete(&product).Error
}

// GetProductInfoList
func GetProductInfoList(info request.ProductPageInfo, ctx *gin.Context) ([]response.ProductList, int64, error) {
	var productList []response.ProductList
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.Product{})
	if info.Type != "" {
		t, err := strconv.Atoi(info.Type)
		if err != nil {
			return productList, total, err
		}
		cond := getProductConditionByType(ctx, t)
		if cond.IsDeleted {
			db = db.Unscoped()
		}
		for key, cn := range cond.Conditions {
			if cn == nil {
				db = db.Where(fmt.Sprintf("%s%s", "products.", key))
			} else {
				db = db.Where(fmt.Sprintf("%s%s = ?", "products.", key), cn)
			}
		}
	}

	if multi.IsTenancy(ctx) {
		db = db.Where("products.sys_tenancy_id = ?", multi.GetTenancyId(ctx))
	}
	if info.Keyword != "" {
		db = db.Where(g.TENANCY_DB.Where("products.store_name like ?", info.Keyword+"%").Or("products.store_info like ?", info.Keyword+"%").Or("products.keyword like ?", info.Keyword+"%").Or("products.bar_code like ?", info.Keyword+"%"))
	}

	// 平台分类id
	if info.ProductCategoryId > 0 {
		productIds, err := getProductIdsByProductCategoryId(info.ProductCategoryId, multi.GetTenancyId(ctx))
		if err != nil {
			return productList, total, err
		}
		db = db.Where("products.id in ?", productIds)
	}

	// 平台分类id
	if info.CateId > 0 {
		db = db.Where("products.product_category_id = ?", info.CateId)
	}
	// 平台分类id
	if info.IsGiftBag != "" {
		db = db.Where("products.is_gift_bag = ?", info.IsGiftBag)
	}

	err := db.Count(&total).Error
	if err != nil {
		return productList, total, err
	}
	err = db.Select("products.*,sys_tenancies.name as sys_tenancy_name,sys_brands.brand_name as brand_name,product_categories.cate_name as cate_name").
		Joins("left join sys_tenancies on products.sys_tenancy_id = sys_tenancies.id").
		Joins("left join sys_brands on products.sys_brand_id = sys_brands.id").
		Joins("left join product_categories on products.product_category_id = product_categories.id").
		Limit(limit).Offset(offset).Find(&productList).Error

	for i := 0; i < len(productList); i++ {
		productCates, err := getProductCatesByProductId(productList[i].ID, productList[i].SysTenancyID)
		if err != nil {
			continue
		}
		productList[i].ProductCates = productCates
	}

	return productList, total, err
}
