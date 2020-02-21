# Amazon Pay Go SDK

Amazon Pay SDK for Golang

## Install

```sh
go get GoTenancy/libs/amazon-pay-sdk-go
```

## Usage

```go
import amazonpay "GoTenancy/libs/amazon-pay-sdk-go"

func main() {
  var client = amazonpay.New(&amazonpay.Config{
    MerchantID: "my merchant id",
    AccessKey:  "my access key",
    SecretKey:  "my secret key",
    Sandbox:    true,
    Region:     "jp",
  })

  // Set order details such as order amount and explanation in Order Reference
  client.SetOrderReferenceDetails(orderReferenceID, amazonpay.OrderReferenceAttributes) (amazonpay.SetOrderReferenceDetailsResult, error)

  // confirm order details
  client.ConfirmOrderReference(orderReferenceID) error

  // Returns the details and current state of the Order Reference object
  client.GetOrderReferenceDetails(orderReferenceID, addressToken) (amazonpay.GetOrderReferenceDetailsResponse, error)

  // Process secures the funds specified for the payment method stored in the Order Reference
  client.Authorize(orderReferenceID, authorizationReferenceID, amount, amazonpay.AuthorizeInput) (amazonpay.AuthorizeResponse, error)

  // Returns the total authorized amount for authorization status and authorization
  client.GetAuthorizationDetails(authorizationID) (amazonpay.GetAuthorizationDetailsResponse, error)

  // CloseAuthorization Close authorization
  client.CloseAuthorization(authorizationID, closureReason) error

  // Request funds from the authorized payment method
  client.Capture(authorizationID, captureReferenceID, captureAmount, amazonpay.CaptureInput) (amazonpay.CaptureResponse, error)

  // Returns the detailed sales request status and the total amount refunded by sales request
  client.GetCaptureDetails(captureID) (amazonpay.GetCaptureDetailsResponse, error)

  // Complete order reference and will not be able to generate a new authorization from this Order Reference
  client.CloseOrderReference(orderReferenceID, closureReason) error

  // CancelOrderReference Cancels a previously confirmed order reference
  client.CancelOrderReference(orderReferenceID, reason) error

  // Refund refund the funds requested
  client.Refund(captureID, refundReferenceID, refundAmount, amazonpay.RefundInput) (amazonpay.RefundResponse, error)

  // Get refund details
  client.GetRefundDetails(refundID) (amazonpay.GetRefundDetailsResponse, error)
}
```

## Verify IPN Notification

Verify ipn notification

```go
amazonpay.VerifyIPNRequest(req)
```

## Demo

We have deployed a demo that integrated Amazon Pay sandbox mode, visit it here: [https://demo.getqor.com](https://demo.getqor.com)

You can place orders with our Amazon sandbox account: `demo@getqor.com` / `qordemo`

After placed orders, you can manage them via [our admin interface]((https://demo.getqor.com/admin)), like `take auth`, `capture`, `refund` orders, the admin interface is generated with [QOR Admin](http://GoTenancy/libs/admin)

#### Source code of this demo

[https://GoTenancy/libs/qor-example](https://GoTenancy/libs/qor-example)

## License

Released under the [MIT License](http://opensource.org/licenses/MIT).
