package models

import (
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/anti-lgbt/learning-be/controllers/entities"
	"github.com/anti-lgbt/learning-be/types"
)

type User struct {
	ID        uint64          `gorm:"primaryKey;autoIncrement;not null;index"`
	Email     string          `gorm:"type:character varying(50);not null;index,unique"`
	Password  string          `gorm:"type:character varying(255);not null"`
	FullName  string          `gorm:"type:character varying(255);not null;index"`
	Avatar    sql.NullString  `gorm:"type:character varying(50)"`
	State     types.UserState `gorm:"type:character varying(10);not null;index;default:active"`
	Role      types.UserRole  `gorm:"type:character varying(10);not null;index;default:member"`
	CreatedAt time.Time       `gorm:"type:timestamp(0);not null;index"`
	UpdatedAt time.Time       `gorm:"type:timestamp(0);not null;index"`
	Comments  []*Comment      `gorm:"constraint:OnDelete:CASCADE"`
}

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func ComparePassword(x, y string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(x), []byte(y))
	return err == nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) ToJSON() *entities.User {
	return &entities.User{
		ID:        u.ID,
		Email:     u.Email,
		FullName:  u.FullName,
		State:     u.State,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
