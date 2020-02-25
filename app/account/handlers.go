package account

import (
	"GoTenancy/models/orders"
	"GoTenancy/models/users"
	"GoTenancy/utils"
	"github.com/kataras/iris/v12"
	"github.com/qor/render"
)

// Controller products controller
type Controller struct {
	View *render.Render
}

// Profile profile show page
func (ctrl Controller) Profile(ctx iris.Context) {
	var (
		currentUser                     = utils.GetCurrentUser(ctx.Request())
		tx                              = utils.GetDB(ctx.Request())
		billingAddress, shippingAddress users.Address
	)

	// TODO refactor
	tx.Model(currentUser).Related(&currentUser.Addresses, "Addresses")
	tx.Model(currentUser).Related(&billingAddress, "DefaultBillingAddress")
	tx.Model(currentUser).Related(&shippingAddress, "DefaultShippingAddress")

	ctrl.View.Execute("profile", map[string]interface{}{
		"CurrentUser": currentUser, "DefaultBillingAddress": billingAddress, "DefaultShippingAddress": shippingAddress,
	}, ctx.Request(), ctx.ResponseWriter())
}

// Orders orders page
func (ctrl Controller) Orders(ctx iris.Context) {
	var (
		Orders      []orders.Order
		currentUser = utils.GetCurrentUser(ctx.Request())
		tx          = utils.GetDB(ctx.Request())
	)

	tx.Preload("OrderItems").Where("state <> ? AND state != ?", orders.DraftState, "").Where(&orders.Order{UserID: &currentUser.ID}).Find(&Orders)

	ctrl.View.Execute("orders", map[string]interface{}{"Orders": Orders}, ctx.Request(), ctx.ResponseWriter())
}

// Update update profile page
func (ctrl Controller) Update(ctx iris.Context) {
	// FIXME
}

// AddCredit add credit
func (ctrl Controller) AddCredit(ctx iris.Context) {
	// FIXME
}
