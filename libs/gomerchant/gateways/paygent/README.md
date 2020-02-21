# Paygent

Paygent Golang SDK

## Usage

```go
// Initalize
import "GoTenancy/libs/gomerchant/gateways/paygent"
Paygent = paygent.New(&paygent.Config{
  MerchantID:      "PaygentMerchantID",
  ConnectID:       "PaygentConnectID",
  ConnectPassword: "PaygentConnectPassword",
  ClientFilePath:  "PaygentClientFilePath",
  CertPassword:    "CertPassword",
  CAFilePath:      "CAFilePath",
  TelegramVersion: "1.0",
  ProductionMode:  false, // production or sandbox mode
})

// Store Credit Card
Paygent.CreateCreditCard(gomerchant.CreateCreditCardParams{
  CustomerID: "customer_id",
  CreditCard: &gomerchant.CreditCard{
    Name:     "holder name",
    Number:   "3580876521284076",
    ExpMonth: 10,
    ExpYear:  2017,
  },
})

// Get Credit Card
Paygent.GetCreditCard(gomerchant.GetCreditCardParams{CustomerID: "customer_id", CreditCardID: "3580876521284076"})

// Delete Stored Credit Card
Paygent.DeleteCreditCard(gomerchant.DeleteCreditCardParams{CustomerID: "customer_id", CreditCardID: "3580876521284076"})

// List Stored Credit Cards
Paygent.ListCreditCards(gomerchant.ListCreditCardsParams{CustomerID: "customer_id"})

// Take Auth
Paygent.Authorize(100, gomerchant.AuthorizeParams{
  Currency: "JPY",
  OrderID:  "order_id",
  PaymentMethod: &gomerchant.PaymentMethod{
    CreditCard: &gomerchant.CreditCard{
      Name:     "holder name",
      Number:   "3580876521284076",
      ExpMonth: 2,
      ExpYear:  2017,
    },
  },
})

// Take Auth with stored credit card
Paygent.Authorize(100, gomerchant.AuthorizeParams{
  Currency: "JPY",
  OrderID:  "order_id",
  PaymentMethod: &gomerchant.PaymentMethod{
    SavedCreditCard: &gomerchant.SavedCreditCard{
      CustomerID:   "customer id",
      CreditCardID: "stored card id",
    },
  },
})

// Capture
Paygent.Capture("payment_id from paygent", gomerchant.CaptureParams{})

// Refund Auth, 100 is the refuned amount
refundResponse, err := Paygent.Refund("payment id from paygent", 100, gomerchant.RefundParams{})
// after refund, paygent will return a new transaction id, get it from response
refundResponse.TransactionID

// Refund & Capture
refundResponse, err := Paygent.Refund("payment id from paygent", 100, gomerchant.RefundParams{Captured: true})

// Void Auth
Paygent.Void("payment id from paygent", gomerchant.VoidParams{})

// Void Captured Transaction
Paygent.Void("payment id from paygent", gomerchant.VoidParams{Captured: true})

// Query payment
transaction, err := Paygent.Query("payment id from paygent")

type Transaction struct {
	ID        string      // transaction id
	Amount    int         // payment amount
	Currency  string      // currency
	Captured  bool        // authorized and captured
	Paid      bool        // authorized OR captured
	Cancelled bool        // cancelled
	Status    string      // status code from paygent
	CreatedAt *time.Time  // payment created time
	Params                // extra params
}
```

## 3D Mode (SecureCode Mode)

```go
authorizeResult, err := Paygent.SecureCodeAuthorize(100,
  paygent.SecureCodeParams{
    UserAgent: "User-Agent	Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12) AppleWebKit/602.3.12 (KHTML, like Gecko) Version/10.0.2 Safari/602.3.12",
    TermURL:    "http://getqor.com/order/return",
    HttpAccept: "http",
  },
  gomerchant.AuthorizeParams{
    OrderID:  "order id",
    PaymentMethod: &gomerchant.PaymentMethod{
      CreditCard: &gomerchant.CreditCard{
        Name:     "holder name",
        Number:   "3580876521284076",
        ExpMonth: 10,
        ExpYear:  2019,
      },
    },
})

// In your controller
if authorizeResult.HandleRequest {
  if err := authorizeResult.RequestHandler(writer, request, gomerchant.Params{}); err == nil {
    return
  }
}

// In return controller (http://getqor.com/order/return)
var params gomerchant.CompleteAuthorizeParams
params.Set("request", request)
Paygent.CompleteAuthorize("payment id from paygent when get auth", params)
```

## Advanced Mode

```go
// Construct API request by yourself
Paygent.Request(telegramKind, gomerchant.Params{})

// For example
paygent.Request("094", gomerchant.Params{"payment_id": "payment id from paygent"})
```

## License

Released under the [MIT License](http://opensource.org/licenses/MIT).
