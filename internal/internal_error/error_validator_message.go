package internal_error

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) string {
	var sb strings.Builder

	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			switch e.Tag() {
			case "required":
				sb.WriteString(e.Field() + " is required. ")
			case "email":
				sb.WriteString("Invalid email format. ")
			case "min":
				sb.WriteString(e.Field() + " is too short. ")
			default:
				sb.WriteString("Invalid value for " + e.Field() + ". ")
			}
		}
		return sb.String()
	}

	return err.Error()
}
