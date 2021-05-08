package source

import (
	"time"

	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"

	"gorm.io/gorm"
)

var Admin = new(admin)

type admin struct{}

//  NickName: "超级管理员", HeaderImg: "http://qmplusimg.henrongyi.top/gva_header.jpg",
//  NickName: "QMPlusUser", HeaderImg: "http://qmplusimg.henrongyi.top/1572075907logo.png",
var admins = []model.SysUser{
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Username: "admin", Password: "e10adc3949ba59abbe56e057f20f883e", AuthorityId: AdminAuthorityId, AdminInfo: model.SysAdminInfo{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Email: "admin@admin.com", Phone: "13800138000", Name: "超级管理员", SysUserID: 1}},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Username: "a303176530", Password: "3ec063004a6f31642261936a379fde3d", AuthorityId: TenancyAuthorityId, TenancyInfo: model.SysTenancyInfo{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Email: "a303176530@admin.com", Phone: "13800138000", Name: "商户管理员", SysUserID: 2}},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Username: "oZM5VwD_PCaPKQZ8zRGt-NUdU2uM", Password: "3ec063004a6f31642261936a379fde3d", AuthorityId: GeneralAuthorityId, GeneralInfo: model.SysGeneralInfo{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Email: "a303176530@admin.com", Phone: "13800138000", Name: "商户管理员", SysUserID: 3}},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_users 表数据初始化
func (a *admin) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]model.SysUser{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_users 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&admins).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_users 表初始数据成功!")
		return nil
	})
}
