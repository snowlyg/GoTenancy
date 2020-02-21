package home

import (
	"net/http"

	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/utils"
	"GoTenancy/libs/render"
	"github.com/kataras/iris/v12"
)

// Controller home controller
type Controller struct {
	View *render.Render
}

// Index home index page
func (ctrl Controller) Index(ctx iris.Context) {
	ctrl.View.Execute("index", map[string]interface{}{}, ctx.Request(), ctx.ResponseWriter())
}

// SwitchLocale switch locale
func (ctrl Controller) SwitchLocale(ctx iris.Context) {
	utils.SetCookie(http.Cookie{Name: "locale", Value: ctx.Request().URL.Query().Get("locale")}, &qor.Context{Request: ctx.Request(), Writer: ctx.ResponseWriter()})
	http.Redirect(ctx.ResponseWriter(), ctx.Request(), ctx.Request().Referer(), http.StatusSeeOther)
}
