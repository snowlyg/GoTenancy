package products

import (
	"GoTenancy/libs/l10n"
	"github.com/jinzhu/gorm"
)

type Collection struct {
	gorm.Model
	Name string
	l10n.LocaleCreatable
}
