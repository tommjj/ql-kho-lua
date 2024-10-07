package validator

import (
	"github.com/go-playground/validator/v10"
)

// LocationValidator is a custom validator for validating location
var LocationValidator validator.Func = func(fl validator.FieldLevel) bool {
	location, ok := fl.Field().Interface().([]float64)
	if !ok {
		return false
	}

	if len(location) != 2 {
		return false
	}

	latitude := location[0]
	longitude := location[1]

	if latitude < -90 || latitude > 90 || longitude < -180 || longitude > 180 {
		return false
	}

	return true
}
