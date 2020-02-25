package products

import (
	"github.com/qor/l10n"
	"github.com/jinzhu/gorm"
)

type Collection struct {
	gorm.Model
	Name string
	l10n.LocaleCreatable
}
