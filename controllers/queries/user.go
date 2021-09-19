package queries

import (
	"mime/multipart"

	"github.com/volatiletech/null"
)

type UserPassword struct {
	OldPassword     string `json:"old_password" form:"old_password"`
	NewPassword     string `json:"new_password" form:"new_password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

type UserPayload struct {
	FullName null.String           `json:"full_name" form:"full_name"`
	Avatar   *multipart.FileHeader `json:"avatar" form:"avatar" validate:"file/isFile|image/isImage"`
}
