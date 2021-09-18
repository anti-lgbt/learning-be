package models

import "time"

type ProductType struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement;not null;index"`
	Name      string    `gorm:"type:character varying(20);not null;index"`
	CreatedAt time.Time `gorm:"type:timestamp(0);not null;index"`
	UpdatedAt time.Time `gorm:"type:timestamp(0);not null;index"`
}
