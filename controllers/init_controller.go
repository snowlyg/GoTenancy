package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/go-tenancy/common"
	"github.com/snowlyg/go-tenancy/services"
	"github.com/snowlyg/go-tenancy/sysinit"
)

type InitController struct {
	Ctx     iris.Context
	Service services.PermService
}

// GetInfo handles GET: http://localhost:8080/user/table.
func (c *InitController) GetInfo() interface{} {
	perms := sysinit.PermService.GetAll(map[string]interface{}{"is_menu": 1, "parent_id": ""}, true)

	return common.Menus{
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

}
