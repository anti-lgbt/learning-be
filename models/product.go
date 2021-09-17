package models

import (
	"database/sql"
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID                 uint64
	Type               string
	Name               string
	Description        sql.NullString
	Price              decimal.Decimal
	DiscountPercentage decimal.Decimal
	StockLeft          uint64
	Special            bool
	ViewCount          uint64
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
