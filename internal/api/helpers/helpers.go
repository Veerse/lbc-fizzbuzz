package helpers

import (
	"github.com/go-playground/validator/v10"
)

func ParseValidationErrors(errors validator.ValidationErrors) string {
	ret := "Incorrect form fields: "

	for _, err := range errors {
		ret = ret + err.Field() + ", "
	}

	return ret[:len(ret)-2] + ". Please refer to documentation."
}
