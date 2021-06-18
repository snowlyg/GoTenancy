package utils

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/snowlyg/go-tenancy/g"
)

func RegisterValidation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("dev-required", ValidateDevRequired)
	}
}

var ValidateDevRequired validator.Func = func(fl validator.FieldLevel) bool {
	if g.TENANCY_CONFIG.System.Env != "pro" {
		return true
	}
	return fl.Field().String() != ""
}
