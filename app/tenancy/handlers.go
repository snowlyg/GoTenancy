package tenancy

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/kataras/iris/v12"
	"github.com/qor/render"
	"go-tenancy/models/tenancy"
	"go-tenancy/utils"
)

// Controller Tenants controller
type Controller struct {
	View *render.Render
}

// Index Tenants index page
func (ctrl Controller) Index(ctx iris.Context) {
	var (
		Tenants []tenancy.Tenant
		tx      = utils.GetDB(ctx.Request())
	)

	tx.Preload("Category").Find(&Tenants)

	if err := ctrl.View.Execute("index", map[string]interface{}{}, ctx.Request(), ctx.ResponseWriter()); err != nil {
		color.Red(fmt.Sprintf("View.Execute error %v\n", err))
	}
}

// Show product show page
func (ctrl Controller) Show(ctx iris.Context) {
	var (
		product tenancy.Tenant
		tx      = utils.GetDB(ctx.Request())
	)

	if tx.First(&product).RecordNotFound() {
		http.Redirect(ctx.ResponseWriter(), ctx.Request(), "/", http.StatusFound)
	}

	if err := ctrl.View.Execute("show", map[string]interface{}{}, ctx.Request(), ctx.ResponseWriter()); err != nil {
		color.Red(fmt.Sprintf("View.Execute error %v\n", err))
	}
}
