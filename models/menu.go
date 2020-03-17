package models

import (
	"github.com/jinzhu/gorm"
)

type Menu struct {
	gorm.Model

	Title    string  `json:"title"`
	Href     string  `json:"href"`
	Icon     string  `json:"icon"`
	Target   string  `json:"target"`
	ParentId uint    `json:"parent_id" gorm:"default:0"`
	Child    []*Menu `json:"child" gorm:"foreignkey:ParentId"`
}
