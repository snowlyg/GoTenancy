package models

import (
	"github.com/jinzhu/gorm"
)

type Tenant struct {
	gorm.Model
	UId      uint32 `json:"uid"  gorm:"unique;not null;size:255"`
	Name     string `json:"name" validate:"required,gte=6,lte=50"  comment:"名称" form:"name" gorm:"not null;size:50"`
	FullName string `json:"full_name" validate:"required,gte=6,lte=255" comment:"全称" form:"full_name" gorm:"size:255"`
	Email    string `json:"email" validate:"email" comment:"邮箱" form:"email" gorm:"unique" `
	Telphone string `json:"telphone" form:"telphone" gorm:"unique;size:11"`
	Rmk      string `json:"rmk" form:"rmk" gorm:"type(text)"`
}
