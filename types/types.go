package types

type UserState string

var (
	UserStateActive  UserRole = "active"
	UserStateDeleted UserRole = "deleted"
)

type UserRole string

var (
	UserRoleAdmin  UserRole = "admin"
	UserRoleMember UserRole = "member"
)

type Error struct {
	Error string
}
