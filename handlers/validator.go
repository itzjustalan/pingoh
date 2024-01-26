package handlers

import (
	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func init() {
	NewValidator()
}

func NewValidator() {
	Validator = validator.New(validator.WithRequiredStructEnabled())
}

// this check is only needed when your code could produce
// an invalid value for validation such as interface with nil
// value most including myself do not usually have code like this.
// func CheckInvalidValidation(err error) {
// 	if _, ok := err.(*validator.InvalidValidationError); ok {
// 		return
// 	}
// }
