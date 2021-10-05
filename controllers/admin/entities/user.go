package entities

import (
	"time"

	"github.com/anti-lgbt/learning-be/types"
)

type User struct {
	ID        uint64          `json:"id"`
	Email     string          `json:"email"`
	FullName  string          `json:"full_name"`
	State     types.UserState `json:"state"`
	Role      types.UserRole  `json:"role"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
