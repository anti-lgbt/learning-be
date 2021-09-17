package models

type Product struct {
	ID uint64
	Type string
	Name string
	Description string
	Price decimal.Decimal
	DiscountPercentage decimal.Decimal
	StockLeft uint64
	Special bool
	ViewCount uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}