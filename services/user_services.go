package services

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/snowlyg/go-tenancy/models"
)

type UserService interface {
	GetAll() []models.User
	GetByID(id int64) (models.User, bool)
	GetByUsernameAndPassword(username, userPassword string) (*models.User, bool)
	DeleteByID(id int64) bool

	Update(id int64, user models.User) (models.User, error)
	UpdatePassword(id int64, newPassword string) (models.User, error)
	UpdateUsername(id int64, newUsername string) (models.User, error)

	Create(userPassword string, user models.User) (models.User, error)
}

func NewUserService(gdb *gorm.DB) UserService {
	return &userService{
		gdb: gdb,
	}
}

type userService struct {
	gdb *gorm.DB
}

func (s *userService) GetAll() []models.User {
	//return s.repo.SelectMany(func(_ models.User) bool {
	//	return true
	//}, -1)
	return nil
}

func (s *userService) GetByID(id int64) (models.User, bool) {
	return models.User{}, true
}

func (s *userService) GetByUsernameAndPassword(username, password string) (*models.User, bool) {
	user := &models.User{Username: username, Password: []byte(password)}
	if notFound := s.gdb.Find(user).RecordNotFound(); notFound {
		return nil, false
	}
	return user, true
}

func (s *userService) Update(id int64, user models.User) (models.User, error) {
	return models.User{}, nil
}

func (s *userService) UpdatePassword(id int64, newPassword string) (models.User, error) {

	hashed, err := models.GeneratePassword(newPassword)
	if err != nil {
		return models.User{}, err
	}

	return s.Update(id, models.User{
		Password: hashed,
	})
}

func (s *userService) UpdateUsername(id int64, newUsername string) (models.User, error) {
	return s.Update(id, models.User{
		Username: newUsername,
	})
}

func (s *userService) Create(userPassword string, user models.User) (models.User, error) {
	if user.ID > 0 || userPassword == "" || user.Firstname == "" || user.Username == "" {
		return models.User{}, errors.New("unable to create this user")
	}

	hashed, err := models.GeneratePassword(userPassword)
	if err != nil {
		return models.User{}, err
	}
	user.Password = hashed

	return models.User{}, nil
}

func (s *userService) DeleteByID(id int64) bool {
	return true
}
