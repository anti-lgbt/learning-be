package helpers

import (
	"errors"
	"mime/multipart"
	"net/http"

	"github.com/gookit/validate"
)

func Vaildate(value interface{}) error {
	v := validate.Struct(value)
	if !v.Validate() {
		for _, errs := range v.Errors.All() {
			for _, err := range errs {
				return errors.New(err)
			}
		}
	}

	return nil
}

func ValidateIsImage(file_header *multipart.FileHeader) bool {
	file, err := file_header.Open()
	if err != nil {
		return false
	}

	buff := make([]byte, 512) // docs tell that it take only first 512 bytes into consideration
	if _, err = file.Read(buff); err != nil {
		return false
	}

	content_type := http.DetectContentType(buff)

	if content_type != "image/jpeg" && content_type != "image/png" {
		return false
	}

	return true
}
