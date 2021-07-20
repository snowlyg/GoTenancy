package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var LabelRule = new(labelRule)

type labelRule struct{}

var labelRules = []model.LabelRule{
	{Type: model.LabelRuleTypeDDJE, Min: 999999.99, Max: 999999.99, UserNum: 0, SysTenancyID: 1, UserLabelID: 4},
}

func (m *labelRule) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.LabelRule{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> label_rules 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&labelRules).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> label_rules 表初始数据成功!")
		return nil
	})
}
