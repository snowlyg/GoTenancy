package response

import (
	"github.com/snowlyg/go-tenancy/model"
)

type GeneralUser struct {
	TenancyResponse
	Uid           uint              `json:"uid"`
	Username      string            `json:"userName"`
	AuthorityName string            `json:"authorityName"`
	AuthorityType int               `json:"authorityType"`
	AuthorityId   string            `json:"authorityId"`
	GroupName     string            `json:"groupName"`
	LabelID       uint              `json:"labelId"` // 用户标签 id
	Label         []model.UserLabel `gorm:"-" json:"label"`

	model.BaseGeneralInfo
}

type GeneralUserDetail struct {
	TenancyResponse
	AvatarUrl     string  `json:"avatarUrl"`
	NickName      string  `json:"nickName"`
	NowMoney      float64 `json:"nowMoney"`
	PayCount      int     `json:"payCount"`
	PayPrice      float64 `json:"payPrice"`
	TotalPayCount int     `json:"totalPayCount"`
	TotalPayPrice float64 `json:"totalPayPrice"`
}
