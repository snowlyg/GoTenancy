package service

import (
	"errors"
	"fmt"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

// DeleteBaseMenu 删除基础路由
func DeleteBaseMenu(id uint) error {
	err := g.TENANCY_DB.Where("pid = ?", id).First(&model.SysBaseMenu{}).Error
	if err != nil {
		var menu model.SysBaseMenu
		err := g.TENANCY_DB.Preload("SysAuthoritys").Where("id = ?", id).First(&menu).Error
		if err != nil {
			return fmt.Errorf("not found menu %w", err)
		}
		if len(menu.SysAuthoritys) > 0 {
			if err = g.TENANCY_DB.Model(&menu).Association("SysAuthoritys").Delete(&menu.SysAuthoritys); err != nil {
				return fmt.Errorf("del SysAuthoritys %w", err)
			}
		}
		if err := g.TENANCY_DB.Delete(&menu).Error; err != nil {
			return fmt.Errorf("del menu %w", err)
		}

	} else {
		return errors.New("此菜单存在子菜单不可删除")
	}
	return nil
}

// UpdateBaseMenu 更新路由
func UpdateBaseMenu(menu model.SysBaseMenu) error {
	var oldMenu model.SysBaseMenu
	upDateMap := make(map[string]interface{})
	upDateMap["pid"] = menu.Pid
	upDateMap["path"] = menu.Path
	upDateMap["menu_name"] = menu.MenuName
	upDateMap["hidden"] = menu.Hidden
	upDateMap["route"] = menu.Route
	upDateMap["icon"] = menu.Icon
	upDateMap["sort"] = menu.Sort

	err := g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		db := tx.Where("id = ?", menu.ID).Find(&oldMenu)
		if oldMenu.MenuName != menu.MenuName {
			err := tx.Where("id <> ? AND menu_name = ?", menu.ID, menu.MenuName).First(&model.SysBaseMenu{}).Error
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				g.TENANCY_LOG.Debug("存在相同name修改失败")
				return errors.New("存在相同name修改失败")
			}
		}

		txErr := db.Updates(upDateMap).Error
		if txErr != nil {
			g.TENANCY_LOG.Debug(txErr.Error())
			return txErr
		}
		return nil
	})
	return err
}

// GetBaseMenuById 返回当前选中menu
func GetBaseMenuById(id uint) (model.SysBaseMenu, error) {
	var menu model.SysBaseMenu
	err := g.TENANCY_DB.Where("id = ?", id).First(&menu).Error
	return menu, err
}
