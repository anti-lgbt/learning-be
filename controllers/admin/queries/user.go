package queries

import (
	"github.com/anti-lgbt/learning-be/controllers/queries"
	"github.com/anti-lgbt/learning-be/types"
	"github.com/volatiletech/null"
)

type UserQuery struct {
	Email    string `query:"email" validate:"email"`
	FullName string `query:"full_name"`
	State    string `query:"state"`
	Role     string `query:"role"`
	queries.Order
	queries.Pagination
	queries.Period
}

type UserPayload struct {
	ID         uint64          `json:"id" form:"id"`
	Email      string          `json:"email" form:"email" validate:"email"`
	Password   string          `json:"password" form:"password"`
	FullName   string          `json:"full_name" form:"full_name" validate:"minLength:4|maxLength:50|required"`
	State      types.UserState `json:"state" form:"state" validate:"StateValidator|required"`
	Role       types.UserRole  `json:"role" form:"role" validate:"RoleValidator|required"`
	ReferralID null.Int64      `json:"referral_id" form:"referral_id"`
}

func (p UserPayload) StateValidator(val types.UserState) bool {
	return val == types.UserStateActive || val == types.UserStateDeleted
}

func (p UserPayload) RoleValidator(val types.UserRole) bool {
	return val == types.UserRoleAdmin || val == types.UserRoleMember
}
