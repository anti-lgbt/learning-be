package models

type ProductType struct {
	ID   uint64 `gorm:"primaryKey;autoIncrement;not null;index"`
	Name string `gorm:"type:character varying(20);not null;index"`
}
