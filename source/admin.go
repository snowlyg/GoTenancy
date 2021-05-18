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

var birthday, _ = time.Parse("2006-01-02", "1994-11-28")
var admins = []model.SysUser{
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Username: "admin", Password: "e10adc3949ba59abbe56e057f20f883e", AuthorityId: AdminAuthorityId, AdminInfo: model.SysAdminInfo{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Email: "admin@admin.com", Phone: "13800138000", NickName: "超级管理员", SysUserID: 1}},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Username: "a303176530", Password: "3ec063004a6f31642261936a379fde3d", AuthorityId: TenancyAuthorityId, TenancyInfo: model.SysTenancyInfo{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Email: "a303176530@admin.com", Phone: "13800138000", NickName: "商户管理员", SysUserID: 2, SysTenancyID: 1}},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Username: "oZM5VwD_PCaPKQZ8zRGt-NUdU2uM", Password: "3ec063004a6f31642261936a379fde3d", AuthorityId: GeneralAuthorityId, GeneralInfo: model.SysGeneralInfo{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Email: "a303176530@admin.com", Phone: "13800138000", NickName: "C端用户", SysUserID: 3, AvatarUrl: "https://thirdwx.qlogo.cn/mmopen/vi_32/PEyYoZmTJtaJdeYWWibrnDUadmXKVYyTtyRq2nxtWbBic5jJTLTT4KHmox1tNvOicgIXxspgmxicghpCFob1icAIWFw/132", Sex: model.Female, Subscribe: 1, OpenId: "own1t5TysymNUqcZm-8giuEvT68M", UnionId: "oZM5VwCgvGUZvkrnrGrdJZI4e12k", IdCard: "445281199411285861", IsAuth: 0, Birthday: birthday, RealName: "余思琳"}},
}

//@description: sys_users 表数据初始化
func (a *admin) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2, 3}).Find(&[]model.SysUser{}).RowsAffected == 3 {
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
