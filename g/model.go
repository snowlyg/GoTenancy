package g

import (
	"time"

	"gorm.io/gorm"
)

type TENANCY_MODEL struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
