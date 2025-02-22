package validator

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
)

func TestUserRoleValidator(t *testing.T) {
	validate := validator.New()

	// Đăng ký UserRoleValidator với tag "user_role"
	validate.RegisterValidation("user_role", UserRoleValidator)

	tests := []struct {
		name     string
		input    any
		expected bool
	}{
		{"Valid Role: Root", domain.Root, true},
		{"Valid Role: Member", domain.Member, true},
		{"Invalid Role: Admin", "Admin", false},
		{"Invalid Role: Empty", "", false},
		{"Invalid Role: invalid type", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Var(tt.input, "user_role")
			assert.Equal(t, tt.expected, err == nil)
		})
	}
}
