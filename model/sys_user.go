package model

import (
	"github.com/snowlyg/go-tenancy/g"
)

type SysUser struct {
	g.TENANCY_MODEL
	Username     string       `json:"userName" gorm:"comment:用户登录名"`
	Password     string       `json:"-"  gorm:"comment:用户登录密码"`
	Authority    SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	AuthorityId  string       `json:"authorityId" gorm:"default:888;comment:用户角色ID"`
	SysAdminInfo SysAdminInfo `json:"admin_info" gorm:"foreignKey:SysAdminID;references:ID;comment:用户信息"`
	SysAdminID   int          `json:"admin_id" gorm:"default:0;comment:用户信息ID"`
}

// `email` varchar(30) NOT NULL DEFAULT '' COMMENT '员工邮箱',
// `phone` varchar(15) NOT NULL DEFAULT '' COMMENT '员工手机号',
// `name` varchar(30) NOT NULL DEFAULT '' COMMENT '员工姓名',
type SysAdminInfo struct {
	g.TENANCY_MODEL
	Email     string `json:"email" gorm:"default:'';comment:员工邮箱" `
	Phone     string `json:"phone" gorm:"default:'';comment:员工手机号" `
	Name      string `json:"name" gorm:"default:'员工姓名';comment:员工姓名" `
	NickName  string `json:"nickName" gorm:"default:系统用户;comment:用户昵称" `
	HeaderImg string `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
}
