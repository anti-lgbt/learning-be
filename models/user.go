package models

import (
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/anti-lgbt/learning-be/types"
)

type User struct {
	ID        uint64          `gorm:"primaryKey;autoIncrement;not null;index"`
	Email     string          `gorm:"type:character varying(50);not null;index"`
	Password  string          `gorm:"type:character varying(255);not null"`
	FullName  sql.NullString  `gorm:"type:character varying(255);index"`
	Avatar    sql.NullString  `gorm:"type:character varying(50)"`
	State     types.UserState `gorm:"type:character varying(10);not null;index;default:active"`
	Role      types.UserRole  `gorm:"type:character varying(10);not null;index;default:member"`
	CreatedAt time.Time       `gorm:"type:timestamp(0);not null;index"`
	UpdatedAt time.Time       `gorm:"type:timestamp(0);not null;index"`
	Comments  []*Comment      `gorm:"constraint:OnDelete:CASCADE"`
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
