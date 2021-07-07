package model

import "github.com/snowlyg/go-tenancy/g"

// Cart 购物车表
type Cart struct {
	g.TENANCY_MODEL

	ProductType       int32  `gorm:"column:product_type;type:tinyint;not null;default:1" json:"productType"`                   // 类型 1=普通产品，2.预售商品
	ProductAttrUnique string `gorm:"column:product_attr_unique;type:varchar(16);not null;default:''" json:"productAttrUnique"` // 商品属性
	CartNum           uint16 `gorm:"column:cart_num;type:smallint unsigned;not null;default:0" json:"cartNum"`                 // 商品数量
	Source            uint8  `gorm:"column:source;type:tinyint unsigned;not null;default:0" json:"source"`                     // 来源 1.直播间,2.预售商品,3.助力商品
	SourceID          uint   `gorm:"column:source_id;type:int unsigned;not null;default:0" json:"sourceId"`                    // 来源关联 id
	IsPay             int    `gorm:"column:is_pay;type:tinyint(1);not null;default:2" json:"isPay"`                            // 2 = 未购买 1 = 已购买
	IsDel             int    `gorm:"column:is_del;type:tinyint(1);not null;default:2" json:"isDel"`                            // 是否删除
	IsNew             int    `gorm:"column:is_new;type:tinyint(1);not null;default:2" json:"isNew"`                            // 是否为立即购买
	IsFail            int    `gorm:"column:is_fail;type:tinyint unsigned;not null;default:2" json:"isFail"`                    // 是否失效

	ProductID    uint `gorm:"index:product_id;column:product_id;type:int unsigned;not null" json:"productId"` // 商品ID
	SysUserID    uint `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:关联标记"`
	SysTenancyID uint `gorm:"index:sys_tenancy_id;column:sys_tenancy_id;type:int;not null" json:"sysTenancyId"` // 商户 id
}
