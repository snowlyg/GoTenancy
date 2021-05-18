package model

import (
	"github.com/snowlyg/go-tenancy/g"
)

// TenancyAddress 用户收货地址
type TenancyAddress struct {
	g.TENANCY_MODEL
	Name          string `json:"name" gorm:"default:'姓名';comment:姓名" `
	Phone         string `json:"phone" gorm:"default:'';comment:手机号" `
	Sex           Sex    `json:"sex" form:"sex" gorm:"column:sex;comment:性别 1:男，2：女"`
	Country       string `json:"country" form:"country" gorm:"column:country;comment:国家"`
	Province      string `json:"provice" form:"provice" gorm:"column:provice;comment:省份"`
	City          string `json:"city" form:"city" gorm:"column:city;comment:城市"`
	IsDefault     int    `json:"isDefault" form:"isDefault" gorm:"column:is_default;comment:是否默认"`
	DetailAddress string `json:"detailAddress" form:"detailAddress" gorm:"column:detail_address;comment:详细地址"`

	SysUserID int `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:关联标记"`
}
