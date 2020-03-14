package services

import (
	"errors"

	"github.com/snowlyg/go-tenancy/models"
	"github.com/snowlyg/go-tenancy/repositories"
)

type UserService interface {
	GetAll() []models.User
	GetByID(id int64) (models.User, bool)
	GetByUsernameAndPassword(username, userPassword string) (models.User, bool)
	DeleteByID(id int64) bool

	Update(id int64, user models.User) (models.User, error)
	UpdatePassword(id int64, newPassword string) (models.User, error)
	UpdateUsername(id int64, newUsername string) (models.User, error)

	Create(userPassword string, user models.User) (models.User, error)
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

type userService struct {
	repo repositories.UserRepository
}

func (s *userService) GetAll() []models.User {
	return s.repo.SelectMany(func(_ models.User) bool {
		return true
	}, -1)
}

func (s *userService) GetByID(id int64) (models.User, bool) {
	return s.repo.Select(func(m models.User) bool {
		return m.ID == id
	})
}

func (s *userService) GetByUsernameAndPassword(username, userPassword string) (models.User, bool) {
	if username == "" || userPassword == "" {
		return models.User{}, false
	}

	return s.repo.Select(func(m models.User) bool {
		if m.Username == username {
			hashed := m.HashedPassword
			if ok, _ := models.ValidatePassword(userPassword, hashed); ok {
				return true
			}
		}
		return false
	})
}

func (s *userService) Update(id int64, user models.User) (models.User, error) {
	user.ID = id
	return s.repo.InsertOrUpdate(user)
}

func (s *userService) UpdatePassword(id int64, newPassword string) (models.User, error) {

	hashed, err := models.GeneratePassword(newPassword)
	if err != nil {
		return models.User{}, err
	}

	return s.Update(id, models.User{
		HashedPassword: hashed,
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
	user.HashedPassword = hashed

	return s.repo.InsertOrUpdate(user)
}

func (s *userService) DeleteByID(id int64) bool {
	return s.repo.Delete(func(m models.User) bool {
		return m.ID == id
	}, 1)
}
