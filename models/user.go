package models

type User struct {
	ID uint64
	UID string
	Email string
	Password string
	FullName sql.NullString
	Avatar sql.NullString
	State types.UserState
	Role types.UserRole
	CreatedAt time.Time
	UpdatedAt time.Time
}
