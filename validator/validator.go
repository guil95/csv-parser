package validator

import (
	"strings"

	validatorplayground "github.com/go-playground/validator/v10"
)

func notBlank(fl validatorplayground.FieldLevel) bool {
	return strings.TrimSpace(fl.Field().String()) != ""
}

func Validate(v interface{}) error {
	validate := validatorplayground.New()
	_ = validate.RegisterValidation("not_blank", notBlank)

	return validate.Struct(v)
}
