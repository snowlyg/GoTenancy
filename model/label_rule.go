package model

// LabelRule 自定标签规则
type LabelRule struct {
	Type    uint8   `gorm:"column:type;type:tinyint unsigned;default:0" json:"type"`                // 0=订单数 1=订单金额
	Min     float64 `gorm:"column:min;type:decimal(8,2) unsigned;not null;default:0.00" json:"min"` // 最小值
	Max     float64 `gorm:"column:max;type:decimal(8,2) unsigned;not null;default:0.00" json:"max"` //  最大值
	UserNum uint    `gorm:"column:user_num;type:int unsigned;not null;default:0" json:"userNum"`    // 用户数

	UserLabelID  uint `gorm:"column:user_label_id;" json:"userLabelId"`
	SysTenancyID uint `gorm:"index:sys_tenancy_id;column:sys_tenancy_id;type:int;not null" json:"sysTenancyId"` // 商户 id
}
