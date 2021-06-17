package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"gorm.io/gorm"
)

var AuthoritiesMenus = new(authoritiesMenus)

type authoritiesMenus struct{}

type AuthorityMenus struct {
	AuthorityId string `gorm:"column:sys_authority_authority_id"`
	BaseMenuId  uint   `gorm:"column:sys_base_menu_id"`
}

func getAuthorityMenus() []AuthorityMenus {
	var authorityMenus = []AuthorityMenus{}
	for _, menu := range menus {
		if menu.IsTenancy == g.StatusTrue {
			authorityMenus = append(authorityMenus, AuthorityMenus{TenancyAuthorityId, menu.ID})
		} else {
			authorityMenus = append(authorityMenus, AuthorityMenus{AdminAuthorityId, menu.ID})
		}
	}
	return authorityMenus
}

//@description: sys_authority_menus 表数据初始化
func (a *authoritiesMenus) Init() error {
	return g.TENANCY_DB.Table("sys_authority_menus").Transaction(func(tx *gorm.DB) error {
		if tx.Where("sys_authority_authority_id IN ('"+AdminAuthorityId+"', "+TenancyAuthorityId+")").Find(&[]AuthorityMenus{}).RowsAffected == 48 {
			color.Danger.Println("\n[Mysql] --> sys_authority_menus 表的初始数据已存在!")
			return nil
		}
		authorityMenus := getAuthorityMenus()
		if err := tx.Create(&authorityMenus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_authority_menus 表初始数据成功!")
		return nil
	})
}
