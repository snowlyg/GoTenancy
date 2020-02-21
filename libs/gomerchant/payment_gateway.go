package gomerchant

import "net/http"

// PaymentGateway interface
type PaymentGateway interface {
	Authorize(amount uint64, params AuthorizeParams) (AuthorizeResponse, error)
	CompleteAuthorize(paymentID string, params CompleteAuthorizeParams) (CompleteAuthorizeResponse, error)
	Capture(transactionID string, params CaptureParams) (CaptureResponse, error)
	Refund(transactionID string, amount uint, params RefundParams) (RefundResponse, error)
	Void(transactionID string, params VoidParams) (VoidResponse, error)

	Query(transactionID string) (Transaction, error)
}

// AuthorizeParams authorize params
type AuthorizeParams struct {
	Amount          uint64
	Currency        string
	Customer        string
	Description     string
	OrderID         string
	BillingAddress  *Address
	ShippingAddress *Address
	PaymentMethod   *PaymentMethod
	Params
}

// AuthorizeResponse authorize response
type AuthorizeResponse struct {
	TransactionID  string
	HandleRequest  bool                                                   // need process request after authorize or not
	RequestHandler func(http.ResponseWriter, *http.Request, Params) error // process request
	Params
}

// CompleteAuthorizeParams complete authorize params
type CompleteAuthorizeParams struct {
	Params
}

// CompleteAuthorizeResponse complete authorize response
type CompleteAuthorizeResponse struct {
	Params
}

// CaptureParams capture params
type CaptureParams struct {
	Params
}

// CaptureResponse capture response
type CaptureResponse struct {
	TransactionID string
	Params
}

// RefundParams refund params
type RefundParams struct {
	Captured bool
	Params
}

// RefundResponse refund response
type RefundResponse struct {
	TransactionID string
	Params
}

// VoidParams void params
type VoidParams struct {
	Captured bool
	Params
}

// VoidResponse void response
type VoidResponse struct {
	TransactionID string
	Params
}
