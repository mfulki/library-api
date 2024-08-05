package utils

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

func GetValidationErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "should be required"
	case "url", "email":
		return fmt.Sprintf("should be %s format", fe.Tag())
	case "e164":
		return "should be phone number (e164) format"
	case "date":
		return "should be date (yyyy-mm-dd) format"
	case "timestamp":
		return fmt.Sprintf("should be timestamp format like %s", time.RFC3339)
	case "unique":
		return "should be unique"
	case "min":
		return "length should be more than " + fe.Param() + " characters"
	case "max":
		return "length exceeds " + fe.Param() + " characters"
	case "lte":
		return "should be less than " + fe.Param()
	case "gte":
		return "should be greater or equal than " + fe.Param()
	case "oneof":
		return "should be one of " + fe.Param()
	}

	return fe.Error()
}
