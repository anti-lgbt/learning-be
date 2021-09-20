package queries

type UserPassword struct {
	OldPassword     string `json:"old_password" form:"old_password"`
	NewPassword     string `json:"new_password" form:"new_password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

type UserPayload struct {
	FullName string `json:"full_name" form:"full_name"`
}

type IDQuery struct {
	ID uint64 `query:"id" validate:"uint"`
}
