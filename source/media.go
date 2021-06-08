package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var TenancyMedia = new(media)

type media struct{}

var medias = []model.TenancyMedia{
	{Name: "杜蕾斯", Tag: "杜蕾斯", Url: "http://qmplusimg.henrongyi.top/head.png", Key: "001", SysTenancyID: 0},
	{Name: "杜蕾斯1", Tag: "杜蕾斯1", Url: "http://qmplusimg.henrongyi.top/head.png", Key: "0011", SysTenancyID: 1},
}

func (m *media) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.TenancyMedia{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> tenancy_medias 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&medias).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> tenancy_medias 表初始数据成功!")
		return nil
	})
}
