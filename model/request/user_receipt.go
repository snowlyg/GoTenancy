package request

type CreateReceipt struct {
	ReceiptType      int    `json:"receiptType" binding:"required,oneof=1 2"`
	ReceiptTitle     string `json:"receiptTitle" binding:"required"`
	ReceiptTitleType int    `json:"receiptTitleType" binding:"required,oneof=1 2"`
	DutyGaragraph    string `json:"dutyGaragraph" binding:"required"`
	Email            string `json:"email" binding:"required,email"`
	BankName         string `json:"bankName" binding:"required"`
	BankCode         string `json:"bankCode" binding:"required"`
	Address          string `json:"address" binding:"required"`
	Tel              string `json:"tel" binding:"required"`
	IsDefault        int    `json:"isDefault"`
}

type UpdateReceipt struct {
	Id               uint   `json:"id" form:"id" binding:"required,gt=0"`
	ReceiptType      int    `json:"receiptType" binding:"required,oneof=1 2"`
	ReceiptTitle     string `json:"receiptTitle" binding:"required"`
	ReceiptTitleType int    `json:"receiptTitleType" binding:"required,oneof=1 2"`
	DutyGaragraph    string `json:"dutyGaragraph" binding:"required"`
	Email            string `json:"email" binding:"required,email"`
	BankName         string `json:"bankName" binding:"required"`
	BankCode         string `json:"bankCode" binding:"required"`
	Address          string `json:"address" binding:"required"`
	Tel              string `json:"tel" binding:"required"`
	IsDefault        int    `json:"isDefault"`
}
