package models

import (
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/anti-lgbt/learning-be/types"
)

type User struct {
	ID        uint64
	Email     string
	Password  string
	FullName  sql.NullString
	Avatar    sql.NullString
	State     types.UserState
	Role      types.UserRole
	CreatedAt time.Time
	UpdatedAt time.Time
}

func HashPassword(password string) (string, error) {

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return string(hash), err
	}

	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash), nil
}
