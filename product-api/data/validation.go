package data

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Wraps validator field error
// so that we don't expose out to code
type ValidationError struct {
	validator.FieldError
}

// Convert Validation Error object to string
func (v *ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s', Error, field validation failed on '%s' with tag '%s' ",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

type ValidationErrors []ValidationError

// Convert slice of objects to string slice
func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

// Validation container
type Validation struct {
	validate *validator.Validate
}

func NewValidation() *Validation {
	v := validator.New()
	v.RegisterValidation("sku", ValidateSKU)
	return &Validation{v}
}

func (v *Validation) Validate(i interface{}) ValidationErrors {
	ReturnErrors := []ValidationError{}
	errs := v.validate.Struct(i)
	if errs == nil {
		return ReturnErrors
	}

	for _, err := range errs.(validator.ValidationErrors) {
		// Cast err to FieldError
		// then convert to ValidationError
		ve := ValidationError{
			err.(validator.FieldError),
		}
		ReturnErrors = append(ReturnErrors, ve)
	}
	return ReturnErrors
}

func ValidateSKU(fl validator.FieldLevel) bool {
	r := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := r.FindAllString(fl.Field().String(), -1)
	if len(matches) != 1 {
		return false
	}
	return true
}
