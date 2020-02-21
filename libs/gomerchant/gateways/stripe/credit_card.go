package stripe

import (
	"fmt"

	"GoTenancy/libs/gomerchant"
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/card"
)

func (*Stripe) CreateCreditCard(creditCardParams gomerchant.CreateCreditCardParams) (gomerchant.CreditCardResponse, error) {
	var (
		expMonth = fmt.Sprint(creditCardParams.CreditCard.ExpMonth)
		expYear  = fmt.Sprint(creditCardParams.CreditCard.ExpYear)
	)

	c, err := card.New(&stripe.CardParams{
		Customer: &creditCardParams.CustomerID,
		Name:     &creditCardParams.CreditCard.Name,
		Number:   &creditCardParams.CreditCard.Number,
		ExpMonth: &expMonth,
		ExpYear:  &expYear,
		CVC:      &creditCardParams.CreditCard.CVC,
	})

	resp := gomerchant.CreditCardResponse{CreditCardID: c.ID}

	if c.Customer != nil {
		resp.CustomerID = c.Customer.ID
	}

	return resp, err
}

func (*Stripe) GetCreditCard(creditCardParams gomerchant.GetCreditCardParams) (gomerchant.GetCreditCardResponse, error) {
	c, err := card.Get(creditCardParams.CreditCardID, &stripe.CardParams{Customer: &creditCardParams.CustomerID})

	resp := gomerchant.GetCreditCardResponse{
		CreditCard: &gomerchant.CustomerCreditCard{
			CustomerName: c.Name,
			CreditCardID: c.ID,
			MaskedNumber: fmt.Sprint(c.Last4),
			ExpMonth:     uint(c.ExpMonth),
			ExpYear:      uint(c.ExpYear),
			Brand:        string(c.Brand),
		},
	}

	if c.Customer != nil {
		resp.CreditCard.CustomerID = c.Customer.ID
	}

	return resp, err
}

func (*Stripe) ListCreditCards(listCreditCardsParams gomerchant.ListCreditCardsParams) (gomerchant.ListCreditCardsResponse, error) {
	iter := card.List(&stripe.CardListParams{Customer: &listCreditCardsParams.CustomerID})
	resp := gomerchant.ListCreditCardsResponse{}
	for iter.Next() {
		c := iter.Card()
		customerCreditCard := &gomerchant.CustomerCreditCard{
			CustomerName: c.Name,
			CreditCardID: c.ID,
			MaskedNumber: fmt.Sprint(c.Last4),
			ExpMonth:     uint(c.ExpMonth),
			ExpYear:      uint(c.ExpYear),
			Brand:        string(c.Brand),
		}

		if c.Customer != nil {
			customerCreditCard.CustomerID = c.Customer.ID
		}

		resp.CreditCards = append(resp.CreditCards, customerCreditCard)
	}
	return resp, iter.Err()
}

func (*Stripe) DeleteCreditCard(deleteCreditCardParams gomerchant.DeleteCreditCardParams) (gomerchant.DeleteCreditCardResponse, error) {
	_, err := card.Del(deleteCreditCardParams.CreditCardID, &stripe.CardParams{Customer: &deleteCreditCardParams.CustomerID})
	return gomerchant.DeleteCreditCardResponse{}, err
}
