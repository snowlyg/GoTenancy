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

var authorityMenus = []AuthorityMenus{
	{AdminAuthorityId, 1},
	{AdminAuthorityId, 2},
	{AdminAuthorityId, 3},
	{AdminAuthorityId, 4},
	{AdminAuthorityId, 5},
	{AdminAuthorityId, 6},
	{AdminAuthorityId, 7},
	{AdminAuthorityId, 8},
	{AdminAuthorityId, 9},
	{AdminAuthorityId, 10},
	{AdminAuthorityId, 11},
	{AdminAuthorityId, 12},
	{AdminAuthorityId, 13},
	{AdminAuthorityId, 14},
	{AdminAuthorityId, 15},
	{AdminAuthorityId, 16},
	{AdminAuthorityId, 17},
	{AdminAuthorityId, 18},
	{AdminAuthorityId, 19},
	{AdminAuthorityId, 20},
	{AdminAuthorityId, 21},
	{AdminAuthorityId, 22},
	{AdminAuthorityId, 23},
	{GeneralAuthorityId, 1},
	{GeneralAuthorityId, 2},
	{GeneralAuthorityId, 8},
	{TenancyAuthorityId, 1},
	{TenancyAuthorityId, 2},
	{TenancyAuthorityId, 3},
	{TenancyAuthorityId, 4},
	{TenancyAuthorityId, 5},
	{TenancyAuthorityId, 6},
	{TenancyAuthorityId, 7},
	{TenancyAuthorityId, 8},
	{TenancyAuthorityId, 9},
	{TenancyAuthorityId, 10},
	{TenancyAuthorityId, 11},
	{TenancyAuthorityId, 12},
	{TenancyAuthorityId, 14},
	{TenancyAuthorityId, 15},
	{TenancyAuthorityId, 16},
	{TenancyAuthorityId, 17},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_authority_menus 表数据初始化
func (a *authoritiesMenus) Init() error {
	return g.TENANCY_DB.Table("sys_authority_menus").Transaction(func(tx *gorm.DB) error {
		if tx.Where("sys_authority_authority_id IN ('"+AdminAuthorityId+"', "+GeneralAuthorityId+", "+TenancyAuthorityId+")").Find(&[]AuthorityMenus{}).RowsAffected == 48 {
			color.Danger.Println("\n[Mysql] --> sys_authority_menus 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&authorityMenus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_authority_menus 表初始数据成功!")
		return nil
	})
}
