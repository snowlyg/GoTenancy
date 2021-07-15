package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var UserGroup = new(userGroup)

type userGroup struct{}

var userGroups = []model.UserGroup{
	{GroupName: "123"},
	{GroupName: "456"},
}

func (m *userGroup) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.UserGroup{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> user_groups 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&userGroups).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> user_groups 表初始数据成功!")
		return nil
	})
}
