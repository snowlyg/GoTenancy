package gomerchant

import "errors"

var ErrNotSupportedPaymentMethod = errors.New("not supported payment method")

type PaymentMethod struct {
	SavedCreditCard *SavedCreditCard
	CreditCard      *CreditCard
}
