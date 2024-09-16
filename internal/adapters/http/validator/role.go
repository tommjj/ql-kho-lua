package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
)

// userRoleValidator is a custom validator for validating user roles
var UserRoleValidator validator.Func = func(fl validator.FieldLevel) bool {
	userRole, ok := fl.Field().Interface().(domain.Role)
	if !ok {
		return false
	}

	switch userRole {
	case domain.Root, domain.Member:
		return true
	default:
		return false
	}
}
