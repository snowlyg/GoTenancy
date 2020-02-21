package tests

import (
	"fmt"
	"testing"
	"time"

	"GoTenancy/libs/gomerchant"
)

type TestSuite struct {
	CreditCardManager   gomerchant.CreditCardManager
	Gateway             gomerchant.PaymentGateway
	GetRandomCustomerID func() string
}

func (testSuite TestSuite) TestAll(t *testing.T) {
	testSuite.TestCreateCreditCard(t)
	testSuite.TestAuthorizeAndCapture(t)
	testSuite.TestAuthorizeAndCaptureWithSavedCreditCard(t)
	testSuite.TestRefund(t)
	testSuite.TestVoid(t)

	testSuite.TestListCreditCards(t)
	testSuite.TestListCreditCardsWithNoResult(t)
	testSuite.TestGetCreditCard(t)
	testSuite.TestDeleteCreditCard(t)
}

func (testSuite TestSuite) createSavedCreditCard() (gomerchant.CreditCardResponse, error) {
	return testSuite.CreditCardManager.CreateCreditCard(gomerchant.CreateCreditCardParams{
		CustomerID: testSuite.GetRandomCustomerID(),
		CreditCard: &gomerchant.CreditCard{
			Name:     "VISA",
			Number:   "4242424242424242",
			ExpMonth: 1,
			ExpYear:  uint(time.Now().Year() + 1),
		},
	})
}

func (testSuite TestSuite) TestCreateCreditCard(t *testing.T) {
	if result, err := testSuite.createSavedCreditCard(); err != nil || result.CreditCardID == "" {
		t.Error(err, result)
	}
}

func (testSuite TestSuite) TestAuthorizeAndCapture(t *testing.T) {
	authorizeResult, err := testSuite.Gateway.Authorize(100, gomerchant.AuthorizeParams{
		Currency: "JPY",
		OrderID:  fmt.Sprint(time.Now().Unix()),
		PaymentMethod: &gomerchant.PaymentMethod{
			CreditCard: &gomerchant.CreditCard{
				Name:     "VISA",
				Number:   "4242424242424242",
				ExpMonth: 1,
				ExpYear:  uint(time.Now().Year() + 1),
			},
		},
	})

	if err != nil || authorizeResult.TransactionID == "" {
		t.Error(err, authorizeResult)
	}

	captureResult, err := testSuite.Gateway.Capture(authorizeResult.TransactionID, gomerchant.CaptureParams{})

	if err != nil || captureResult.TransactionID == "" {
		t.Error(err, captureResult)
	}
}

func (testSuite TestSuite) TestAuthorizeAndCaptureWithSavedCreditCard(t *testing.T) {
	if savedCreditCard, err := testSuite.createSavedCreditCard(); err == nil {
		authorizeResult, err := testSuite.Gateway.Authorize(100, gomerchant.AuthorizeParams{
			Currency: "JPY",
			OrderID:  fmt.Sprint(time.Now().Unix()),
			PaymentMethod: &gomerchant.PaymentMethod{
				SavedCreditCard: &gomerchant.SavedCreditCard{
					CustomerID:   savedCreditCard.CustomerID,
					CreditCardID: savedCreditCard.CreditCardID,
				},
			},
		})

		if err != nil || authorizeResult.TransactionID == "" {
			t.Error(err, authorizeResult)
		}

		captureResult, err := testSuite.Gateway.Capture(authorizeResult.TransactionID, gomerchant.CaptureParams{})

		if err != nil || captureResult.TransactionID == "" {
			t.Error(err, captureResult)
		}
	}
}

func (testSuite TestSuite) createAuth() gomerchant.AuthorizeResponse {
	authorizeResponse, _ := testSuite.Gateway.Authorize(1000, gomerchant.AuthorizeParams{
		Currency: "JPY",
		OrderID:  fmt.Sprint(time.Now().Unix()),
		PaymentMethod: &gomerchant.PaymentMethod{
			CreditCard: &gomerchant.CreditCard{
				Name:     "VISA",
				Number:   "4242424242424242",
				ExpMonth: 1,
				ExpYear:  uint(time.Now().Year() + 1),
			},
		},
	})

	return authorizeResponse
}

func (testSuite TestSuite) TestRefund(t *testing.T) {
	// refund authorized transaction
	authorizeResponse := testSuite.createAuth()
	if refundResponse, err := testSuite.Gateway.Refund(authorizeResponse.TransactionID, 100, gomerchant.RefundParams{}); err == nil {
		if transaction, err := testSuite.Gateway.Query(refundResponse.TransactionID); err == nil {
			if !(transaction.Amount == 900 && transaction.Paid == true && transaction.Cancelled == false && transaction.CreatedAt != nil) { // &&transaction.Captured == false) {
				t.Errorf("transaction after refund authorized transaction is not correct, but got %#v", transaction)
			}
		} else {
			t.Errorf("no error should happen when query transaction, but got %v, %#v", err, transaction)
		}
	} else {
		t.Errorf("no error should happen when refund transaction, but got %v", err)
	}

	// refund authorized transaction, and capture it
	authorizeResponse = testSuite.createAuth()
	if refundResponse, err := testSuite.Gateway.Refund(authorizeResponse.TransactionID, 150, gomerchant.RefundParams{Captured: true}); err == nil {
		if transaction, err := testSuite.Gateway.Query(refundResponse.TransactionID); err == nil {
			if !(transaction.Amount == 850 && transaction.Paid == true && transaction.Captured == true && transaction.Cancelled == false && transaction.CreatedAt != nil) {
				t.Errorf("transaction after refund authorized transaction is not correct, but got %#v", transaction)
			}
		} else {
			t.Errorf("no error should happen when query transaction, but got %v, %#v", err, transaction)
		}
	} else {
		t.Errorf("no error should happen when refund transaction, but got %v", err)
	}

	// refund captured transaction
	authorizeResponse = testSuite.createAuth()
	captureResponse, _ := testSuite.Gateway.Capture(authorizeResponse.TransactionID, gomerchant.CaptureParams{})
	if refundResponse, err := testSuite.Gateway.Refund(captureResponse.TransactionID, 200, gomerchant.RefundParams{Captured: true}); err == nil {
		if transaction, err := testSuite.Gateway.Query(refundResponse.TransactionID); err == nil {
			if !(transaction.Amount == 800 && transaction.Paid == true && transaction.Captured == true && transaction.Cancelled == false && transaction.CreatedAt != nil) {
				t.Errorf("transaction after refund captured transaction is not correct, but got %#v", transaction)
			}
		} else {
			t.Errorf("no error should happen when query transaction, but got %v, %#v", err, transaction)
		}
	} else {
		t.Errorf("no error should happen when refund transaction, but got %v", err)
	}
}

func (testSuite TestSuite) TestVoid(t *testing.T) {
	// void authorized transaction
	authorizeResponse := testSuite.createAuth()
	if refundResponse, err := testSuite.Gateway.Void(authorizeResponse.TransactionID, gomerchant.VoidParams{}); err == nil {
		if transaction, err := testSuite.Gateway.Query(refundResponse.TransactionID); err == nil {
			if !(transaction.Paid == false && transaction.Captured == false && transaction.Cancelled == true && transaction.CreatedAt != nil) { // && transaction.Amount == 1000) {
				t.Errorf("transaction after refund auth is not correct, but got %#v", transaction)
			}
		} else {
			t.Errorf("no error should happen when query transaction, but got %v, %#v", err, transaction)
		}
	} else {
		t.Errorf("no error should happen when refund transaction, but got %v", err)
	}

	// void captured transaction
	authorizeResponse = testSuite.createAuth()
	captureResponse, _ := testSuite.Gateway.Capture(authorizeResponse.TransactionID, gomerchant.CaptureParams{})
	if refundResponse, err := testSuite.Gateway.Void(captureResponse.TransactionID, gomerchant.VoidParams{Captured: true}); err == nil {
		if transaction, err := testSuite.Gateway.Query(refundResponse.TransactionID); err == nil {
			if !(transaction.Paid == false && transaction.Captured == false && transaction.Cancelled == true && transaction.CreatedAt != nil) { // && transaction.Amount == 1000) {
				t.Errorf("transaction after refund captured is not correct, but got %#v", transaction)
			}
		} else {
			t.Errorf("no error should happen when query transaction, but got %v, %#v", err, transaction)
		}
	} else {
		t.Errorf("no error should happen when refund transaction, but got %v", err)
	}
}

func (testSuite TestSuite) TestListCreditCards(t *testing.T) {
	if response, err := testSuite.createSavedCreditCard(); err == nil {
		// create anotther credit card
		testSuite.CreditCardManager.CreateCreditCard(gomerchant.CreateCreditCardParams{
			CustomerID: response.CustomerID,
			CreditCard: &gomerchant.CreditCard{
				Name:     "VISA",
				Number:   "4242424242424242",
				ExpMonth: 1,
				ExpYear:  uint(time.Now().Year() + 1),
			},
		})

		if response, err := testSuite.CreditCardManager.ListCreditCards(gomerchant.ListCreditCardsParams{CustomerID: response.CustomerID}); err == nil {
			if len(response.CreditCards) != 2 {
				t.Errorf("Should found two saved credit cards, but got %v", response.CreditCards)
			}

			for _, creditCard := range response.CreditCards {
				if creditCard.MaskedNumber == "" || creditCard.ExpYear == 0 || creditCard.ExpMonth == 0 || creditCard.CustomerID == "" || creditCard.CreditCardID == "" {
					t.Errorf("Credit card's information should be correct, but got %v", creditCard)
				}
			}
		} else {
			t.Errorf("no error should happen when query saved credit cards, but got %v", err)
		}
	}
}

func (testSuite TestSuite) TestListCreditCardsWithNoResult(t *testing.T) {
	if response, err := testSuite.CreditCardManager.ListCreditCards(gomerchant.ListCreditCardsParams{CustomerID: testSuite.GetRandomCustomerID()}); err != nil {
		t.Errorf("should not return error, but got %v", err)
	} else if len(response.CreditCards) != 0 {
		t.Errorf("credit card's count should be zero")
	}
}

func (testSuite TestSuite) TestGetCreditCard(t *testing.T) {
	if response, err := testSuite.createSavedCreditCard(); err == nil {
		if response, err := testSuite.CreditCardManager.GetCreditCard(gomerchant.GetCreditCardParams{CustomerID: response.CustomerID, CreditCardID: response.CreditCardID}); err == nil {
			creditCard := response.CreditCard
			if creditCard == nil {
				t.Errorf("Should found saved credit cards, but got %v", response)
			} else if creditCard.Brand == "" || creditCard.MaskedNumber == "" || creditCard.ExpYear == 0 || creditCard.ExpMonth == 0 || creditCard.CustomerID == "" || creditCard.CustomerName == "" || creditCard.CreditCardID == "" {
				t.Errorf("Credit card's information should be correct, but got %v", creditCard)
			}
		} else {
			t.Errorf("no error should happen when query saved credit card, but got %v", err)
		}
	}
}

func (testSuite TestSuite) TestDeleteCreditCard(t *testing.T) {
	if response, err := testSuite.createSavedCreditCard(); err == nil {
		if _, err := testSuite.CreditCardManager.DeleteCreditCard(gomerchant.DeleteCreditCardParams{CustomerID: response.CustomerID, CreditCardID: response.CreditCardID}); err == nil {
			if response, err := testSuite.CreditCardManager.GetCreditCard(gomerchant.GetCreditCardParams{CustomerID: response.CustomerID, CreditCardID: response.CreditCardID}); err == nil {
				t.Errorf("Should failed to get credit card, but got %v", response)
			}
		} else {
			t.Errorf("no error should happen when delete saved credit card, but got %v", err)
		}
	}
}
