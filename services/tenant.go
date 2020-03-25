package services

import (
	"database/sql"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/snowlyg/go-tenancy/common"
	"github.com/snowlyg/go-tenancy/config"
	"github.com/snowlyg/go-tenancy/models"
)

type TenantService interface {
	GetAll(args map[string]interface{}, pagination *common.Pagination, ispreload bool) (int64, []*models.Tenant)
	GetByID(id uint) (models.Tenant, bool)
	DeleteByID(id uint) error
	Update(id uint, tenant *models.Tenant) error
	Create(tenant *models.Tenant) error
	DeleteMnutil(tenantIds []common.Id) error
}

func NewTenantService(gdb *gorm.DB, userService UserService, roleService RoleService) TenantService {
	return &tenantService{
		gdb:         gdb,
		UserService: userService,
		RoleService: roleService,
	}
}

type tenantService struct {
	gdb         *gorm.DB
	UserService UserService
	RoleService RoleService
}

//GetAll 查询所有数据
//args 过滤条件 {"parent_id = ?" : 0}
func (s *tenantService) GetAll(args map[string]interface{}, pagination *common.Pagination, ispreload bool) (int64, []*models.Tenant) {
	var tenants []*models.Tenant
	var count int64

	db := s.gdb.Where(args).Order("id desc")

	if ispreload {
		//db = db.Preload("Child")
	}

	db.Find(&tenants).Count(&count)

	if pagination != nil {
		db = db.Limit(pagination.Limit).Offset(pagination.Limit * (pagination.Page - 1))
	}

	if err := db.Find(&tenants).Error; err != nil {
		panic(err)
	}
	return count, tenants
}

func (s *tenantService) GetByID(id uint) (models.Tenant, bool) {
	tenant := models.Tenant{Model: gorm.Model{ID: id}}
	if notFound := s.gdb.Find(&tenant).RecordNotFound(); notFound {
		return tenant, false
	}
	return tenant, true
}

func (s *tenantService) UpdateTenant(id uint, tenant *models.Tenant) error {

	user := models.User{Model: gorm.Model{ID: id}}
	if config.Config.DB.Adapter != "mysql" {
		if err := s.gdb.Model(user).Where(NotAdmin).Update(tenant).Error; err != nil {
			return err
		}

		return nil
	} else {
		return s.gdb.Transaction(func(tx *gorm.DB) error {
			if err := tx.Model(user).Where(NotAdmin).Update(tenant).Error; err != nil {
				return err
			}

			return nil
		})
	}

}

func (s *tenantService) Update(id uint, tenant *models.Tenant) error {

	if err := s.gdb.Where("id = ?", id).Where(NotAdmin).Update(tenant).Error; err != nil {
		return err
	}
	return nil

}

func (s *tenantService) Create(tenant *models.Tenant) error {

	if tenant.ID > 0 {
		return errors.New("unable to create this tenant")
	}

	if config.Config.DB.Adapter != "mysql" {
		if err := s.gdb.Create(tenant).Error; err != nil {
			return err
		}

		tenantrole, _ := s.RoleService.GetByName(config.Config.Tenant.RoleName)
		tenantuser := &models.User{
			Username: "tenantname",
			Name:     "商户管理员",
			Email:    "tenant@admin.com",
			Telphone: "13800138001",
			IsAdmin:  sql.NullBool{Bool: false, Valid: true},
			Model:    gorm.Model{CreatedAt: time.Now()},
			RoleIds:  []uint{tenantrole.ID},
		}

		if err := s.UserService.Create("password", tenantuser); err != nil {
			return err
		}

		return nil
	} else {
		return s.gdb.Transaction(func(tx *gorm.DB) error {
			if err := s.gdb.Create(tenant).Error; err != nil {
				return err
			}

			tenantrole, _ := s.RoleService.GetByName(config.Config.Tenant.RoleName)
			tenantuser := &models.User{
				Username: "tenantname",
				Name:     "商户管理员",
				Email:    "tenant@admin.com",
				Telphone: "13800138001",
				IsAdmin:  sql.NullBool{Bool: false, Valid: true},
				Model:    gorm.Model{CreatedAt: time.Now()},
				RoleIds:  []uint{tenantrole.ID},
			}

			if err := s.UserService.Create("password", tenantuser); err != nil {
				return err
			}

			return nil
		})
	}

}

func (s *tenantService) DeleteByID(id uint) error {
	if err := s.gdb.Delete(models.Tenant{Model: gorm.Model{ID: id}}).Error; err != nil {
		return err
	}
	return nil
}

func (s *tenantService) DeleteMnutil(tenantIds []common.Id) error {
	return s.gdb.Transaction(func(tx *gorm.DB) error {
		for _, tenantId := range tenantIds {
			if err := tx.Delete(models.Tenant{Model: gorm.Model{ID: uint(tenantId.Id)}}).Error; err != nil {
				return err
			}
		}

		// 返回 nil 提交事务
		return nil
	})
}
