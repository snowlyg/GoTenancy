package source

import (
	"github.com/gookit/color"
	uuid "github.com/satori/go.uuid"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var Mini = new(mini)

type mini struct{}

var minis = []model.SysMini{
	{UUID: uuid.NewV4(), Name: "中德澳上线护理商城", AppID: "YJ3s1abt7MAfT6gWVKoD", AppSecret: "tRE49zaf5NCm6PidFZoaFg3u4WCHDok7fxgL63yV0pF4AM", Remark: "中德澳上线护理商城"},
}

func (m *mini) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.SysMini{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> sys_minis 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&minis).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_minis 表初始数据成功!")
		return nil
	})
}
