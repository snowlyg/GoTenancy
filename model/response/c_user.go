package response

import (
	"time"

	"github.com/snowlyg/go-tenancy/model"
)

type GeneralUser struct {
	TenancyResponse
	Uid           uint      `json:"uid"`
	Username      string    `json:"userName"`
	AuthorityName string    `json:"authorityName"`
	AuthorityType int       `json:"authorityType"`
	AuthorityId   string    `json:"authorityId"`
	FirstPayTime  time.Time `json:"firstPayTime"`
	LastPayTime   time.Time `json:"lastPayTime"`
	GroupName     string    `json:"groupName"`
	LabelId       string    `json:"labelId"`
	Label         []string  `gorm:"-" json:"label"`

	model.BaseGeneralInfo
}

type GeneralUserDetail struct {
	TenancyResponse
	Uid           uint           `json:"uid"`
	AvatarUrl     string         `json:"avatarUrl"`
	NickName      string         `json:"nickName"`
	NowMoney      float64        `json:"nowMoney"`
	PayCount      int            `json:"payCount"`
	PayPrice      float64        `json:"payPrice"`
	TotalPayCount int            `json:"totalPayCount"`
	TotalPayPrice float64        `json:"totalPayPrice"`
	IdCard        string         `json:"idCard"`
	RealName      string         `json:"realName"`
	Birthday      model.Birthday `json:"birthday"`
	Mark          string         `json:"mark"`
	Address       string         `json:"address"`
	Phone         string         `json:"phone"`
	GroupId       uint           `json:"groupId"`
	LabelIds      []uint         `gorm:"-"  json:"labelId"`
}
