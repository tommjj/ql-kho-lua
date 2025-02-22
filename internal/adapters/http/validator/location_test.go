package validator

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestLocationValidator(t *testing.T) {
	validate := validator.New()

	// Đăng ký LocationValidator với tag "location"
	validate.RegisterValidation("location", LocationValidator)

	tests := []struct {
		name     string
		input    any
		expected bool
	}{
		{"Tọa độ trung tâm (0,0)", []float64{0, 0}, true},
		{"Cực Bắc và Kinh tuyến gốc", []float64{90, 0}, true},
		{"Cực Nam và Kinh tuyến gốc", []float64{-90, 0}, true},
		{"Cực Đông và Xích đạo", []float64{0, 180}, true},
		{"Cực Tây và Xích đạo", []float64{0, -180}, true},
		{"Vượt quá giới hạn vĩ độ", []float64{91, 0}, false},
		{"Vượt quá giới hạn vĩ độ", []float64{-91, 0}, false},
		{"Vượt quá giới hạn kinh độ", []float64{0, 181}, false},
		{"Vượt quá giới hạn kinh độ", []float64{0, -181}, false},
		{"Chỉ có một phần tử", []float64{0}, false},
		{"Không có phần tử nào", []float64{}, false},
		{"Nhiều hơn 2 phần tử", []float64{10, 20, 30}, false},
		{"Giá trị không phải float64", []any{"10", 20}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Var(tt.input, "location")
			assert.Equal(t, tt.expected, err == nil)
		})
	}
}
