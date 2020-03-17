package controllers

import (
	"fmt"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/snowlyg/go-tenancy/common"
	"github.com/snowlyg/go-tenancy/models"
	"github.com/snowlyg/go-tenancy/sysinit"
	"github.com/snowlyg/go-tenancy/transformer"
	"github.com/snowlyg/gotransformer"
)

//GetInitMenus 站点头部，侧边栏菜单接口
func GetInitMenus(ctx iris.Context) {
	menuinfo := sysinit.MenuService.GetAll()

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
		MenuInfo: menuinfo,
	}
	if _, err := ctx.JSON(init); err != nil {
		panic(err)
	}
}

//GetMenus 菜单表格接口
func GetMenus(ctx iris.Context) {
	menus := sysinit.MenuService.GetAll()
	ttms := transformerTableMenus(menus)

	ctx.StatusCode(iris.StatusOK)
	if _, err := ctx.JSON(common.Table{Code: 0, Msg: "", Count: len(ttms), Data: ttms}); err != nil {
		panic(err)
	}
}

// transformerTableMenus 菜单表格接口数据转换
func transformerTableMenus(menus []*models.Menu) []*transformer.TableMenu {
	var tableMenus []*transformer.TableMenu

	for mun, menu := range menus {
		tableMenu := &transformer.TableMenu{}
		g := gotransformer.NewTransform(tableMenu, menu, time.RFC3339)
		err := g.Transformer()
		if err != nil {
			panic(fmt.Sprintf("菜单表格接口数据转换错误：%v", err))
		}

		tableMenu.OrderNumber = mun + 1

		if menu.ParentId == 0 {
			tableMenu.ParentId = -1
		}
		tableMenus = append(tableMenus, tableMenu)
	}

	return tableMenus
}
