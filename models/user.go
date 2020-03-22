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
	Email    string       `json:"email" validate:"email" comment:"邮箱" form:"email" gorm:"unique" `
	Telphone string       `json:"telphone" form:"telphone" gorm:"unique;size:11"`
	IsAdmin  sql.NullBool `json:"is_admin" gorm:"not null;default:0"`
	Password string       `json:"password" validate:"required,gte=6 ,lte=14"  comment:"密码" form:"-" gorm:"not null"`
	TenantId uint         `json:"tenant_id" form:"tenant_id"`
	RoleIds  []uint       `json:"role_ids" form:"role_ids" validate:"required" gorm:"-"`
}

func (u User) IsValid() bool {
	return u.ID > 0
}

func GeneratePassword(userPassword string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	return string(password), err
}

func ValidatePassword(userPassword string, hashed string) (bool, error) {

	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, err
	}

	return true, nil
}
