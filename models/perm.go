package models

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

type Perm struct {
	gorm.Model

	OrderNumber int64         `json:"order_number"`
	Checked     int8          `json:"checked"`
	Type        int8          `json:"type"` // 1：通用，2：admin ,3:tenant
	IsMenu      int8          `json:"is_menu"`
	Title       string        `json:"title" gorm:"not null;size:255"`
	Href        string        `json:"href" gorm:"size:255"`
	Icon        string        `json:"icon"`
	Target      string        `json:"target"`
	Method      string        `json:"method"`
	ParentId    sql.NullInt64 `json:"parent_id" gorm:"default:0"`
	Child       []*Perm       `json:"child" gorm:"foreignkey:ParentId"`
}
