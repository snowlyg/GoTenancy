package models

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	Name        string       `json:"name" validate:"required,gte=6,lte=50"  comment:"标识名称" form:"name" gorm:"unique;not null;size:50"`
	DisplayName string       `json:"display_name" validate:"required,gte=6,lte=255" comment:"显示名称" form:"display_name" gorm:"size:255"`
	Rmk         string       `json:"rmk" form:"rmk" gorm:"type(text)"`
	IsAdmin     sql.NullBool `json:"is_admin" gorm:"not null;default:0"`
	PermIds     []string     `json:"perm_ids" form:"perm_ids" gorm:"-"`
}
