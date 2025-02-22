package validator

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestImageValidator(t *testing.T) {
	validate := validator.New()

	validate.RegisterValidation("image", ImageFileValidator)

	tests := []struct {
		name     string
		input    any
		expected bool
	}{
		{"Valid image: .png", "image.png", true},
		{"Valid image: .jpg", "image.jpg", true},
		{"Valid image: .jpeg", "image.jpeg", true},
		{"Valid image: .jfif", "image.jfif", true},
		{"Valid image: .gif", "image.gif", true},
		{"Invalid image: Tệp không có phần mở rộng", "file_without_ext", false},
		{"Invalid image: Tệp có phần mở rộng không hợp lệ", "image.pdf", false},
		{"Invalid image: Đường dẫn chứa thư mục", "folder/image.png", false},
		{"Invalid image: Đường dẫn tuyệt đối", "/absolute/path/image.png", false},
		{"Invalid image: Đường dẫn tương đối", "../image.png", false},
		{"Invalid image: Giá trị rỗng", "", false},
		{"Invalid image: Giá trị không phải string", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Var(tt.input, "image")
			assert.Equal(t, tt.expected, err == nil)
		})
	}
}
