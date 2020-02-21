package gomerchant

import "errors"

var (
	// Mostly copied from Stripe, might need to be modified when integrating other payment gateways.
	ErrInvalidNumber      = errors.New("gomerchant: the card number is not a valid credit card number.")
	ErrInvalidExpiryMonth = errors.New("gomerchant: the card's expiration month is invalid.")
	ErrInvalidExpiryYear  = errors.New("gomerchant: the card's expiration year is invalid.")
	ErrInvalidCVC         = errors.New("gomerchant: the card's security code is invalid.")
	ErrIncorrectNumber    = errors.New("gomerchant: the card number is incorrect.")
	ErrExpiredCard        = errors.New("gomerchant: the card has expired.")
	ErrIncorrectCVC       = errors.New("gomerchant: the card's security code is incorrect.")
	ErrIncorrectZip       = errors.New("gomerchant: the card's zip code failed validation.")
	ErrCardDeclined       = errors.New("gomerchant: the card was declined.")
	ErrMissing            = errors.New("gomerchant: there is no card on a customer that is being charged.")
	ErrProcessingError    = errors.New("gomerchant: an error occurred while processing the card.")
)
