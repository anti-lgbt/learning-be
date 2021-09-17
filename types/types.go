package types

type UserState string;

var (
	UserStateActive = "active"
	UserStateDeleted = "deleted"
)

type UserRole string;

var (
	UserRoleAdmin = "admin"
	UserRoleMember = "member"
)
