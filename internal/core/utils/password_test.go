package utils

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Hash password valid", "password123", true},
		{"Hash mật khẩu rỗng", "", true},
		{"Hash mật khẩu có đúng 72 bytes", GenerateRandomString(72), true},
		{"Hash mật khẩu có nhiều hơn 72 bytes", GenerateRandomString(73), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := HashPassword(tt.input)
			if (err == nil) != tt.expected {
				t.Errorf("HashPassword() = %v, want %v", (err == nil), tt.expected)
			}
		})
	}
}

func TestComparePassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		hash     string
		expected bool
	}{
		{"So sánh mật khẩu đúng", "password123", MustHashPassword("password123"), true},
		{"So sánh mật khẩu sai", "password123", MustHashPassword("password1234"), false},
		{"So sánh mật khẩu rỗng", "", MustHashPassword(""), true},
		{"So sánh với hash không hợp lệ", "password123", "invalid_hash", false},
		{"So sánh với hash bị cắt bớt", "password123", MustHashPassword("password123")[:len(MustHashPassword("password123"))-1], false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComparePassword(tt.password, tt.hash); (got == nil) != tt.expected {
				t.Errorf("ComparePassword() = %v, want %v", (got == nil), tt.expected)
			}
		})
	}
}

func MustHashPassword(password string) string {
	hash, err := HashPassword(password)
	if err != nil {
		panic(err)
	}
	return hash
}
