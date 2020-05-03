package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/go-tenancy/common"
	"github.com/snowlyg/go-tenancy/sysinit"
)

type InitController struct {
	Ctx iris.Context
}

// GetInfo handles GET: http://localhost:8080/user/table.
func (c *InitController) GetInfo() interface{} {

	authUser := GetAuthUser()
	// 管理员 返回管理菜单 	filter["type"] = 2
	// 商户 返回商户菜单 filter["type"] = 3
	typefilters := []string{"1"}
	if authUser.IsAdmin.Bool {
		typefilters = append(typefilters, "2")
	} else {
		typefilters = append(typefilters, "3")
	}

	_, perms := sysinit.PermService.GetAll(map[string]interface{}{"is_menu": 1, "parent_id": 0}, typefilters, true)

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
