package services

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/snowlyg/go-tenancy/models"
)

type UserService interface {
	GetAll(args map[string]interface{}, ispreload bool) []*models.User
	GetByID(id int64) (models.User, bool)
	GetByUsernameAndPassword(username, userPassword string) (*models.User, bool)
	DeleteByID(id int64) bool

	Update(id int64, user *models.User) error
	UpdatePassword(id int64, newPassword string) error
	UpdateUsername(id int64, newUsername string) error

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

func (s *userService) GetAll(args map[string]interface{}, ispreload bool) []*models.User {
	var users []*models.User
	db := s.gdb
	if len(args) > 0 {
		for parm, arg := range args {
			db = db.Where(parm, arg)
		}
	}

	if ispreload {
		//db = db.Preload("Child")
	}

	if err := db.Find(&users).Error; err != nil {
		panic(err)
	}
	return users
}

func (s *userService) GetByID(id int64) (models.User, bool) {
	return models.User{}, true
}

func (s *userService) GetByUsernameAndPassword(username, password string) (*models.User, bool) {
	user := &models.User{Username: username, Password: []byte(password)}
	if notFound := s.gdb.Find(user).RecordNotFound(); notFound {
		return user, false
	}
	return user, true
}

func (s *userService) Update(id int64, user *models.User) error {
	return nil
}

func (s *userService) UpdatePassword(id int64, newPassword string) error {

	hashed, err := models.GeneratePassword(newPassword)
	if err != nil {
		return err
	}

	return s.Update(id, &models.User{
		Password: hashed,
	})
}

func (s *userService) UpdateUsername(id int64, newUsername string) error {
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

func (s *userService) DeleteByID(id int64) bool {
	return true
}
