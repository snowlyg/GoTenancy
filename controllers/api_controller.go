package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/go-tenancy/api_types"
)

func GetMenus(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	menus := api_types.Menus{
		HomeInfo: struct {
			Title string `json:"title"`
			Href  string `json:"href"`
		}{
			Title: "首页",
			Href:  "/control",
		},
		LogoInfo: struct {
			Title string `json:"title"`
			Href  string `json:"href"`
			Image string `json:"image"`
		}{
			Title: "LAYUI MINI",
			Href:  "",
			Image: "./public/images/logo.png",
		},
		MenuInfo: []struct {
			Title  string `json:"title"`
			Href   string `json:"href"`
			Icon   string `json:"icon"`
			Target string `json:"target"`
			Child  []struct {
				Title  string `json:"title"`
				Href   string `json:"href"`
				Icon   string `json:"icon"`
				Target string `json:"target"`
				Child  []struct {
					Title  string `json:"title"`
					Href   string `json:"href"`
					Icon   string `json:"icon"`
					Target string `json:"target"`
				} `json:"child"`
			} `json:"child"`
		}{
			{
				Title:  "常规管理",
				Href:   "fa fa-address-book",
				Icon:   "",
				Target: "_self",
				Child: []struct {
					Title  string `json:"title"`
					Href   string `json:"href"`
					Icon   string `json:"icon"`
					Target string `json:"target"`
					Child  []struct {
						Title  string `json:"title"`
						Href   string `json:"href"`
						Icon   string `json:"icon"`
						Target string `json:"target"`
					} `json:"child"`
				}{
					{
						Title:  "主页一",
						Href:   "page/welcome-1.html",
						Icon:   "fa fa-tachometer",
						Target: "_self",
					},
					{
						Title:  "主页二",
						Href:   "page/welcome-2.html",
						Icon:   "fa fa-tachometer",
						Target: "_self",
					},
					{
						Title:  "主页三",
						Href:   "page/welcome-3.html",
						Icon:   "fa fa-tachometer",
						Target: "_self",
					},
				},
			},
			{
				Title:  "菜单管理",
				Href:   "page/menu.html",
				Icon:   "fa fa-window-maximize",
				Target: "_self",
			},
			{
				Title:  "系统设置",
				Href:   "page/setting.html",
				Icon:   "fa fa-gears",
				Target: "_self",
			},
			{
				Title:  "表格示例",
				Href:   "page/table.html",
				Icon:   "fa fa-file-text",
				Target: "_self",
			},
			{
				Title:  "表单示例",
				Href:   "",
				Icon:   "fa fa-calendar",
				Target: "_self",
				Child: []struct {
					Title  string `json:"title"`
					Href   string `json:"href"`
					Icon   string `json:"icon"`
					Target string `json:"target"`
					Child  []struct {
						Title  string `json:"title"`
						Href   string `json:"href"`
						Icon   string `json:"icon"`
						Target string `json:"target"`
					} `json:"child"`
				}{
					{
						Title:  "普通表单",
						Href:   "page/form.html",
						Icon:   "fa fa-list-alt",
						Target: "_self",
					},
					{
						Title:  "分步表单",
						Href:   "page/form-step.html",
						Icon:   "fa fa-navicon",
						Target: "_self",
					},
				},
			},
			{
				Title:  "登录模板",
				Href:   "",
				Icon:   "fa fa-flag-o",
				Target: "_self",
				Child: []struct {
					Title  string `json:"title"`
					Href   string `json:"href"`
					Icon   string `json:"icon"`
					Target string `json:"target"`
					Child  []struct {
						Title  string `json:"title"`
						Href   string `json:"href"`
						Icon   string `json:"icon"`
						Target string `json:"target"`
					} `json:"child"`
				}{
					{
						Title:  "登录-1",
						Href:   "page/login-1.html",
						Icon:   "fa fa-stumbleupon-circle",
						Target: "_blank",
					},
					{
						Title:  "登录-2",
						Href:   "page/login-2.html",
						Icon:   "fa fa-viacoin",
						Target: "_blank",
					},
				},
			},
			{
				Title:  "异常页面",
				Href:   "",
				Icon:   "fa fa-home",
				Target: "_self",
				Child: []struct {
					Title  string `json:"title"`
					Href   string `json:"href"`
					Icon   string `json:"icon"`
					Target string `json:"target"`
					Child  []struct {
						Title  string `json:"title"`
						Href   string `json:"href"`
						Icon   string `json:"icon"`
						Target string `json:"target"`
					} `json:"child"`
				}{
					{
						Title:  "404页面",
						Href:   "page/404.html",
						Icon:   "fa fa-hourglass-end",
						Target: "_self",
					},
				},
			},
			{
				Title:  "其它界面",
				Href:   "",
				Icon:   "fa fa-snowflake-o",
				Target: "",
				Child: []struct {
					Title  string `json:"title"`
					Href   string `json:"href"`
					Icon   string `json:"icon"`
					Target string `json:"target"`
					Child  []struct {
						Title  string `json:"title"`
						Href   string `json:"href"`
						Icon   string `json:"icon"`
						Target string `json:"target"`
					} `json:"child"`
				}{
					{
						Title:  "按钮示例",
						Href:   "page/button.html",
						Icon:   "fa fa-snowflake-o",
						Target: "_self",
					},
					{
						Title:  "弹出层",
						Href:   "page/layer.html",
						Icon:   "fa fa-shield",
						Target: "_self",
					},
				},
			},
		},
	}
	if _, err := ctx.JSON(menus); err != nil {
		panic(err)
	}
}
