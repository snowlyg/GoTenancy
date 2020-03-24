package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/snowlyg/go-tenancy/common"
	"github.com/snowlyg/go-tenancy/sysinit"
)

type InitController struct {
	Ctx     iris.Context
	Session *sessions.Session
}

// GetInfo handles GET: http://localhost:8080/user/table.
func (c *InitController) GetInfo() interface{} {

	authUser := GetAuthUser(c.Session)
	// 管理员 返回管理菜单 	filter["type"] = 2
	// 商户 返回商户菜单 filter["type"] = 3
	filter := map[string]interface{}{"is_menu": 1, "parent_id": 0}
	if authUser.TenantId == 0 {
		filter["type"] = 2
	} else if authUser.TenantId > 0 {
		filter["type"] = 3
	}

	_, perms := sysinit.PermService.GetAll(filter, true)

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
