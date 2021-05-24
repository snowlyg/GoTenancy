package response

import (
	"time"

	"github.com/snowlyg/go-tenancy/model"
)

type LoginResponse struct {
	User  interface{} `json:"user"`
	Token string      `json:"AccessToken"`
}

type SysAdminUser struct {
	TenancyResponse
	Username      string `json:"userName"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	NickName      string `json:"nickName"`
	HeaderImg     string `json:"headerImg"`
	AuthorityName string `json:"authorityName"`
	AuthorityType int    `json:"authorityType"`
	AuthorityId   string `json:"authorityId"`
	DefaultRouter string `json:"defaultRouter"`
}

type SysTenancyUser struct {
	TenancyResponse
	Username      string `json:"userName"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	NickName      string `json:"nickName"`
	HeaderImg     string `json:"headerImg"`
	AuthorityName string `json:"authorityName"`
	AuthorityType int    `json:"authorityType"`
	AuthorityId   string `json:"authorityId"`
	TenancyName   string `json:"tenancyName"`
	DefaultRouter string `json:"defaultRouter"`
}

type SysGeneralUser struct {
	TenancyResponse
	Username      string    `json:"userName"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	NickName      string    `json:"nickName"`
	AuthorityName string    `json:"authorityName"`
	AuthorityType int       `json:"authorityType"`
	AuthorityId   string    `json:"authorityId"`
	AvatarUrl     string    `json:"avatarUrl"`
	Sex           model.Sex `json:"sex"`
	Subscribe     int       `json:"subscribe"`
	OpenId        string    `json:"openId"`
	UnionId       string    `json:"unionId"`
	Country       string    `json:"country"`
	Province      string    `json:"province"`
	City          string    `json:"city"`
	IdCard        string    `json:"idCard"`
	IsAuth        int       `json:"isAuth"`
	RealName      string    `json:"realName"`
	Birthday      time.Time `json:"birthday"`
}
