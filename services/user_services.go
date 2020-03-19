package services

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/snowlyg/go-tenancy/common"
	"github.com/snowlyg/go-tenancy/models"
)

var IsAdmin = map[string]interface{}{"is_admin": 0}

type UserService interface {
	GetAll(args map[string]interface{}, pagination *common.Pagination, ispreload bool) (int64, []*models.User)
	GetByID(id uint) (models.User, bool)
	GetByUsernameAndPassword(username, userPassword string) (*models.User, bool)
	DeleteByID(id uint) error

	Update(id uint, user *models.User) error
	UpdatePassword(id uint, newPassword string) error
	UpdateUsername(id uint, newUsername string) error

	Create(userPassword string, user *models.User) error
}

func NewUserService(gdb *gorm.DB) UserService {
	return &userService{
		gdb: gdb,
	}
}

type userService struct {
	gdb *gorm.DB
}

func (s *userService) GetAll(args map[string]interface{}, pagination *common.Pagination, ispreload bool) (int64, []*models.User) {
	var users []*models.User
	var count int64

	args["is_admin"] = 0
	db := s.gdb.Where(args)
	if ispreload {
		//db = db.Preload("Child")
	}

	db.Find(&users).Count(&count)

	if pagination != nil {
		db = db.Limit(pagination.Limit).Offset(pagination.Limit * (pagination.Page - 1))
	}

	if err := db.Find(&users).Error; err != nil {
		panic(err)
	}
	return count, users
}

func (s *userService) GetByID(id uint) (models.User, bool) {
	user := models.User{Model: gorm.Model{ID: id}}
	if notFound := s.gdb.Find(&user).RecordNotFound(); notFound {
		return user, false
	}
	return user, true
}

func (s *userService) GetByUsernameAndPassword(username, password string) (*models.User, bool) {
	user := &models.User{Username: username, Password: []byte(password)}
	if notFound := s.gdb.Find(user).RecordNotFound(); notFound {
		return nil, false
	}
	return user, true
}

func (s *userService) Update(id uint, user *models.User) error {
	if err := s.gdb.Where("id = ?", id).Where(IsAdmin).Update(user).Error; err != nil {
		return err
	}
	return nil
}

func (s *userService) UpdatePassword(id uint, newPassword string) error {

	hashed, err := models.GeneratePassword(newPassword)
	if err != nil {
		return err
	}

	return s.Update(id, &models.User{
		Password: hashed,
	})
}

func (s *userService) UpdateUsername(id uint, newUsername string) error {
	return s.Update(id, &models.User{
		Username: newUsername,
	})
}

func (s *userService) Create(userPassword string, user *models.User) error {
	var (
		hashed []byte
		err    error
	)
	if user.ID > 0 || userPassword == "" || user.Name == "" || user.Username == "" {
		return errors.New("unable to create this user")
	}

	hashed, err = models.GeneratePassword(userPassword)
	user.Password = hashed

	err = s.gdb.Create(user).Error

	if err != nil {
		return err
	}

	return nil
}

func (s *userService) DeleteByID(id uint) error {
	if err := s.gdb.Where(IsAdmin).Delete(models.User{Model: gorm.Model{ID: id}}).Error; err != nil {
		return err
	}
	return nil
}
