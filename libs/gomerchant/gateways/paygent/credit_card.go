package paygent

import (
	"encoding/csv"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"GoTenancy/libs/gomerchant"
)

func getValidTerm(creditCard *gomerchant.CreditCard) string {
	year := fmt.Sprint(creditCard.ExpYear)
	if len(year) >= 2 {
		year = year[len(year)-2:]
	} else {
		year = ""
	}
	return fmt.Sprintf("%02d", creditCard.ExpMonth) + year
}

var brandsMap = map[string]string{"visa": "V", "master": "M", "american_express": "X", "diners_club": "C", "jcb": "J"}

func (paygent *Paygent) CreateCreditCard(creditCardParams gomerchant.CreateCreditCardParams) (gomerchant.CreditCardResponse, error) {
	var (
		response   = gomerchant.CreditCardResponse{CustomerID: creditCardParams.CustomerID}
		creditCard = creditCardParams.CreditCard
		brand, _   = brandsMap[creditCard.Brand()]
	)

	results, err := paygent.Request("025", gomerchant.Params{
		"customer_id":     creditCardParams.CustomerID,
		"card_number":     creditCard.Number,
		"card_valid_term": getValidTerm(creditCard),
		"cardholder_name": creditCard.Name,
		"card_brand":      brand,
	}.IgnoreBlankFields())

	if err == nil {
		if customerCardID, ok := results.Get("customer_card_id"); ok {
			response.CreditCardID = fmt.Sprint(customerCardID)
		}
	}
	response.Params = results.Params

	return response, err
}

func (paygent *Paygent) GetCreditCard(getCreditCardParams gomerchant.GetCreditCardParams) (gomerchant.GetCreditCardResponse, error) {
	var response gomerchant.GetCreditCardResponse
	results, err := paygent.Request("027", gomerchant.Params{"customer_id": getCreditCardParams.CustomerID, "credit_card_id": getCreditCardParams.CreditCardID})

	if err == nil {
		cards, err := parseListCreditCardsResponse(&results)
		if len(cards) > 0 {
			response.CreditCard = cards[0]
		} else {
			err = errors.New("credit card not found")
		}
		return response, err
	}

	return response, err
}

func (paygent *Paygent) DeleteCreditCard(deleteCreditCardParams gomerchant.DeleteCreditCardParams) (gomerchant.DeleteCreditCardResponse, error) {
	var response = gomerchant.DeleteCreditCardResponse{}

	results, err := paygent.Request("026", gomerchant.Params{"customer_id": deleteCreditCardParams.CustomerID, "customer_card_id": deleteCreditCardParams.CreditCardID}.IgnoreBlankFields())
	response.Params = results.Params
	return response, err
}

func (paygent *Paygent) ListCreditCards(listCreditCardsParams gomerchant.ListCreditCardsParams) (gomerchant.ListCreditCardsResponse, error) {
	var response = gomerchant.ListCreditCardsResponse{}

	results, err := paygent.Request("027", gomerchant.Params{"customer_id": listCreditCardsParams.CustomerID})

	if err == nil {
		response.CreditCards, err = parseListCreditCardsResponse(&results)
	}

	if results.ResponseCode == "P026" {
		err = nil
	}

	return response, err
}

func parseListCreditCardsResponse(response *Response) (cards []*gomerchant.CustomerCreditCard, err error) {
	var headers []string

	for _, str := range strings.Split(response.RawBody, "\r\n") {
		if str != "" {
			row := csv.NewReader(strings.NewReader(str))
			if record, err := row.Read(); err == nil {
				if len(record) == 0 {
					return nil, errors.New("wrong format")
				}

				switch record[0] {
				case "1":
					// response information
					if len(record) != 4 {
						return nil, errors.New("wrong format")
					}

					response.Result = record[1]
					response.ResponseCode = record[2]
					response.ResponseDetail = record[3]
				case "2":
					// card header
					headers = record
				case "3":
					// card information
					params := gomerchant.Params{}
					for idx, value := range record {
						params[headers[idx]] = value
					}
					customerCard := &gomerchant.CustomerCreditCard{Params: params}

					if v, ok := params.Get("customer_id"); ok {
						customerCard.CustomerID = fmt.Sprint(v)
					}

					if v, ok := params.Get("customer_card_id"); ok {
						customerCard.CreditCardID = fmt.Sprint(v)
					}

					if v, ok := params.Get("cardholder_name"); ok {
						customerCard.CustomerName = fmt.Sprint(v)
					}

					if v, ok := params.Get("card_number"); ok {
						customerCard.MaskedNumber = fmt.Sprint(v)
					}

					if v, ok := params.Get("card_brand"); ok {
						for key, value := range brandsMap {
							if fmt.Sprint(v) == value {
								customerCard.Brand = key
							}
						}

						if customerCard.Brand == "" {
							customerCard.Brand = fmt.Sprint(v)
						}
					}

					if v, ok := params.Get("card_valid_term"); ok {
						if u, err := strconv.Atoi(fmt.Sprint(v)[0:2]); err == nil {
							customerCard.ExpMonth = uint(u)
						}
						if u, err := strconv.Atoi(fmt.Sprint(v)[2:4]); err == nil {
							customerCard.ExpYear = uint(time.Now().Year()/100*100 + u)
						}
					}

					cards = append(cards, customerCard)
				case "4":
					// card numbers
				}
			} else {
				return nil, err
			}
		}
	}

	if response.Result == "1" {
		if response.ResponseDetail != "" {
			err = errors.New(response.ResponseDetail)
		} else {
			err = errors.New("failed to process this request")
		}
	}

	return cards, err
}
