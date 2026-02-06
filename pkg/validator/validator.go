package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Init initializes custom validators
func Init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// Register custom validators here if needed
		_ = v
	}
}

// GetValidationErrors extracts validation errors from validator.ValidationErrors
func GetValidationErrors(err error) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			errors[fieldError.Field()] = getErrorMessage(fieldError)
		}
	}

	return errors
}

// getErrorMessage returns a human-readable error message
func getErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Value is too short"
	case "max":
		return "Value is too long"
	case "gte":
		return "Value must be greater than or equal to " + fe.Param()
	case "lte":
		return "Value must be less than or equal to " + fe.Param()
	default:
		return "Invalid value"
	}
}
