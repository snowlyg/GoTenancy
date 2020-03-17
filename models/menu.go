package models

import (
	"github.com/jinzhu/gorm"
)

type Menu struct {
	gorm.Model

	Title  string `json:"title"`
	Href   string `json:"href"`
	Icon   string `json:"icon"`
	Target string `json:"target"`
	Child  []*Menu
}
