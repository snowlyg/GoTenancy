package pages

import (
	"github.com/kataras/iris/v12"
	"github.com/qor/render"
)

// Controller page controller
type Controller struct {
	View *render.Render
}

// Index page index page
func (ctrl Controller) Index(ctx iris.Context) {
	ctrl.View.Execute("index", map[string]interface{}{}, ctx.Request(), ctx.ResponseWriter())
}
