package blogs

import (
	"GoTenancy/models/users"
	"github.com/jinzhu/gorm"
	"github.com/qor/publish2"
)

type Article struct {
	gorm.Model
	Author   users.User
	AuthorID uint
	Title    string
	Content  string `gorm:"type:text"`
	publish2.Version
	publish2.Schedule
	publish2.Visible
}
