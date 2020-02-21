package gomerchant

import (
	"regexp"
	"strconv"
)

var Brands = map[string]*regexp.Regexp{
	`visa`:               regexp.MustCompile(`^4\d{12}(\d{3})?$`),
	`master`:             regexp.MustCompile(`^(5[1-5]\d{4}|677189)\d{10}$`),
	`discover`:           regexp.MustCompile(`^(6011|65\d{2}|64[4-9]\d)\d{12}|(62\d{14})$`),
	`american_express`:   regexp.MustCompile(`^3[47]\d{13}$`),
	`diners_club`:        regexp.MustCompile(`^3(0[0-5]|[68]\d)\d{11}$`),
	`jcb`:                regexp.MustCompile(`^35(27|28|29|[3-8]\d)\d{12}$`),
	`switch`:             regexp.MustCompile(`^6759\d{12}(\d{2,3})?$`),
	`solo`:               regexp.MustCompile(`^6767\d{12}(\d{2,3})?$`),
	`dankort`:            regexp.MustCompile(`^5019\d{12}$`),
	`maestro`:            regexp.MustCompile(`^(5[06-8]|6\d)\d{10,17}$`),
	`forbrugsforeningen`: regexp.MustCompile(`^600722\d{10}$`),
	`laser`:              regexp.MustCompile(`^(6304|6706|6771|6709)\d{8}(\d{4}|\d{6,7})?$`),
}

type SavedCreditCard struct {
	CustomerID   string
	CreditCardID string
}

type CreditCard struct {
	Name     string
	Number   string
	ExpMonth uint
	ExpYear  uint
	CVC      string
}

func (creditCard CreditCard) Brand() string {
	for name, match := range Brands {
		if match.MatchString(creditCard.Number) {
			return name
		}
	}
	return ""
}

// https://en.wikipedia.org/wiki/Luhn_algorithm
func (creditCard CreditCard) ValidNumber() bool {
	// number length >= 12
	if len(creditCard.Number) < 12 {
		return false
	}

	var number int
	if n, err := strconv.Atoi(creditCard.Number); err == nil {
		number = n
	} else {
		// should be digits
		return false
	}

	checkNumber := number % 10
	number = number / 10

	for i := 0; number > 0; i++ {
		cur := number % 10

		if i%2 == 0 { // even
			cur = cur * 2
			if cur > 9 {
				cur = cur%10 + cur/10
			}
		}

		checkNumber += cur
		number = number / 10
	}

	return checkNumber%10 == 0
}
