package router

import (
	"github.com/kataras/iris/v12"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
	"github.com/snowlyg/go-tenancy/middleware"
)

func InitMenuRouter(Router iris.Party) (R iris.Party) {
	MenuRouter := Router.Party("/menu", middleware.OperationRecord())
	{
		MenuRouter.Post("/getMenu", v1.GetMenu)                   // 获取菜单树
		MenuRouter.Post("/getMenuList", v1.GetMenuList)           // 分页获取基础menu列表
		MenuRouter.Post("/addBaseMenu", v1.AddBaseMenu)           // 新增菜单
		MenuRouter.Post("/getBaseMenuTree", v1.GetBaseMenuTree)   // 获取用户动态路由
		MenuRouter.Post("/addMenuAuthority", v1.AddMenuAuthority) //	增加menu和角色关联关系
		MenuRouter.Post("/getMenuAuthority", v1.GetMenuAuthority) // 获取指定角色menu
		MenuRouter.Post("/deleteBaseMenu", v1.DeleteBaseMenu)     // 删除菜单
		MenuRouter.Post("/updateBaseMenu", v1.UpdateBaseMenu)     // 更新菜单
		MenuRouter.Post("/getBaseMenuById", v1.GetBaseMenuById)   // 根据id获取菜单
	}
	return MenuRouter
}
