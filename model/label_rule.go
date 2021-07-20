package model

import "github.com/snowlyg/go-tenancy/g"

const (
	LabelRuleTypeUnknown int = iota
	LabelRuleTypeDDS         //订单数
	LabelRuleTypeDDJE        //订单金额
)

// LabelRule 自定标签规则
type LabelRule struct {
	g.TENANCY_MODEL
	Type    int     `gorm:"column:type;type:tinyint unsigned;default:1" json:"type"`                // 1=订单数 2=订单金额
	Min     float64 `gorm:"column:min;type:decimal(8,2) unsigned;not null;default:0.00" json:"min"` // 最小值
	Max     float64 `gorm:"column:max;type:decimal(8,2) unsigned;not null;default:0.00" json:"max"` //  最大值
	UserNum uint    `gorm:"column:user_num;type:int unsigned;not null;default:0" json:"userNum"`    // 用户数

	UserLabelID  uint `gorm:"column:user_label_id;" json:"userLabelId"`
	SysTenancyID uint `gorm:"index:sys_tenancy_id;column:sys_tenancy_id;type:int;not null" json:"sysTenancyId"` // 商户 id
}
