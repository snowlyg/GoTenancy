package pages

import (
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
