package models

import (
	"github.com/jinzhu/gorm"
)

type Perm struct {
	gorm.Model

	OrderNumber int64   `json:"order_number"`
	Authority   string  `json:"authority"`
	Checked     int8    `json:"checked"`
	IsMenu      int8    `json:"is_menu"`
	Title       string  `json:"title"`
	Href        string  `json:"href"`
	Icon        string  `json:"icon"`
	Target      string  `json:"target"`
	ParentId    int64   `json:"parent_id" gorm:"default:0"`
	Child       []*Perm `json:"child" gorm:"foreignkey:ParentId"`
}
