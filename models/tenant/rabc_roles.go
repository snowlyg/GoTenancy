package tenant

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	"go-tenancy/app/home/homevalidates"
	"go-tenancy/config/db"
	"go-tenancy/utils"
)

type RabcRole struct {
	gorm.Model

	Name        string `gorm:"unique;not null VARCHAR(191)"`
	DisplayName string `gorm:"VARCHAR(191)"`
	Description string `gorm:"VARCHAR(191)"`
}

func NewRabcRole(id uint, name string) *RabcRole {
	return &RabcRole{
		Model: gorm.Model{
			ID:        id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name: name,
	}
}

func NewRabcRoleByStruct(rr *homevalidates.RabcRoleRequest) *RabcRole {
	return &RabcRole{
		Model: gorm.Model{
			ID:        0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:        rr.Name,
		DisplayName: rr.DisplayName,
		Description: rr.Description,
	}
}

/**
* 通过 id 获取 role 记录
* @method GetRoleById
* @param  {[type]}       role  *Role [description]
 */
func (r *RabcRole) GetRabcRoleById() {
	_ = db.DB.Where("id = ?", r.ID).First(r).Error
}

/**
* 通过 name 获取 role 记录
* @method GetRoleByName
* @param  {[type]}       role  *Role [description]
 */
func (r *RabcRole) GetRabcRoleByName() {
	_ = db.DB.Where("name = ?", r.Name).First(r).Error
}

/**
* 通过 id 删除角色
* @method DeleteRoleById
 */
func (r *RabcRole) DeleteRabcRoleById() {
	if err := db.DB.Delete(r).Error; err != nil {
		color.Red(fmt.Sprintf("DeleteRoleErr:%s \n", err))
	}
}

/**
* 获取所有的角色
* @method GetAllRole
* @param  {[type]} name string [description]
* @param  {[type]} orderBy string [description]
* @param  {[type]} offset int    [description]
* @param  {[type]} limit int    [description]
 */
func GetAllRabcRoles(name, orderBy string, offset, limit int) (roles []*RabcRole) {

	if err := utils.GetAll(name, orderBy, offset, limit).Find(&roles).Error; err != nil {
		color.Red(fmt.Sprintf("GetAllRoleErr:%s \n", err))
	}
	return
}

/**
* 创建
* @method CreateRole
* @param  {[type]} kw string [description]
* @param  {[type]} cp int    [description]
* @param  {[type]} mp int    [description]
 */
func (r *RabcRole) CreateRabcRole(permIds []uint) error {
	if err := db.DB.Create(r).Error; err != nil {
		color.Red(fmt.Sprintf("CreateRoleErr:%v \n", err))
		return err
	}

	addPerms(permIds, r)

	return nil
}

func addPerms(permIds []uint, role *RabcRole) {
	if len(permIds) > 0 {
		roleId := strconv.FormatUint(uint64(role.ID), 10)
		if _, err := db.GetCasbinEnforcer().DeletePermissionsForUser(roleId); err != nil {
			color.Red(fmt.Sprintf("AppendPermsErr:%s \n", err))
		}
		var perms []RabcPermission
		db.DB.Where("id in (?)", permIds).Find(&perms)
		for _, perm := range perms {
			if _, err := db.GetCasbinEnforcer().AddPolicy(roleId, perm.Name, perm.Act); err != nil {
				color.Red(fmt.Sprintf("AddPolicy:%s \n", err))
			}
		}
	}
}

/**
* 更新
* @method UpdateRole
* @param  {[type]} kw string [description]
* @param  {[type]} cp int    [description]
* @param  {[type]} mp int    [description]
 */
func (r *RabcRole) UpdateRabcRole(rj *homevalidates.RabcRoleRequest, permIds []uint) {

	if err := utils.Update(r, rj); err != nil {
		color.Red(fmt.Sprintf("UpdatRoleErr:%s \n", err))
	}

	addPerms(permIds, r)

	return
}

// 角色权限
func (r *RabcRole) RabcRoleRabcPermisions() []*RabcPermission {
	perms := utils.GetPermissionsForUser(r.ID)
	var ps []*RabcPermission
	for _, perm := range perms {
		if len(perm) >= 3 && len(perm[1]) > 0 && len(perm[2]) > 0 {
			p := NewRabcPermission(0, perm[1], perm[2])
			p.GetRabcPermissionByNameAct()
			if p.ID > 0 {
				ps = append(ps, p)
			}
		}
	}
	return ps
}
