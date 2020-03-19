package services

import (
	"errors"
	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/jinzhu/gorm"
	"github.com/snowlyg/go-tenancy/common"
	"github.com/snowlyg/go-tenancy/models"
)

var NotAdmin = map[string]interface{}{"is_admin": 0}
var IsAdmin = map[string]interface{}{"is_admin": 1}

type UserService interface {
	GetAll(args map[string]interface{}, pagination *common.Pagination, ispreload bool) (int64, []*models.User)
	GetByID(id uint) (models.User, bool)
	GetByUsername(username string) (*models.User, bool)
	DeleteByID(id uint) error
	DeleteMnutil(userIds []common.Id) error

	Update(id uint, user *models.User) error
	UpdatePassword(id uint, newPassword string) error
	UpdateUsername(id uint, newUsername string) error

	Create(userPassword string, user *models.User, roleIds []uint) error
}

func NewUserService(gdb *gorm.DB, ce *casbin.Enforcer) UserService {
	return &userService{
		gdb: gdb,
		ce:  ce,
	}
}

type userService struct {
	gdb *gorm.DB
	ce  *casbin.Enforcer
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

func (s *userService) GetByUsername(username string) (*models.User, bool) {
	user := &models.User{}
	if notFound := s.gdb.Where("username = ?", username).Find(&user).RecordNotFound(); notFound {
		return nil, false
	}
	return user, true
}

func (s *userService) Update(id uint, user *models.User) error {
	if err := s.gdb.Where("id = ?", id).Where(NotAdmin).Update(user).Error; err != nil {
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

func (s *userService) Create(userPassword string, user *models.User, roleIds []uint) error {
	return s.gdb.Transaction(func(tx *gorm.DB) error {

		if user.ID > 0 || userPassword == "" || user.Name == "" || user.Username == "" {
			return errors.New("unable to create this user")
		}

		hashed, err := models.GeneratePassword(userPassword)
		user.Password = hashed
		if err != nil {
			return err
		}

		if err = tx.Create(user).Error; err != nil {
			return err
		}

		if err = s.addRoles(roleIds, user); err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})

}

func (s *userService) DeleteByID(id uint) error {
	if err := s.gdb.Where(NotAdmin).Delete(models.User{Model: gorm.Model{ID: id}}).Error; err != nil {
		return err
	}
	return nil
}

func (s *userService) DeleteMnutil(userIds []common.Id) error {
	return s.gdb.Transaction(func(tx *gorm.DB) error {
		for _, userid := range userIds {
			if err := tx.Where(NotAdmin).Delete(models.User{Model: gorm.Model{ID: uint(userid.Id)}}).Error; err != nil {
				return err
			}
		}

		// 返回 nil 提交事务
		return nil
	})
}

func (s *userService) addRoles(roleids []uint, user *models.User) error {
	if len(roleids) > 0 {
		userId := strconv.FormatUint(uint64(user.ID), 10)
		if _, err := s.ce.DeleteRolesForUser(userId); err != nil {
			return err
		}

		for _, roleId := range roleids {
			roleId := strconv.FormatUint(uint64(roleId), 10)
			if _, err := s.ce.AddRoleForUser(userId, roleId); err != nil {
				return err
			}
		}
	}

	return nil
}
