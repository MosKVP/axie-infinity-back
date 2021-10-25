package util

import (
	"breeding/internal/model"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ConvertValidationError(e validator.ValidationErrors) *model.Error {
	var errList []string
	for _, err := range e {
		errList = append(errList, fieldErrorToText(err))
	}
	intErr := &model.Error{
		HTTPCode: http.StatusBadRequest,
		Message:  strings.Join(errList, ", "),
	}
	return intErr
}

func fieldErrorToText(e validator.FieldError) string {
	field := e.Field()

	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s", field, e.Param())
	case "min":
		return fmt.Sprintf("%s must be longer than %s", field, e.Param())
	case "email":
		return "Invalid email format"
	case "len":
		return fmt.Sprintf("%s must be %s characters long", field, e.Param())
	case "numeric":
		return fmt.Sprintf("%s must be numeric", field)
	}
	return fmt.Sprintf("%s is not valid", field)
}
