package account

import (
	"github.com/kataras/iris/v12"
	"github.com/qor/render"
	"go-tenancy/models/users"
	"go-tenancy/utils"
)

// Controller Profile controller
type Controller struct {
	View *render.Render
}

// Profile profile show page
func (ctrl Controller) Profile(ctx iris.Context) {
	var (
		currentUser                     = utils.GetTenantCurrentUser(ctx.Request())
		tx                              = utils.GetDB(ctx.Request())
		billingAddress, shippingAddress users.Address
	)

	// TODO refactor
	tx.Model(currentUser).Related(&currentUser.Addresses, "Addresses")
	tx.Model(currentUser).Related(&billingAddress, "DefaultBillingAddress")
	tx.Model(currentUser).Related(&shippingAddress, "DefaultShippingAddress")

	_ = ctrl.View.Execute("profile", map[string]interface{}{
		"CurrentUser": currentUser, "DefaultBillingAddress": billingAddress, "DefaultShippingAddress": shippingAddress,
	}, ctx.Request(), ctx.ResponseWriter())
}

// Update update profile page
func (ctrl Controller) Update(ctx iris.Context) {
	// FIXME
}

// AddCredit add credit
func (ctrl Controller) AddCredit(ctx iris.Context) {
	// FIXME
}
