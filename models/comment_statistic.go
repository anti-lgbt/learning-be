package models

import "time"

type CommentStatistic struct {
	ID           uint64    `gorm:"primaryKey;autoIncrement;not null;index"`
	Count        uint64    `gorm:"type:integer;not null;index"`
	ReferralDate time.Time `gorm:"type:timestamp(0);not null;index"`
	CreatedAt    time.Time `gorm:"type:timestamp(0);not null;index"`
	UpdatedAt    time.Time `gorm:"type:timestamp(0);not null;index"`
}
