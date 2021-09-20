package entities

import (
	"time"
)

type Comment struct {
	ID        uint64    `json:"id"`
	UserID    uint64    `json:"user_id"`
	FullName  string    `json:"full_name"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
