package types

type UserState string

var (
	UserStateActive  UserState = "active"
	UserStateDeleted UserState = "deleted"
)

type UserRole string

var (
	UserRoleAdmin  UserRole = "admin"
	UserRoleMember UserRole = "member"
)

type Error struct {
	Error string
}

type Ordering string

var (
	OrderingAsc  Ordering = "asc"
	OrderingDesc Ordering = "desc"
)
