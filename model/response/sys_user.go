package response

import (
	"github.com/snowlyg/go-tenancy/model"
)

type SysUserResponse struct {
	User model.SysUser `json:"user"`
}

type LoginResponse struct {
	Token string `json:"AccessToken"`
}
