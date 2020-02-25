package products

import (
	"fmt"
	"strings"

	"github.com/qor/l10n"
	"github.com/qor/sorting"
	"github.com/qor/validations"
	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	l10n.Locale
	sorting.Sorting
	Name string
	Code string

	Categories []Category
	CategoryID uint
}

func (category Category) Validate(db *gorm.DB) {
	if strings.TrimSpace(category.Name) == "" {
		if err := db.AddError(validations.NewError(category, "Name", "Name can not be empty")); err != nil {
			color.Red(fmt.Sprintf("db.AddError error: %v", err))
		}
	}
}

func (category Category) DefaultPath() string {
	if len(category.Code) > 0 {
		return fmt.Sprintf("/category/%s", category.Code)
	}
	return "/"
}
