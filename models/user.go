package models

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name     string       `json:"name" validate:"gte=6,lte=50"  comment:"姓名" form:"name" gorm:"size:50"`
	Username string       `json:"username" validate:"required,gte=6,lte=12" comment:"用户名" form:"username" gorm:"unique;not null;size:50"`
	Email    string       `json:"email" form:"email" gorm:"unique" validate:"email"`
	Telphone string       `json:"telphone" form:"telphone" gorm:"unique;size:11"`
	IsAdmin  sql.NullBool `json:"is_admin" gorm:"not null;default:0"`
	Password []byte       `json:"password" validate:"gte=6"  comment:"密码" form:"-" gorm:"not null"`
}

func (u User) IsValid() bool {
	return u.ID > 0
}

func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

func ValidatePassword(userPassword string, hashed []byte) (bool, error) {

	if err := bcrypt.CompareHashAndPassword(hashed, []byte(userPassword)); err != nil {
		return false, err
	}

	return true, nil
}
