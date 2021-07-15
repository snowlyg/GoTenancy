package model

import (
	"time"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/multi"
)

type SysUser struct {
	g.TENANCY_MODEL

	Username    string       `json:"userName" gorm:"not null;type:varchar(32);comment:用户登录名"`
	Password    string       `json:"-"  gorm:"not null;type:varchar(128);comment:用户登录密码"`
	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	AuthorityId string       `json:"authorityId" gorm:"not null;comment:用户角色ID"`

	Status int `gorm:"column:status;type:tinyint(1);not null;default:1" json:"status"` // 1为正常，2为禁止

	AdminInfo   AdminInfo   `json:"adminInfo" gorm:"foreignKey:SysUserID;references:ID;comment:管理员信息"`
	TenancyInfo TenancyInfo `json:"tenancyInfo" gorm:"foreignKey:SysUserID;references:ID;comment:商户信息"`
	GeneralInfo GeneralInfo `json:"generalInfo" gorm:"foreignKey:SysUserID;references:ID;comment:普通用户信息"`
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

type AdminInfo struct {
	g.TENANCY_MODEL
	Email     string `json:"email" gorm:"default:'';comment:员工邮箱"`
	Phone     string `json:"phone" gorm:"type:char(15);default:'';comment:员工手机号" `
	NickName  string `json:"nickName" gorm:"type:varchar(16);default:'员工姓名';comment:员工姓名" `
	HeaderImg string `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`

	SysUserID uint `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:关联标记"`
}

type TenancyInfo struct {
	g.TENANCY_MODEL
	Email     string `json:"email" gorm:"default:'';comment:员工邮箱"`
	Phone     string `json:"phone" gorm:"type:char(15);default:'';comment:员工手机号" `
	NickName  string `json:"nickName" gorm:"type:varchar(16);default:'员工姓名';comment:员工姓名" `
	HeaderImg string `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`

	SysUserID    uint `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:关联标记"`
	SysTenancyID uint `json:"sysTenancyId" form:"sysTenancyId" gorm:"column:sys_tenancy_id;comment:关联标记"`
}

const (
	Unknown int = iota
	Male
	Female
)

type GeneralInfo struct {
	g.TENANCY_MODEL
	BaseGeneralInfo
	LabelID   uint `gorm:"column:label_id;type:varchar(64)" json:"labelId"`                     // 用户标签 id
	GroupID   uint `gorm:"column:group_id;type:int unsigned;not null;default:0" json:"groupId"` // 用户分组id
	SysUserID uint `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:关联标记"`
}

type BaseGeneralInfo struct {
	Email     string    `json:"email" gorm:"default:'';comment:员工邮箱"`
	Phone     string    `json:"phone" gorm:"type:char(15);default:'';comment:手机号"`
	NickName  string    `json:"nickName" gorm:"type:varchar(16);comment:昵称"`
	AvatarUrl string    `json:"avatarUrl" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	Sex       int       `json:"sex" form:"sex" gorm:"column:sex;comment:性别 1:男，2：女"`
	Subscribe int       `json:"subscribe" form:"subscribe" gorm:"column:subscribe;comment:是否订阅"`
	OpenId    string    `json:"openId" form:"openId" gorm:"type:varchar(30);column:open_id;comment:openid"`
	UnionId   string    `json:"unionId" form:"unionId" gorm:"type:varchar(30);column:union_id;comment:unionId"`
	Country   string    `json:"country" form:"country" gorm:"type:varchar(32);column:country;comment:国家"`
	Province  string    `json:"province" form:"province" gorm:"type:varchar(32);column:province;comment:省份"`
	City      string    `json:"city" form:"city" gorm:"type:varchar(32);column:city;comment:城市"`
	IdCard    string    `json:"idCard" form:"idCard" gorm:"type:varchar(20);column:id_card;comment:身份证号"`
	IsAuth    int       `json:"isAuth" form:"isAuth" gorm:"column:is_auth;comment:是否实名认证"`
	RealName  string    `json:"realName" form:"realName" gorm:"type:varchar(64);column:real_name;comment:真实IP"`
	Birthday  time.Time `json:"birthday" form:"birthday" gorm:"column:birthday;comment:生日"`

	Mark     string    `gorm:"column:mark;type:varchar(255);not null;default:''" json:"mark"`                      // 用户备注
	Addres   string    `gorm:"column:addres;type:varchar(128)" json:"addres"`                                      // 地址
	LastTime time.Time `gorm:"column:last_time;type:timestamp" json:"lastTime"`                                    // 最后一次登录时间
	LastIP   string    `gorm:"column:last_ip;type:varchar(16);not null" json:"lastIp"`                             // 最后一次登录ip
	NowMoney float64   `gorm:"column:now_money;type:decimal(8,2) unsigned;not null;default:0.00" json:"nowMoney"`  // 用户余额
	UserType string    `gorm:"column:user_type;type:varchar(32);not null" json:"userType"`                         // 用户类型 h5,小程序 routine ,微信 wechat
	MainUId  uint      `gorm:"index:main_uid;column:main_uid;type:int unsigned;default:0" json:"mainUid"`          // 主账号
	PayCount uint      `gorm:"column:pay_count;type:int unsigned;not null;default:0" json:"payCount"`              // 用户购买次数
	PayPrice float64   `gorm:"column:pay_price;type:decimal(10,2) unsigned;not null;default:0.00" json:"payPrice"` // 用户消费金额
}
