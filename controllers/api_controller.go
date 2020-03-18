package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/go-tenancy/common"
	"github.com/snowlyg/go-tenancy/models"
	"github.com/snowlyg/go-tenancy/sysinit"
)

//GetInitMenus 站点头部，侧边栏菜单接口
func GetInitMenus(ctx iris.Context) {
	args := map[string]interface{}{
		"parent_id = ?": 0,
		"is_menu = ?":   1,
	}
	perms := sysinit.PermService.GetAll(args, true)

	ctx.StatusCode(iris.StatusOK)
	init := common.Menus{
		HomeInfo: struct {
			Title string `json:"title"`
			Href  string `json:"href"`
		}{
			Title: "首页",
			Href:  "control",
		},
		LogoInfo: struct {
			Title string `json:"title"`
			Href  string `json:"href"`
			Image string `json:"image"`
		}{
			Title: "GOTENACY",
			Href:  "",
			Image: "./public/images/logo.png",
		},
		MenuInfo: perms,
	}
	if _, err := ctx.JSON(init); err != nil {
		panic(err)
	}
}

//GetMenus 菜单表格接口
func GetMenus(ctx iris.Context) {
	args := map[string]interface{}{}
	perms := sysinit.PermService.GetAll(args, false)
	transformerTableMenus(perms)

	ctx.StatusCode(iris.StatusOK)
	if _, err := ctx.JSON(common.Table{Code: 0, Msg: "", Count: len(perms), Data: perms}); err != nil {
		panic(err)
	}
}

// transformerTableMenus 菜单表格接口数据转换
func transformerTableMenus(perms []*models.Perm) {
	for _, perm := range perms {
		if perm.ParentId == 0 {
			perm.ParentId = -1
		}
	}

}
