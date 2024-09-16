package validator

import (
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

// LocationValidator is a custom validator for validating location
var LocationValidator validator.Func = func(fl validator.FieldLevel) bool {
	location, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	before, after, found := strings.Cut(location, ",")
	if !found {
		return false
	}

	latitude, err := strconv.ParseFloat(strings.TrimSpace(before), 64)
	if err != nil {
		return false
	}
	longitude, err := strconv.ParseFloat(strings.TrimSpace(after), 64)
	if err != nil {
		return false
	}

	if latitude < -90 || latitude > 90 || longitude < -180 || longitude > 180 {
		return false
	}

	return true
}
