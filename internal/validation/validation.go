package validation

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateStruct validates any struct.
func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

// // RegisterCustomValidator registers a custom validator function.
// func RegisterCustomValidator(tag string, fn validator.Func) error {
// 	return validate.RegisterValidation(tag, fn)
// }