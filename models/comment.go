package models

import "time"

type Comment struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement;not null;index"`
	UserID    uint64    `gorm:"type:integer;not null;index"`
	ProductID uint64    `gorm:"type:integer;not null;index"`
	Content   string    `gorm:"type:character varying(255);not null;index"`
	CreatedAt time.Time `gorm:"type:timestamp(0);not null;index"`
	UpdatedAt time.Time `gorm:"type:timestamp(0);not null;index"`
	User      *User
	Product   *Product
}
