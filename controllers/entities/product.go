package entities

import (
	"time"

	"github.com/shopspring/decimal"
	"github.com/volatiletech/null"
)

type Product struct {
	ID                 uint64          `json:"id"`
	Type               string          `json:"type"`
	Name               string          `json:"name"`
	Description        null.String     `json:"description"`
	Price              decimal.Decimal `json:"price"`
	DiscountPercentage decimal.Decimal `json:"discount_percentage"`
	StockLeft          uint64          `json:"stock_left"`
	Special            bool            `json:"special"`
	ViewCount          uint64          `json:"view_count"`
	CreatedAt          time.Time       `json:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at"`
}
