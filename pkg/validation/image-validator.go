package validation

import (
	"github.com/go-playground/validator/v10"
	"mime/multipart"
)

func imageValidation(fl validator.FieldLevel) bool {
	image, ok := fl.Field().Interface().(*multipart.FileHeader)

	if !ok || image == nil {
		return true
	}

	const maxSize = 5 << 20 // 5MB

	if image.Size > maxSize {
		return false
	}

	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/jpg":  true,
		"image/png":  true,
		"image/gif":  true,
		"image/bmp":  true,
		"image/webp": true,
	}

	if _, ok := allowedTypes[image.Header.Get("Content-Type")]; ok {
		return true
	}
	
	return false
}
