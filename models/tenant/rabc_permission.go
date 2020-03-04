package tenant

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	"go-tenancy/app/home/homevalidates"
	"go-tenancy/config/db"
	"go-tenancy/utils"
)

type RabcPermission struct {
	gorm.Model
	Name        string `gorm:"not null VARCHAR(191)"`
	DisplayName string `gorm:"VARCHAR(191)"`
	Description string `gorm:"VARCHAR(191)"`
	Act         string `gorm:"VARCHAR(191)"`
}

func NewRabcPermission(id uint, name, act string) *RabcPermission {
	return &RabcPermission{
		Model: gorm.Model{
			ID:        id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name: name,
		Act:  act,
	}
}

/**
* 通过 id 获取 permission 记录
* @method GetPermissionById
* @param  {[type]}       permission  *Permission [description]
 */
func (p *RabcPermission) GetRabcPermissionById() {
	_ = db.DB.Where("id = ?", p.ID).First(p).Error
}

/**
* 通过 name 获取 permission 记录
* @method GetPermissionByName
* @param  {[type]}       permission  *Permission [description]
 */
func (p *RabcPermission) GetRabcPermissionByNameAct() {
	_ = db.DB.Where("name = ?", p.Name).Where("act = ?", p.Act).First(p).Error
}

/**
* 通过 id 删除权限
* @method DeletePermissionById
 */
func (p *RabcPermission) DeleteRabcPermissionById() {
	if err := db.DB.Delete(p).Error; err != nil {
		color.Red(fmt.Sprintf("DeletePermissionByIdError:%s \n", err))
	}
}

/**
* 获取所有的权限
* @method GetAllPermissions
* @param  {[type]} name string [description]
* @param  {[type]} orderBy string [description]
* @param  {[type]} offset int    [description]
* @param  {[type]} limit int    [description]
 */
func GetAllRabcPermissions(name, orderBy string, offset, limit int) (permissions []*RabcPermission) {
	if err := utils.GetAll(name, orderBy, offset, limit).Find(&permissions).Error; err != nil {
		color.Red(fmt.Sprintf("GetAllPermissionsError:%s \n", err))
	}

	return
}

/**
* 创建
* @method CreatePermission
* @param  {[type]} kw string [description]
* @param  {[type]} cp int    [description]
* @param  {[type]} mp int    [description]
 */
func (p *RabcPermission) CreateRabcPermission() error {
	if err := db.DB.Create(p).Error; err != nil {
		color.Red(fmt.Sprintf("CreatePermissionError:%s \n", err))
		return err
	}
	return nil
}

/**
* 更新
* @method UpdatePermission
* @param  {[type]} kw string [description]
* @param  {[type]} cp int    [description]
* @param  {[type]} mp int    [description]
 */
func (p *RabcPermission) UpdateRabcPermission(pj *homevalidates.RabcPermissionRequest) {
	if err := utils.Update(p, pj); err != nil {
		color.Red(fmt.Sprintf("UpdatePermissionError:%s \n", err))
	}
}
