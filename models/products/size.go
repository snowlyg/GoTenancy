package products

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	"github.com/qor/l10n"
	"github.com/qor/sorting"
	"github.com/qor/validations"
)

type Size struct {
	gorm.Model
	l10n.Locale
	sorting.Sorting
	Name string
	Code string `l10n:"sync"`
}

func (size Size) Validate(db *gorm.DB) {
	if strings.TrimSpace(size.Name) == "" {
		if err := db.AddError(validations.NewError(size, "Name", "Name can not be empty")); err != nil {
			color.Red(fmt.Sprintf("db.AddError error: %v", err))
		}
	}

	if strings.TrimSpace(size.Code) == "" {
		if err := db.AddError(validations.NewError(size, "Code", "Code can not be empty")); err != nil {
			color.Red(fmt.Sprintf("db.AddError error: %v", err))
		}
	}
}
