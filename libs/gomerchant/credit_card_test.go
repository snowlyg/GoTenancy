package gomerchant

import "testing"

func TestCreditCardLuhnAlgorithm(t *testing.T) {
	validNumbers := []string{"4111111111111111", "5431111111111111", "341111111111111", "6011601160116611", "5105105105105100", "5555555555554444", "4222222222222", "378282246310005", "371449635398431", "378734493671000", "38520000023237", "30569309025904", "6011111111111117", "6011000990139424", "3530111333300000", "3566002020360505"}

	for _, number := range validNumbers {
		creditCard := CreditCard{Number: number}
		if !creditCard.ValidNumber() {
			t.Errorf("%v should be valid", number)
		}
	}
}
