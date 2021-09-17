package models

import "time"

type Comment struct {
	ID        uint64
	UserID    uint64
	ProductID uint64
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
