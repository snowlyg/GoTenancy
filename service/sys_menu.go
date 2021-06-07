package service

import (
	"errors"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"gorm.io/gorm"
)

// getMenuTreeMap 获取路由总树map
func getMenuTreeMap(authorityId string) (map[uint][]model.SysMenu, error) {
	var allMenus []model.SysMenu
	treeMap := make(map[uint][]model.SysMenu, 1000)
	err := g.TENANCY_DB.Where("authority_id = ?", authorityId).Where("is_tenancy = ?", g.StatusFalse).Where("is_menu = ?", g.StatusTrue).Order("sort desc").Find(&allMenus).Error
	if err != nil {
		return nil, err
	}
	for _, v := range allMenus {
		treeMap[v.Pid] = append(treeMap[v.Pid], v)
	}
	return treeMap, err
}

// GetMenuTree 获取动态菜单树
func GetMenuTree(authorityId string) ([]model.SysMenu, error) {
	menuTree, err := getMenuTreeMap(authorityId)
	menus := menuTree[0]
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(&menus[i], menuTree)
	}
	return menus, err
}

// getChildrenList 获取子菜单
func getChildrenList(menu *model.SysMenu, treeMap map[uint][]model.SysMenu) error {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		err := getChildrenList(&menu.Children[i], treeMap)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetInfoList 获取路由分页
func GetInfoList() ([]model.SysBaseMenu, error) {
	var menuList []model.SysBaseMenu
	treeMap, err := getBaseMenuTreeMap()
	menuList = treeMap[0]
	for i := 0; i < len(menuList); i++ {
		err = getBaseChildrenList(&menuList[i], treeMap)
	}
	return menuList, err
}

// getBaseChildrenList 获取菜单的子菜单
func getBaseChildrenList(menu *model.SysBaseMenu, treeMap map[uint][]model.SysBaseMenu) (err error) {
	menu.Children = treeMap[menu.ID]
	for i := 0; i < len(menu.Children); i++ {
		err = getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

// AddBaseMenu 添加基础路由
func AddBaseMenu(menu model.SysBaseMenu) (model.SysBaseMenu, error) {
	err := g.TENANCY_DB.Where("menu_name = ?", menu.MenuName).First(&model.SysBaseMenu{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return model.SysBaseMenu{}, errors.New("存在重复name，请修改name")
	}
	err = g.TENANCY_DB.Create(&menu).Error
	return menu, err
}

// getBaseMenuTreeMap 获取路由总树map
func getBaseMenuTreeMap() (map[uint][]model.SysBaseMenu, error) {
	var allMenus []model.SysBaseMenu
	treeMap := make(map[uint][]model.SysBaseMenu)
	err := g.TENANCY_DB.Order("sort desc").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.Pid] = append(treeMap[v.Pid], v)
	}
	return treeMap, err
}

// GetBaseMenuTree 获取基础路由树
func GetBaseMenuTree() ([]model.SysBaseMenu, error) {
	treeMap, err := getBaseMenuTreeMap()
	if err != nil {
		return nil, err
	}
	menus := treeMap[0]
	for i := 0; i < len(menus); i++ {
		err = getBaseChildrenList(&menus[i], treeMap)
	}
	return menus, err
}

// AddMenuAuthority 为角色增加menu树
func AddMenuAuthority(menus []model.SysBaseMenu, authorityId string) error {
	var auth model.SysAuthority
	auth.AuthorityId = authorityId
	auth.SysBaseMenus = menus
	return SetMenuAuthority(&auth)
}

// GetMenuAuthority 查看当前角色树
func GetMenuAuthority(info *request.GetAuthorityId) ([]model.SysMenu, error) {
	var menus []model.SysMenu
	err := g.TENANCY_DB.Where("authority_id = ? ", info.AuthorityId).Order("sort desc").Find(&menus).Error
	//sql := "SELECT authority_menu.keep_alive,authority_menu.default_menu,authority_menu.created_at,authority_menu.updated_at,authority_menu.deleted_at,authority_menu.menu_level,authority_menu.parent_id,authority_menu.path,authority_menu.`name`,authority_menu.hidden,authority_menu.component,authority_menu.title,authority_menu.icon,authority_menu.sort,authority_menu.menu_id,authority_menu.authority_id FROM authority_menu WHERE authority_menu.authority_id = ? ORDER BY authority_menu.sort ASC"
	//err = g.TENANCY_DB.Raw(sql, authorityId).Scan(&menus).Error
	return menus, err
}
