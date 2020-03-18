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
	Title       string  `json:"title" gorm:"not null;size:255"`
	Href        string  `json:"href" gorm:"size:255"`
	Icon        string  `json:"icon"`
	Target      string  `json:"target" gorm:"default(_self)"`
	ParentId    int64   `json:"parent_id" gorm:"default(0))"`
	Child       []*Perm `json:"child" gorm:"foreignkey:ParentId"`
}
