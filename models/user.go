package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"gte=2,lte=50"  comment:"姓名" form:"name" gorm:"size:255"`
	Username string `json:"username" validate:"required,gte=2,lte=50" comment:"用户名" form:"username" gorm:"unique;not null;size:255"`
	Email    string `json:"email" form:"email" gorm:"unique"`
	Telphone string `json:"telphone" form:"telphone" gorm:"unique;size:13"`
	Password []byte `json:"password" validate:"required"  comment:"密码" form:"-" gorm:"not null"`
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
