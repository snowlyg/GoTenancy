package request

type CreateReceipt struct {
	ReceiptType      int    `json:"receiptType" validate:"required,oneof=1 2"`
	ReceiptTitle     string `json:"receiptTitle" validate:"required"`
	ReceiptTitleType int    `json:"receiptTitleType" validate:"required,oneof=1 2"`
	DutyGaragraph    string `json:"dutyGaragraph" validate:"required"`
	Email            string `json:"email" validate:"required,email"`
	BankName         string `json:"bankName" validate:"required"`
	BankCode         string `json:"bankCode" validate:"required"`
	Address          string `json:"address" validate:"required"`
	Tel              string `json:"tel" validate:"required"`
	IsDefault        bool   `json:"isDefault"`
}

type UpdateReceipt struct {
	Id               uint   `json:"id" form:"id" validate:"required,gt=0"`
	ReceiptType      int    `json:"receiptType" validate:"required,oneof=1 2"`
	ReceiptTitle     string `json:"receiptTitle" validate:"required"`
	ReceiptTitleType int    `json:"receiptTitleType" validate:"required,oneof=1 2"`
	DutyGaragraph    string `json:"dutyGaragraph" validate:"required"`
	Email            string `json:"email" validate:"required,email"`
	BankName         string `json:"bankName" validate:"required"`
	BankCode         string `json:"bankCode" validate:"required"`
	Address          string `json:"address" validate:"required"`
	Tel              string `json:"tel" validate:"required"`
	IsDefault        bool   `json:"isDefault"`
}
