package services

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/snowlyg/go-tenancy/common"
	"github.com/snowlyg/go-tenancy/models"
)

type RoleService interface {
	GetAll(args map[string]interface{}, pagination *common.Pagination, ispreload bool) (int64, []*models.Role)
	GetByID(id uint) (models.Role, bool)
	DeleteByID(id uint) error
	DeleteMnutil(ids []common.Id) error
	Update(id uint, role *models.Role) error
	Create(role *models.Role) error
}

func NewRoleService(gdb *gorm.DB) RoleService {
	return &roleService{
		gdb: gdb,
	}
}

type roleService struct {
	gdb *gorm.DB
}

//GetAll 查询所有数据
func (s *roleService) GetAll(args map[string]interface{}, pagination *common.Pagination, ispreload bool) (int64, []*models.Role) {
	var roles []*models.Role
	var count int64

	args["is_admin"] = 0
	db := s.gdb.Where(args)

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
	return models.Role{}, true
}

func (s *roleService) Update(id uint, role *models.Role) error {
	return nil
}

func (s *roleService) Create(role *models.Role) error {
	var (
		err error
	)
	if role.ID > 0 {
		return errors.New("unable to create this role")
	}

	err = s.gdb.Create(role).Error

	if err != nil {
		return err
	}

	return nil
}

func (s *roleService) DeleteByID(id uint) error {
	if err := s.gdb.Where(IsAdmin).Delete(models.Role{Model: gorm.Model{ID: id}}).Error; err != nil {
		return err
	}
	return nil
}

func (s *roleService) DeleteMnutil(ids []common.Id) error {
	return s.gdb.Transaction(func(tx *gorm.DB) error {
		for _, id := range ids {
			if err := tx.Where(IsAdmin).Delete(models.Role{Model: gorm.Model{ID: uint(id.Id)}}).Error; err != nil {
				return err
			}
		}

		// 返回 nil 提交事务
		return nil
	})
}
