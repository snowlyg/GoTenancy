package services

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/snowlyg/go-tenancy/models"
)

type RoleService interface {
	GetAll(args map[string]interface{}, ispreload bool) []*models.Role
	GetByID(id uint) (models.Role, bool)
	DeleteByID(id uint) error
	Update(id uint, menu *models.Role) error
	Create(menu *models.Role) error
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
func (s *roleService) GetAll(args map[string]interface{}, ispreload bool) []*models.Role {
	var meuns []*models.Role

	args["is_admin"] = 0
	db := s.gdb.Where(args)

	if ispreload {
		//db = db.Preload("Child")
	}

	if err := db.Find(&meuns).Error; err != nil {
		panic(err)
	}
	return meuns
}

func (s *roleService) GetByID(id uint) (models.Role, bool) {
	return models.Role{}, true
}

func (s *roleService) Update(id uint, menu *models.Role) error {
	return nil
}

func (s *roleService) Create(menu *models.Role) error {
	var (
		err error
	)
	if menu.ID > 0 {
		return errors.New("unable to create this menu")
	}

	err = s.gdb.Create(menu).Error

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
