package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/multi"
	"gorm.io/gorm"
)

// GetMenuMap
func GetMenuMap(id uint, ctx *gin.Context, isTenancy bool) (Form, error) {
	var form Form
	var formStr string
	if id > 0 {
		menu, err := GetMenuByID(id)
		if err != nil {
			return form, err
		}
		formStr = fmt.Sprintf(`{"rule":[{"type":"cascader","field":"pid","value":%d,"title":"父级分类","props":{"type":"other","options":[],"placeholder":"请选择父级分类","props":{"checkStrictly":true,"emitPath":false}}},{"type":"select","field":"is_menu","value":%d,"title":"权限类型","props":{"multiple":false,"placeholder":"请选择权限类型"},"control":[{"value":2,"rule":[{"type":"input","field":"menu_name","value":"%s","title":"路由名称","props":{"type":"text","placeholder":"请输入路由名称"},"validate":[{"message":"请输入路由名称","required":true,"type":"string","trigger":"change"}]},{"type":"input","field":"params","value":"%s","title":"参数","props":{"type":"textarea","placeholder":"路由参数:\r\nkey1:value1\r\nkey2:value2"}}]},{"value":1,"rule":[{"type":"switch","field":"hidden","value":%d,"title":"是否显示","props":{"activeValue":2,"inactiveValue":1,"inactiveText":"关闭","activeText":"开启"}},{"type":"frame","field":"icon","value":"%s","title":"菜单图标","props":{"type":"input","maxLength":1,"title":"请选择菜单图标","src":"\/admin\/setting\/icons?field=icon","icon":"el-icon-circle-plus-outline","height":"338px","width":"700px","modal":{"modal":false}}},{"type":"input","field":"menu_name","value":"%s","title":"菜单名称","props":{"type":"text","placeholder":"请输入菜单名称"},"validate":[{"message":"请输入菜单名称","required":true,"type":"string","trigger":"change"}]}]}],"options":[{"value":1,"label":"菜单"}]},{"type":"input","field":"route","value":"%s","title":"路由","props":{"type":"text","placeholder":"请输入路由"}},{"type":"inputNumber","field":"sort","value":%d,"title":"排序","props":{"placeholder":"请输入排序"}}],"action":"","method":"POST","title":"编辑菜单","config":{}}`, menu.Pid, menu.IsMenu, menu.MenuName, menu.Params, menu.Hidden, menu.Icon, menu.MenuName, menu.Route, menu.Sort)

	} else {
		formStr = fmt.Sprintf(`{"rule":[{"type":"cascader","field":"pid","value":0,"title":"父级分类","props":{"type":"other","options":[],"placeholder":"请选择父级分类","props":{"checkStrictly":true,"emitPath":false}}},{"type":"select","field":"is_menu","value":%d,"title":"权限类型","props":{"multiple":false,"placeholder":"请选择权限类型"},"control":[{"value":2,"rule":[{"type":"input","field":"menu_name","value":"%s","title":"路由名称","props":{"type":"text","placeholder":"请输入路由名称"},"validate":[{"message":"请输入路由名称","required":true,"type":"string","trigger":"change"}]},{"type":"input","field":"params","value":"%s","title":"参数","props":{"type":"textarea","placeholder":"路由参数:\r\nkey1:value1\r\nkey2:value2"}}]},{"value":1,"rule":[{"type":"switch","field":"hidden","value":%d,"title":"是否显示","props":{"activeValue":2,"inactiveValue":1,"inactiveText":"关闭","activeText":"开启"}},{"type":"frame","field":"icon","value":"","title":"菜单图标","props":{"type":"input","maxLength":1,"title":"请选择菜单图标","src":"\/admin\/setting\/icons?field=icon","icon":"el-icon-circle-plus-outline","height":"338px","width":"700px","modal":{"modal":false}}},{"type":"input","field":"menu_name","value":"%s","title":"菜单名称","props":{"type":"text","placeholder":"请输入菜单名称"},"validate":[{"message":"请输入菜单名称","required":true,"type":"string","trigger":"change"}]}]}],"options":[{"value":1,"label":"菜单"}]},{"type":"input","field":"route","value":"%s","title":"路由","props":{"type":"text","placeholder":"请输入路由"}},{"type":"inputNumber","field":"sort","value":%d,"title":"排序","props":{"placeholder":"请输入排序"}}],"action":"","method":"POST","title":"添加菜单","config":{}}`, 1, "", "", 2, "", "", 0)
	}
	err := json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	if id > 0 {
		form.SetAction(fmt.Sprintf("/menu/updateBaseMenu/%d", id), ctx)
	} else {
		if isTenancy {
			form.SetAction("/menu/addTenancyBaseMenu", ctx)
		} else {
			form.SetAction("/menu/addBaseMenu", ctx)
		}
	}
	opts, err := GetMenusOptions()
	if err != nil {
		return form, err
	}
	form.Rule[0].Props["options"] = opts
	return form, err
}

func GetMenuByID(id uint) (model.SysBaseMenu, error) {
	var menu model.SysBaseMenu
	err := g.TENANCY_DB.Model(&model.SysBaseMenu{}).Where("id = ?", id).First(&menu).Error
	return menu, err
}

// getMenuTreeMap 获取路由总树map
func getMenuTreeMap(ctx *gin.Context) (map[uint][]model.SysMenu, error) {
	var allMenus []model.SysMenu
	treeMap := make(map[uint][]model.SysMenu, 1000)
	db := g.TENANCY_DB.Where("authority_id = ?", multi.GetAuthorityId(ctx))
	if multi.IsAdmin(ctx) {
		db = db.Where("is_tenancy = ?", g.StatusFalse)
	} else if multi.IsTenancy(ctx) {
		db = db.Where("is_tenancy = ?", g.StatusTrue)
	}
	err := db.Where("is_menu = ?", g.StatusTrue).Order("sort desc").Find(&allMenus).Error
	if err != nil {
		return nil, err
	}
	for _, v := range allMenus {
		treeMap[v.Pid] = append(treeMap[v.Pid], v)
	}
	return treeMap, err
}

// GetMenuTree 获取动态菜单树
func GetMenuTree(ctx *gin.Context) ([]model.SysMenu, error) {
	menuTree, err := getMenuTreeMap(ctx)
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
func GetInfoList(userType int) ([]model.SysBaseMenu, error) {
	var menuList []model.SysBaseMenu
	treeMap, err := getBaseMenuTreeMap(userType)
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
	err := g.TENANCY_DB.Where("route = ?", menu.Route).First(&model.SysBaseMenu{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return model.SysBaseMenu{}, errors.New("存在重复route，请修改route")
	}
	err = g.TENANCY_DB.Create(&menu).Error
	return menu, err
}

// getBaseMenuTreeMap 获取路由总树map
func getBaseMenuTreeMap(userType int) (map[uint][]model.SysBaseMenu, error) {
	var allMenus []model.SysBaseMenu
	treeMap := make(map[uint][]model.SysBaseMenu)
	db := g.TENANCY_DB.Where("is_menu = ?", g.StatusTrue)
	if userType == multi.AdminAuthority {
		db = db.Where("is_tenancy = ? ", g.StatusFalse)
	} else if userType == multi.TenancyAuthority {
		db = db.Where("is_tenancy = ? ", g.StatusTrue)
	}
	err := db.Order("sort desc").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.Pid] = append(treeMap[v.Pid], v)
	}
	return treeMap, err
}

// GetBaseMenuTree 获取基础路由树
func GetBaseMenuTree() ([]model.SysBaseMenu, error) {
	treeMap, err := getBaseMenuTreeMap(1)
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
	err := g.TENANCY_DB.Where("authority_id = ? ", info.AuthorityId).Where("is_tenancy = ?", g.StatusFalse).Where("is_menu = ?", g.StatusTrue).Order("sort desc").Find(&menus).Error
	//sql := "SELECT authority_menu.keep_alive,authority_menu.default_menu,authority_menu.created_at,authority_menu.updated_at,authority_menu.deleted_at,authority_menu.menu_level,authority_menu.parent_id,authority_menu.path,authority_menu.`name`,authority_menu.hidden,authority_menu.component,authority_menu.title,authority_menu.icon,authority_menu.sort,authority_menu.menu_id,authority_menu.authority_id FROM authority_menu WHERE authority_menu.authority_id = ? ORDER BY authority_menu.sort ASC"
	//err = g.TENANCY_DB.Raw(sql, authorityId).Scan(&menus).Error
	return menus, err
}

// GetMenusOptions
func GetMenusOptions() ([]Option, error) {
	var options []Option
	options = append(options, Option{Label: "请选择", Value: 0})
	treeMap, err := getBaseMenuTreeMap(1)

	for _, opt := range treeMap[0] {
		options = append(options, Option{Label: opt.MenuName, Value: opt.ID})
	}
	for i := 0; i < len(options); i++ {
		getMenusOption(&options[i], treeMap)
	}

	return options, err
}

// getMenusOption
func getMenusOption(op *Option, treeMap map[uint][]model.SysBaseMenu) {
	id, ok := op.Value.(uint)
	if ok {
		for _, opt := range treeMap[id] {
			op.Children = append(op.Children, Option{Label: opt.MenuName, Value: opt.ID})
		}
		for i := 0; i < len(op.Children); i++ {
			getMenusOption(&op.Children[i], treeMap)
		}
	}
}
