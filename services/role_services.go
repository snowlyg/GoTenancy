package services

import (
	"errors"
	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/jinzhu/gorm"
	"github.com/snowlyg/go-tenancy/common"
	"github.com/snowlyg/go-tenancy/config"
	"github.com/snowlyg/go-tenancy/models"
)

type RoleService interface {
	GetAll(args map[string]interface{}, pagination *common.Pagination, ispreload bool) (int64, []*models.Role)
	GetByID(id uint) (models.Role, bool)
	DeleteByID(id uint) error
	DeleteMnutil(ids []common.Id) error
	Update(id uint, role *models.Role) error
	Create(role *models.Role, permIds []uint) error
	GetAdmin() (models.Role, bool)

	GetPermsByID(id uint) ([]models.Perm, error)
}

func NewRoleService(gdb *gorm.DB, ce *casbin.Enforcer) RoleService {
	return &roleService{
		gdb: gdb,
		ce:  ce,
	}
}

type roleService struct {
	gdb *gorm.DB
	ce  *casbin.Enforcer
}

//GetAll 查询所有数据
func (s *roleService) GetAll(args map[string]interface{}, pagination *common.Pagination, ispreload bool) (int64, []*models.Role) {
	var roles []*models.Role
	var count int64

	args["is_admin"] = 0
	db := s.gdb.Where(args).Order("id desc")

	if ispreload {
		//db = db.Preload("Child")
	}

	db.Find(&roles).Count(&count)

	if pagination != nil {
		db = db.Limit(pagination.Limit).Offset(pagination.Limit * (pagination.Page - 1))
	}

	if err := db.Find(&roles).Error; err != nil {
		panic(err)
	}

	return count, roles
}

func (s *roleService) GetByID(id uint) (models.Role, bool) {
	user := models.Role{Model: gorm.Model{ID: id}}
	if notFound := s.gdb.Find(&user).RecordNotFound(); notFound {
		return user, false
	}
	return user, true
}

func (s *roleService) GetPermsByID(id uint) ([]models.Perm, error) {
	permIds := s.ce.GetPermissionsForUser(strconv.FormatUint(uint64(id), 10))
	var perms []models.Perm
	if err := s.gdb.Where("id in (?)", permIds).Find(&perms).Error; err != nil {
		return nil, err
	}

	return perms, nil
}

func (s *roleService) Update(id uint, role *models.Role) error {
	if err := s.gdb.Where("id = ?", id).Where(NotAdmin).Update(role).Error; err != nil {
		return err
	}
	return nil
}

func (s *roleService) Create(role *models.Role, permIds []uint) error {
	if config.Config.DB.Adapter != "mysql" {
		var err error

		if role.ID > 0 {
			return errors.New("unable to create this role")
		}

		if err = s.gdb.Create(role).Error; err != nil {
			return err
		}

		if err = s.addPerms(permIds, role); err != nil {
			return err
		}

		return nil
	} else {
		return s.gdb.Transaction(func(tx *gorm.DB) error {
			var err error

			if role.ID > 0 {
				return errors.New("unable to create this role")
			}

			if err = tx.Create(role).Error; err != nil {
				return err
			}

			if err = s.addPerms(permIds, role); err != nil {
				return err
			}

			// 返回 nil 提交事务
			return nil
		})
	}

}

func (s *roleService) DeleteByID(id uint) error {
	if err := s.gdb.Where(NotAdmin).Delete(models.Role{Model: gorm.Model{ID: id}}).Error; err != nil {
		return err
	}
	return nil
}

func (s *roleService) DeleteMnutil(ids []common.Id) error {
	return s.gdb.Transaction(func(tx *gorm.DB) error {
		for _, id := range ids {
			if err := tx.Where(NotAdmin).Delete(models.Role{Model: gorm.Model{ID: uint(id.Id)}}).Error; err != nil {
				return err
			}
		}

		// 返回 nil 提交事务
		return nil
	})
}

func (s *roleService) GetAdmin() (models.Role, bool) {
	role := models.Role{}
	if notFound := s.gdb.Where(IsAdmin).Find(&role).RecordNotFound(); notFound {
		return role, false
	}
	return role, true
}

func (s *roleService) addPerms(permIds []uint, role *models.Role) error {
	if len(permIds) > 0 {
		roleId := strconv.FormatUint(uint64(role.ID), 10)
		if _, err := s.ce.DeletePermissionsForUser(roleId); err != nil {
			return err
		}
		var perms []models.Perm
		s.gdb.Where("id in (?)", permIds).Find(&perms)
		for _, perm := range perms {
			if _, err := s.ce.AddPolicy(roleId, perm.Href, perm.Method); err != nil {
				return err
			}
		}
	}

	return nil
}
