package utils

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/snowlyg/go-tenancy/g"
)

func RegisterValidation(validate *validator.Validate) {
	validate.RegisterCustomTypeFunc(ValidateValuer, sql.NullString{}, sql.NullInt64{}, sql.NullBool{}, sql.NullFloat64{})
	validate.RegisterValidation("dev-required", ValidateDevRequired)
}

func ValidateValuer(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(driver.Valuer); ok {
		val, err := valuer.Value()
		if err == nil {
			return val
		}
	}
	return nil
}

func ValidateDevRequired(fl validator.FieldLevel) bool {
	if g.TENANCY_CONFIG.System.Env == "dev" {
		return true
	}
	return fl.Field().String() != ""
}

// Verify 校验方法
func Verify(err error) error {
	if err == nil {
		return nil
	}
	if errs, ok := err.(validator.ValidationErrors); ok {
		// Wrap the errors with JSON format, the underline library returns the errors as interface.
		validationErrors := wrapValidationErrors(errs)

		return validationErrors
	}
	return nil
}

func wrapValidationErrors(errs validator.ValidationErrors) error {
	var validationErrors []string
	for _, validationErr := range errs {
		validationErrors = append(validationErrors, validationErr.Error())
	}

	return errors.New(strings.Join(validationErrors, ";"))
}
