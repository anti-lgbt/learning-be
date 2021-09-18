package config

import (
	"github.com/gookit/validate"
	"github.com/shopspring/decimal"
)

func InitValidator() {
	validate.AddValidator("decimalPositive", func(val decimal.Decimal) bool {
		return val.IsPositive()
	})
}
