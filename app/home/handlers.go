package home

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/kataras/iris/v12"
	"github.com/qor/render"
)

// Controller home controller
type Controller struct {
	View *render.Render
}

// Index home index page
func (ctrl Controller) Index(ctx iris.Context) {
	if err := ctx.View("index.html"); err != nil {
		color.Red(fmt.Sprintf("Home Index View error: %v\n", err))
	}
}
