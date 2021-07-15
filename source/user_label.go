package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var UserLabel = new(userLabel)

type userLabel struct{}

var userLabels = []model.UserLabel{
	{LabelName: "123", Type: model.UserLabelTypeSD},
	{LabelName: "456", Type: model.UserLabelTypeZD},
	{LabelName: "789", Type: model.UserLabelTypeSD, SysTenancyID: 1},
	{LabelName: "567", Type: model.UserLabelTypeZD, SysTenancyID: 1},
}

func (m *userLabel) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.UserLabel{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> user_labels 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&userLabels).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> user_labels 表初始数据成功!")
		return nil
	})
}
