package model

import "time"

type SysRegion struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Code      int        `json:"code" gorm:""`
	PCode     int        `json:"p_code" gorm:""`
	Name      string     `json:"name" gorm:""`
}
