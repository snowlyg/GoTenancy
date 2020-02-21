package settings

import (
	"GoTenancy/libs/l10n"
	"GoTenancy/libs/location"
	"github.com/jinzhu/gorm"
)

type FeeSetting struct {
	ShippingFee     uint
	GiftWrappingFee uint
	CODFee          uint `gorm:"column:cod_fee"`
	TaxRate         int
}

type Setting struct {
	gorm.Model
	FeeSetting
	location.Location `location:"name:Company Address"`
	l10n.Locale
}
