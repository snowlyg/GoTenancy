package services

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/snowlyg/go-tenancy/models"
)

type MenuService interface {
	GetAll() []*models.Menu
	GetByID(id int64) (models.Menu, bool)
	DeleteByID(id int64) bool
	Update(id int64, menu *models.Menu) error
	Create(menu *models.Menu) error
}

func NewMenuService(gdb *gorm.DB) MenuService {
	return &menuService{
		gdb: gdb,
	}
}

type menuService struct {
	gdb *gorm.DB
}

func (s *menuService) GetAll() []*models.Menu {
	var meuns []*models.Menu
	if err := s.gdb.Where("parent_id = ?", 0).Preload("Child").Find(&meuns).Error; err != nil {
		panic(err)
	}
	return meuns
}

func (s *menuService) GetByID(id int64) (models.Menu, bool) {
	return models.Menu{}, true
}

func (s *menuService) Update(id int64, menu *models.Menu) error {
	return nil
}

func (s *menuService) Create(menu *models.Menu) error {
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

func (s *menuService) DeleteByID(id int64) bool {
	return true
}
