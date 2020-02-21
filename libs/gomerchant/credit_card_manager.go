package gomerchant

// CreditCardManager interface
type CreditCardManager interface {
	CreateCreditCard(creditCardParams CreateCreditCardParams) (CreditCardResponse, error)
	GetCreditCard(creditCardParams GetCreditCardParams) (GetCreditCardResponse, error)
	ListCreditCards(listCreditCardsParams ListCreditCardsParams) (ListCreditCardsResponse, error)
	DeleteCreditCard(deleteCreditCardParams DeleteCreditCardParams) (DeleteCreditCardResponse, error)
}

// CreateCreditCard Params
type CreateCreditCardParams struct {
	CustomerID string
	CreditCard *CreditCard
}

type CreditCardResponse struct {
	CustomerID   string
	CreditCardID string
	Params
}

// Get Credit Cards Params
type GetCreditCardParams struct {
	CustomerID   string
	CreditCardID string
}

type GetCreditCardResponse struct {
	CreditCard *CustomerCreditCard
	Params
}

// Delete Credit Cards Params
type DeleteCreditCardParams struct {
	CustomerID   string
	CreditCardID string
}

type DeleteCreditCardResponse struct {
	Params
}

// List Credit Cards Params
type ListCreditCardsParams struct {
	CustomerID string
}

type ListCreditCardsResponse struct {
	CreditCards []*CustomerCreditCard
	Params
}

// CustomerCreditCard CustomerCard defination
type CustomerCreditCard struct {
	CustomerID   string
	CustomerName string
	CreditCardID string
	MaskedNumber string
	ExpMonth     uint
	ExpYear      uint
	Brand        string
	Params
}
