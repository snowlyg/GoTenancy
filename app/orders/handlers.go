package orders

import (
	"fmt"
	"net/http"
	"strconv"

	"GoTenancy/config"
	amazonpay "GoTenancy/libs/amazon-pay-sdk-go"
	"GoTenancy/libs/gomerchant"
	qorrender "GoTenancy/libs/render"
	"GoTenancy/libs/responder"
	"GoTenancy/libs/session/manager"
	"GoTenancy/models/orders"
	"GoTenancy/utils"
	"github.com/gorilla/schema"
	"github.com/kataras/iris/v12"
)

// Controller products controller
type Controller struct {
	View *qorrender.Render
}

var decoder = schema.NewDecoder()

// Cart shopping cart
func (ctrl Controller) Cart(ctx iris.Context) {
	order := getCurrentOrderWithItems(ctx)
	ctrl.View.Execute("cart", map[string]interface{}{"Order": order}, ctx.Request(), ctx.ResponseWriter())
}

// Checkout checkout shopping cart
func (ctrl Controller) Checkout(ctx iris.Context) {
	hasAmazon := ctx.Request().URL.Query().Get("access_token")
	order := getCurrentOrderWithItems(ctx)
	ctrl.View.Execute("checkout", map[string]interface{}{"Order": order, "HasAmazon": hasAmazon}, ctx.Request(), ctx.ResponseWriter())
}

// Complete complete shopping cart
func (ctrl Controller) Complete(ctx iris.Context) {
	ctx.Request().ParseForm()

	order := getCurrentOrder(ctx)
	if order.AmazonOrderReferenceID = ctx.Request().Form.Get("amazon_order_reference_id"); order.AmazonOrderReferenceID != "" {
		order.AmazonAddressAccessToken = ctx.Request().Form.Get("amazon_address_access_token")
		tx := utils.GetDB(ctx.Request())
		err := orders.OrderState.Trigger("checkout", order, tx, "")

		if err == nil {
			tx.Save(order)
			http.Redirect(ctx.ResponseWriter(), ctx.Request(), "/cart/success", http.StatusFound)
			return
		}
		utils.AddFlashMessage(ctx.ResponseWriter(), ctx.Request(), err.Error(), "error")
	} else {
		utils.AddFlashMessage(ctx.ResponseWriter(), ctx.Request(), "Order Reference ID not Found", "error")
	}

	http.Redirect(ctx.ResponseWriter(), ctx.Request(), "/cart", http.StatusFound)
}

// CompleteCreditCard complete shopping cart with credit card
func (ctrl Controller) CompleteCreditCard(ctx iris.Context) {
	ctx.Request().ParseForm()

	order := getCurrentOrder(ctx)

	expMonth, _ := strconv.Atoi(ctx.Request().Form.Get("exp_month"))
	expYear, _ := strconv.Atoi(ctx.Request().Form.Get("exp_year"))

	creditCard := gomerchant.CreditCard{
		Name:     ctx.Request().Form.Get("name"),
		Number:   ctx.Request().Form.Get("creditcard"),
		CVC:      ctx.Request().Form.Get("cvv"),
		ExpYear:  uint(expYear),
		ExpMonth: uint(expMonth),
	}

	if creditCard.ValidNumber() {
		// TODO integrate with https://GoTenancy/libs/gomerchant to handle those information
		tx := utils.GetDB(ctx.Request())
		err := orders.OrderState.Trigger("checkout", order, tx, "")

		if err == nil {
			tx.Save(order)
			http.Redirect(ctx.ResponseWriter(), ctx.Request(), "/cart/success", http.StatusFound)
			return
		}
	}

	utils.AddFlashMessage(ctx.ResponseWriter(), ctx.Request(), "Invalid Credit Card", "error")
	http.Redirect(ctx.ResponseWriter(), ctx.Request(), "/cart", http.StatusFound)
}

// CheckoutSuccess checkout success page
func (ctrl Controller) CheckoutSuccess(ctx iris.Context) {
	ctrl.View.Execute("success", map[string]interface{}{}, ctx.Request(), ctx.ResponseWriter())
}

type updateCartInput struct {
	SizeVariationID  uint `schema:"size_variation_id"`
	Quantity         uint `schema:"quantity"`
	ProductID        uint `schema:"product_id"`
	ColorVariationID uint `schema:"color_variation_id"`
}

// UpdateCart update shopping cart
func (ctrl Controller) UpdateCart(ctx iris.Context) {
	var (
		input updateCartInput
		tx    = utils.GetDB(ctx.Request())
	)

	ctx.Request().ParseForm()
	decoder.Decode(&input, ctx.Request().PostForm)

	order := getCurrentOrder(ctx)

	if input.Quantity > 0 {
		tx.Where(&orders.OrderItem{OrderID: order.ID, SizeVariationID: input.SizeVariationID}).
			Assign(&orders.OrderItem{Quantity: input.Quantity}).
			FirstOrCreate(&orders.OrderItem{})
	} else {
		tx.Where(&orders.OrderItem{OrderID: order.ID, SizeVariationID: input.SizeVariationID}).
			Delete(&orders.OrderItem{})
	}

	responder.With("html", func() {
		http.Redirect(ctx.ResponseWriter(), ctx.Request(), "/cart", http.StatusFound)
	}).With([]string{"json", "xml"}, func() {
		config.Render.JSON(ctx.ResponseWriter(), http.StatusOK, map[string]string{"status": "ok"})
	}).Respond(ctx.Request())
}

// AmazonCallback amazon callback
func (ctrl Controller) AmazonCallback(ctx iris.Context) {
	ipn, ok := amazonpay.VerifyIPNRequest(ctx.Request())
	fmt.Printf("%#v\n", ipn)
	fmt.Printf("%#v\n", ok)
}

func getCurrentOrder(ctx iris.Context) *orders.Order {
	var (
		order       = orders.Order{}
		cartID      = manager.SessionManager.Get(ctx.Request(), "cart_id")
		currentUser = utils.GetCurrentUser(ctx.Request())
		tx          = utils.GetDB(ctx.Request())
	)

	if cartID != "" {
		if currentUser != nil && !tx.NewRecord(currentUser) {
			if !tx.First(&order, "id = ? AND (user_id = ? OR user_id IS NULL)", cartID, currentUser.ID).RecordNotFound() && order.UserID == nil {
				tx.Model(&order).Update("UserID", currentUser.ID)
			}
		} else {
			tx.First(&order, "id = ? AND user_id IS NULL", cartID)
		}
	}

	// only create new shopping cart if updating
	if tx.NewRecord(order) || !order.IsCart() {
		order = orders.Order{}
		if ctx.Request().Method != "GET" {
			if currentUser != nil && !tx.NewRecord(currentUser) {
				order.UserID = &currentUser.ID
			}

			tx.Create(&order)
		}
	}

	manager.SessionManager.Add(ctx.ResponseWriter(), ctx.Request(), "cart_id", order.ID)

	return &order
}

func getCurrentOrderWithItems(ctx iris.Context) *orders.Order {
	order := getCurrentOrder(ctx)
	if tx := utils.GetDB(ctx.Request()); !tx.NewRecord(order) {
		tx.Model(order).Association("OrderItems").Find(&order.OrderItems)
	}
	return order
}
