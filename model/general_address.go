package model

import (
	"github.com/snowlyg/go-tenancy/g"
)

// GeneralAddress 用户收货地址
type GeneralAddress struct {
	g.TENANCY_MODEL
	Name      string `json:"name" gorm:"not null;default:'姓名';comment:姓名"`
	Phone     string `json:"phone" gorm:"not null;default:'';comment:手机号"`
	Sex       Sex    `json:"sex" form:"sex" gorm:"not null;column:sex;comment:性别 0:女 1:男，2：未知"`
	Country   string `json:"country" form:"country" gorm:"not null;column:country;comment:国家"`
	Province  string `json:"province" form:"province" gorm:"not null;column:province;comment:省份"`
	City      string `json:"city" form:"city" gorm:"not null;column:city;comment:城市"`
	District  string `json:"district" form:"district" gorm:"not null;column:district;comment:地区"`
	IsDefault bool   `json:"isDefault" form:"isDefault" gorm:"not null;type:bool;column:is_default;comment:是否默认"`
	Detail    string `json:"detail" form:"detail" gorm:"column:not null;detail;comment:详细地址"`
	Postcode  string `json:"postcode" form:"postcode" gorm:"not null;column:postcode;comment:邮政编码"`

	// 可选字段
	Age          int    `json:"age" form:"age" gorm:"column:age;comment:年龄"`
	HospitalName string `json:"hospitalName" form:"hospitalName" gorm:"column:hospital_name;comment:邮政编码"`
	LocName      string `json:"locName" form:"locName" gorm:"column:loc_name;comment:邮政编码"`
	BedNum       string `json:"bedNum" form:"bedNum" gorm:"column:bed_num;comment:邮政编码"`
	HospitalNO   string `json:"hospitalNo" form:"hospitalNo" gorm:"column:hospital_no;comment:住院号"`
	Disease      string `json:"disease" form:"disease" gorm:"column:disease;comment:病种"`

	SysUserID int `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:关联标记"`
}
