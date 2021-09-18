package helpers

import (
	"errors"

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
