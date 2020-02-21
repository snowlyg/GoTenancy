// Package stripe implements GoMerchant payment gateway for Stripe.
package stripe

import (
	"fmt"
	"time"

	"GoTenancy/libs/gomerchant"
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/refund"
)

// Stripe implements gomerchant.PaymetGateway interface.
type Stripe struct {
	Config *Config
}

var _ gomerchant.PaymentGateway = &Stripe{}

// Config stripe config
type Config struct {
	Key string
}

// New creates Stripe struct.
func New(config *Config) *Stripe {
	stripe.Key = config.Key

	return &Stripe{
		Config: config,
	}
}

var capture bool = false

func (*Stripe) Authorize(amount uint64, params gomerchant.AuthorizeParams) (gomerchant.AuthorizeResponse, error) {
	int64Amount := int64(amount)
	chargeParams := &stripe.ChargeParams{
		Amount:      &int64Amount,
		Currency:    &params.Currency,
		Description: &params.Description,
		Capture:     &capture,
	}
	chargeParams.AddMetadata("order_id", params.OrderID)

	if params.PaymentMethod != nil {
		if params.PaymentMethod.CreditCard != nil {
			chargeParams.SetSource(toStripeCC(params.Customer, params.PaymentMethod.CreditCard, params.BillingAddress))
		}
		if params.PaymentMethod.SavedCreditCard != nil {
			if len(params.PaymentMethod.SavedCreditCard.CustomerID) > 0 {
				chargeParams.Customer = &params.PaymentMethod.SavedCreditCard.CustomerID
			}
			chargeParams.SetSource(params.PaymentMethod.SavedCreditCard.CreditCardID)
		}
	}

	charge, err := charge.New(chargeParams)
	if charge != nil {
		return gomerchant.AuthorizeResponse{TransactionID: charge.ID}, err
	}
	return gomerchant.AuthorizeResponse{}, err
}

func (*Stripe) CompleteAuthorize(paymentID string, params gomerchant.CompleteAuthorizeParams) (gomerchant.CompleteAuthorizeResponse, error) {
	return gomerchant.CompleteAuthorizeResponse{}, nil
}

func (*Stripe) Capture(transactionID string, params gomerchant.CaptureParams) (gomerchant.CaptureResponse, error) {
	_, err := charge.Capture(transactionID, nil)
	return gomerchant.CaptureResponse{TransactionID: transactionID}, err
}

func (s *Stripe) Refund(transactionID string, amount uint, params gomerchant.RefundParams) (gomerchant.RefundResponse, error) {
	transaction, err := s.Query(transactionID)

	if err == nil {
		if transaction.Captured {
			int64Amount := int64(amount)
			_, err = refund.New(&stripe.RefundParams{
				Charge: &transactionID,
				Amount: &int64Amount,
			})
		} else {
			int64Amount := int64(transaction.Amount - int(amount))
			_, err = charge.Capture(transactionID, &stripe.CaptureParams{
				Amount: &int64Amount,
			})
		}
	}

	return gomerchant.RefundResponse{TransactionID: transactionID}, err
}

func (*Stripe) Void(transactionID string, params gomerchant.VoidParams) (gomerchant.VoidResponse, error) {
	refundParams := &stripe.RefundParams{
		Charge: &transactionID,
	}
	_, err := refund.New(refundParams)
	return gomerchant.VoidResponse{TransactionID: transactionID}, err
}

func (*Stripe) Query(transactionID string) (gomerchant.Transaction, error) {
	c, err := charge.Get(transactionID, nil)
	created := time.Unix(c.Created, 0)
	transaction := gomerchant.Transaction{
		ID:        c.ID,
		Amount:    int(c.Amount - c.AmountRefunded),
		Currency:  string(c.Currency),
		Captured:  c.Captured,
		Paid:      c.Paid,
		Cancelled: c.Refunded,
		Status:    c.Status,
		CreatedAt: &created,
	}

	if transaction.Cancelled {
		transaction.Paid = false
		transaction.Captured = false
	}

	return transaction, err
}

func toStripeCC(customer string, cc *gomerchant.CreditCard, billingAddress *gomerchant.Address) *stripe.CardParams {
	var (
		expMonth = fmt.Sprint(cc.ExpMonth)
		expYear  = fmt.Sprint(cc.ExpYear)
	)
	cm := stripe.CardParams{
		Customer: &customer,
		Name:     &cc.Name,
		Number:   &cc.Number,
		ExpMonth: &expMonth,
		ExpYear:  &expYear,
		CVC:      &cc.CVC,
	}

	if billingAddress != nil {
		cm.AddressLine1 = &billingAddress.Address1
		cm.AddressLine1 = &billingAddress.Address2
		cm.AddressCity = &billingAddress.City
		cm.AddressState = &billingAddress.State
		cm.AddressZip = &billingAddress.ZIP
		cm.AddressCountry = &billingAddress.Country
	}

	return &cm
}
