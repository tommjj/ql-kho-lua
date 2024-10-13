package validator

import (
	"path/filepath"
	"slices"

	"github.com/go-playground/validator/v10"
)

var imageExt = []string{".png", ".jpg", ".jpeg", ".jfif", ".gif"}

// ImageFileValidator is a custom validator for validating image file name
var ImageFileValidator validator.Func = func(fl validator.FieldLevel) bool {
	filename, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	dir, file := filepath.Split(filename)
	if dir != "" {
		return false
	}

	ext := filepath.Ext(file)

	isImage := slices.Contains(imageExt, ext)
	return isImage
}
