package response

import "github.com/snowlyg/go-tenancy/model"

type UserLabelWithUserId struct {
	model.UserLabel
	SysUserID uint `json:"sysUserId"`
}
