package services

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/snowlyg/go-tenancy/models"
)

type PermService interface {
	GetAll(args map[string]interface{}, ispreload bool) (int64, []*models.Perm)
	GetByID(id uint) (models.Perm, bool)
	DeleteByID(id uint) bool
	GetPermissionByHrefMethod(href, method string) (models.Perm, bool)
	Update(id uint, menu *models.Perm) error
	Create(menu *models.Perm) error
}

func NewPermService(gdb *gorm.DB) PermService {
	return &permService{
		gdb: gdb,
	}
}

type permService struct {
	gdb *gorm.DB
}

//GetAll 查询所有数据
//args 过滤条件 {"parent_id = ?" : 0}
func (s *permService) GetAll(args map[string]interface{}, ispreload bool) (int64, []*models.Perm) {
	var meuns []*models.Perm
	var count int64

	db := s.gdb.Where(args)

	if ispreload {
		db = db.Preload("Child")
	}

	db.Find(&meuns).Count(&count)

	if err := db.Find(&meuns).Error; err != nil {
		panic(err)
	}
	return count, meuns
}

func (s *permService) GetPermissionByHrefMethod(href, method string) (models.Perm, bool) {
	var perm models.Perm
	if notFound := s.gdb.Where("href = ?", href).Where("method = ?", method).Find(&perm).RecordNotFound(); notFound {
		return perm, false
	}
	return perm, true
}

func (s *permService) GetByID(id uint) (models.Perm, bool) {
	return models.Perm{}, true
}

func (s *permService) Update(id uint, menu *models.Perm) error {
	return nil
}

func (s *permService) Create(menu *models.Perm) error {
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

func (s *permService) DeleteByID(id uint) bool {
	return true
}
