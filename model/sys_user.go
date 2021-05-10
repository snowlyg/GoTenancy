package model

import (
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/multi"
)

type SysUser struct {
	g.TENANCY_MODEL
	Username    string       `json:"userName" gorm:"comment:用户登录名"`
	Password    string       `json:"-"  gorm:"comment:用户登录密码"`
	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	AuthorityId string       `json:"authorityId" gorm:"default:888;comment:用户角色ID"`

	AdminInfo   SysAdminInfo   `json:"adminInfo" gorm:"foreignKey:SysUserID;references:ID;comment:管理员信息"`
	TenancyInfo SysTenancyInfo `json:"tenancyInfo" gorm:"foreignKey:SysUserID;references:ID;comment:商户信息"`
	GeneralInfo SysGeneralInfo `json:"generalInfo" gorm:"foreignKey:SysUserID;references:ID;comment:普通用户信息"`
}

// IsAdmin
func (su *SysUser) IsAdmin() bool {
	return su.Authority.AuthorityType == multi.AdminAuthority
}

// IsTenancy
func (su *SysUser) IsTenancy() bool {
	return su.Authority.AuthorityType == multi.TenancyAuthority
}

// IsGeneral
func (su *SysUser) IsGeneral() bool {
	return su.Authority.AuthorityType == multi.GeneralAuthority
}

// AuthorityType
func (su *SysUser) AuthorityType() int {
	return su.Authority.AuthorityType
}

type SysAdminInfo struct {
	g.TENANCY_MODEL
	Email     string `json:"email" gorm:"default:'';comment:员工邮箱" `
	Phone     string `json:"phone" gorm:"default:'';comment:员工手机号" `
	Name      string `json:"name" gorm:"default:'员工姓名';comment:员工姓名" `
	HeaderImg string `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	SysUserID int    `json:"sys_user_id" form:"sys_user_id" gorm:"column:sys_user_id;comment:关联标记"`
}

type SysTenancyInfo struct {
	g.TENANCY_MODEL
	Email        string `json:"email" gorm:"default:'';comment:员工邮箱" `
	Phone        string `json:"phone" gorm:"default:'';comment:员工手机号" `
	Name         string `json:"name" gorm:"default:'员工姓名';comment:员工姓名" `
	HeaderImg    string `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	SysUserID    int    `json:"sys_user_id" form:"sys_user_id" gorm:"column:sys_user_id;comment:关联标记"`
	SysTenancyID int    `json:"sys_tenancy_id" form:"sys_tenancy_id" gorm:"column:sys_tenancy_id;comment:关联标记"`
}

type SysGeneralInfo struct {
	g.TENANCY_MODEL
	Email        string `json:"email" gorm:"default:'';comment:员工邮箱" `
	Phone        string `json:"phone" gorm:"default:'';comment:员工手机号" `
	Name         string `json:"name" gorm:"default:'员工姓名';comment:员工姓名" `
	HeaderImg    string `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	SysUserID    int    `json:"sys_user_id" form:"sys_user_id" gorm:"column:sys_user_id;comment:关联标记"`
	SysTenancyID int    `json:"sys_tenancy_id" form:"sys_tenancy_id" gorm:"column:sys_tenancy_id;comment:关联标记"`
}
