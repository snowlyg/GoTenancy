package services

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/snowlyg/go-tenancy/models"
)

type PermService interface {
	GetAll(args interface{}, ispreload bool) []*models.Perm
	GetByID(id int64) (models.Perm, bool)
	DeleteByID(id int64) bool
	Update(id int64, menu *models.Perm) error
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
func (s *permService) GetAll(args interface{}, ispreload bool) []*models.Perm {
	var meuns []*models.Perm
	db := s.gdb.Where(args)

	if ispreload {
		db = db.Preload("Child")
	}

	if err := db.Find(&meuns).Error; err != nil {
		panic(err)
	}
	return meuns
}

func (s *permService) GetByID(id int64) (models.Perm, bool) {
	return models.Perm{}, true
}

func (s *permService) Update(id int64, menu *models.Perm) error {
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

func (s *permService) DeleteByID(id int64) bool {
	return true
}
