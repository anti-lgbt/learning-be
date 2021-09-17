package models

import (
	"database/sql"
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID                 uint64          `gorm:"primaryKey;autoIncrement;not null;index"`
	ProductTypeID      uint64          `gorm:"type:integer;not null"`
	Name               string          `gorm:"type:character varying(255);not null;index"`
	Description        sql.NullString  `gorm:"type:character varying(255)"`
	Price              decimal.Decimal `gorm:"type:numeric(32,16);not null;default:0.0"`
	DiscountPercentage decimal.Decimal `gorm:"type:numeric(32,16);not null;default:0.0"`
	StockLeft          uint64          `gorm:"type:integer;not null;default:0"`
	Special            bool            `gorm:"type:boolean;index;default:false"`
	ViewCount          uint64          `gorm:"type:integer;not null;default:0"`
	Image              string          `gorm:"type:character varying(255)"`
	CreatedAt          time.Time       `gorm:"type:timestamp(0);not null;index"`
	UpdatedAt          time.Time       `gorm:"type:timestamp(0);not null;index"`
	Comments           []*Comment      `gorm:"constraint:OnDelete:CASCADE"`
	ProductType        *ProductType    `gorm:"constraint:OnDelete:CASCADE"`
}
