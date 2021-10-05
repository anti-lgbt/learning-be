package queries

import (
	"github.com/anti-lgbt/learning-be/controllers/queries"
	"github.com/shopspring/decimal"
)

type ProductQuery struct {
	queries.ProductQuery
	queries.Period
}

type ProductPayload struct {
	ID                 uint64          `json:"id" form:"id" validate:"uint"`
	ProductTypeID      uint64          `json:"product_type_id" form:"product_type_id" validate:"required"`
	Name               string          `json:"name" form:"name" validate:"required"`
	Description        string          `json:"description" form:"description" validate:"required"`
	Price              decimal.Decimal `json:"price" form:"price" validate:"decimalPositive|required"`
	DiscountPercentage decimal.Decimal `json:"discount_percentage" form:"discount_percentage" validate:"DiscountPercentageValidator|required"`
	StockLeft          uint64          `json:"stock_left" form:"stock_left" validate:"uint|required"`
	Special            bool            `json:"special" form:"special" validate:"bool|required"`
}

func (p ProductPayload) DiscountPercentageValidator(val decimal.Decimal) bool {
	return !val.IsNegative()
}

type ProductTypePayload struct {
	ID   uint64 `json:"id" form:"id" validate:"uint"`
	Name string `json:"name" form:"name" validate:"required"`
}
