package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/snowlyg/go-tenancy/common"
	"github.com/snowlyg/go-tenancy/models"
	"github.com/snowlyg/go-tenancy/sysinit"
	"github.com/snowlyg/go-tenancy/transformer"
	"github.com/snowlyg/gotransformer"
)

type MenuController struct {
	Ctx iris.Context
}

// Get handles GET: http://localhost:8080/menu.
func (c *MenuController) Get() mvc.Result {
	return mvc.View{
		Name: "menu/index.html",
	}
}

// Get handles GET: http://localhost:8080/menu/table.
func (c *MenuController) GetTable() interface{} {
	args := map[string]interface{}{}
	count, perms := sysinit.PermService.GetAll(args, false)
	tablemenus := transformerTableMenus(perms)

	return common.Table{Code: 0, Msg: "", Count: count, Data: tablemenus}
}

// transformerTableMenus 菜单表格接口数据转换
func transformerTableMenus(perms []*models.Perm) []*transformer.MenuTable {
	var tablemenus []*transformer.MenuTable
	for _, perm := range perms {
		tablemenu := &transformer.MenuTable{}
		g := gotransformer.NewTransform(tablemenu, perm, "")
		_ = g.Transformer()

		if perm.ParentId.Int64 == 0 {
			tablemenu.ParentId = -1
		} else {
			tablemenu.ParentId = perm.ParentId.Int64
		}

		tablemenus = append(tablemenus, tablemenu)
	}
	return tablemenus
}
