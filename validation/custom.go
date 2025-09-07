package validation

import (
	"regexp"
	"github.com/go-playground/validator/v10"
)

func init() {
	RegisterCustomValidator("e164", validateE164Phone)
}

func validateE164Phone(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^\+[1-9]\d{1,14}$`)
	return re.MatchString(fl.Field().String())
}