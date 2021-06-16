package model

import (
	"github.com/snowlyg/go-tenancy/g"
)

// GeneralAddress 用户收货地址
type GeneralAddress struct {
	g.TENANCY_MODEL
	Name      string `json:"name" gorm:"type:varchar(32);not null;comment:姓名"`
	Phone     string `json:"phone" gorm:"type:varchar(16);not null;comment:手机号"`
	Sex       int    `json:"sex" form:"sex" gorm:"not null;column:sex;comment:性别 0:女 1:男，2：未知"`
	Country   string `json:"country" form:"country" gorm:"type:varchar(64);not null;column:country;comment:国家"`
	Province  string `json:"province" form:"province" gorm:"type:varchar(64);not null;column:province;comment:省份"`
	City      string `json:"city" form:"city" gorm:"type:varchar(64);not null;column:city;comment:城市"`
	District  string `json:"district" form:"district" gorm:"type:varchar(64);not null;column:district;comment:地区"`
	IsDefault bool   `json:"isDefault" form:"isDefault" gorm:"not null;type:bool;column:is_default;comment:是否默认"`
	Detail    string `json:"detail" form:"detail" gorm:"type:varchar(254);not null;column:detail;comment:详细地址"`
	Postcode  string `json:"postcode" form:"postcode" gorm:"type:varchar(20);not null;column:postcode;comment:邮政编码"`

	// 可选字段
	Age          int    `json:"age" form:"age" gorm:"column:age;comment:年龄"`
	HospitalName string `json:"hospitalName" form:"hospitalName" gorm:"type:varchar(50);column:hospital_name;comment:邮政编码"`
	LocName      string `json:"locName" form:"locName" gorm:"type:varchar(50);column:loc_name;comment:邮政编码"`
	BedNum       string `json:"bedNum" form:"bedNum" gorm:"type:varchar(10);column:bed_num;comment:邮政编码"`
	HospitalNO   string `json:"hospitalNo" form:"hospitalNo" gorm:"type:varchar(20);column:hospital_no;comment:住院号"`
	Disease      string `json:"disease" form:"disease" gorm:"type:varchar(150);column:disease;comment:病种"`

	SysUserID uint `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:关联标记"`
}
