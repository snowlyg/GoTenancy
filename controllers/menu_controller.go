package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/snowlyg/go-tenancy/common"
	"github.com/snowlyg/go-tenancy/models"
	"github.com/snowlyg/go-tenancy/services"
	"github.com/snowlyg/go-tenancy/sysinit"
)

type MenuController struct {
	Ctx     iris.Context
	Service services.PermService
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
	perms := sysinit.PermService.GetAll(args, false)
	transformerTableMenus(perms)

	return common.Table{Code: 0, Msg: "", Count: len(perms), Data: perms}
}

// transformerTableMenus 菜单表格接口数据转换
func transformerTableMenus(perms []*models.Perm) {
	for _, perm := range perms {
		if perm.ParentId == 0 {
			perm.ParentId = -1
		}
	}

}
