package resource

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/anti-lgbt/learning-be/config"
	"github.com/anti-lgbt/learning-be/controllers/entities"
	"github.com/anti-lgbt/learning-be/controllers/helpers"
	"github.com/anti-lgbt/learning-be/controllers/queries"
	"github.com/anti-lgbt/learning-be/models"
	"github.com/anti-lgbt/learning-be/types"
)

func GetUser(c *fiber.Ctx) error {
	CurrentUser := c.Locals("CurrentUser").(*models.User)

	return c.Status(200).JSON(entities.User{
		ID:        CurrentUser.ID,
		Email:     CurrentUser.Email,
		FullName:  CurrentUser.FullName,
		State:     CurrentUser.State,
		Role:      CurrentUser.Role,
		CreatedAt: CurrentUser.CreatedAt,
		UpdatedAt: CurrentUser.UpdatedAt,
	})
}

func UpdatePassword(c *fiber.Ctx) error {
	CurrentUser := c.Locals("CurrentUser").(*models.User)

	var params = new(queries.UserPassword)
	if err := c.BodyParser(params); err != nil {
		return c.Status(500).JSON(types.Error{
			Error: "Không thể xác minh được body",
		})
	}

	if err := helpers.Vaildate(params); err != nil {
		return c.Status(422).JSON(types.Error{
			Error: err.Error(),
		})
	}

	if !CurrentUser.ComparePassword(params.OldPassword) {
		return c.Status(422).JSON(types.Error{
			Error: "Sai mật khẩu",
		})
	}

	if models.ComparePassword(params.OldPassword, params.NewPassword) {
		return c.Status(422).JSON(types.Error{
			Error: "Mật khẩu mới phải khác mật khẩu cũ",
		})
	}

	if models.ComparePassword(params.NewPassword, params.ConfirmPassword) {
		return c.Status(422).JSON(types.Error{
			Error: "Nhập lại mật khẩu không khớp",
		})
	}

	hashed_new_password, err := models.HashPassword(params.NewPassword)
	if err != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không thể đổi mật khẩu",
		})
	}

	CurrentUser.Password = hashed_new_password

	config.DataBase.Save(&CurrentUser)

	return c.Status(200).JSON(200)
}

func UpdateUser(c *fiber.Ctx) error {
	CurrentUser := c.Locals("CurrentUser").(*models.User)

	var params = new(queries.UserPayload)
	if err := c.BodyParser(params); err != nil {
		return c.Status(500).JSON(types.Error{
			Error: "Không thể xác minh được body",
		})
	}

	if err := helpers.Vaildate(params); err != nil {
		return c.Status(422).JSON(types.Error{
			Error: err.Error(),
		})
	}

	if len(params.FullName) > 0 {
		CurrentUser.FullName = params.FullName
	}

	config.DataBase.Save(&CurrentUser)

	return c.Status(200).JSON(CurrentUser.ToJSON())
}

func UploadUserAvatar(c *fiber.Ctx) error {
	CurrentUser := c.Locals("CurrentUser").(*models.User)

	file_header, err := c.FormFile("avatar")
	if err != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không thể upload được ảnh",
		})
	}

	file, err := file_header.Open()
	if err != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không thể upload được ảnh",
		})
	}

	buff := make([]byte, 512) // docs tell that it take only first 512 bytes into consideration
	if _, err = file.Read(buff); err != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không thể upload được ảnh",
		})
	}

	content_type := http.DetectContentType(buff)

	if content_type != "image/jpeg" && content_type != "image/png" {
		return c.Status(422).JSON(types.Error{
			Error: "Sai định dạng ảnh",
		})
	}

	image_path := fmt.Sprintf("./uploads/%s", file_header.Filename)
	if err := c.SaveFile(file_header, image_path); err != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không thể upload được ảnh",
		})
	}

	CurrentUser.Avatar = sql.NullString{
		String: image_path,
		Valid:  true,
	}

	config.DataBase.Save(&CurrentUser)

	return c.Status(200).JSON(200)
}

func GetUserAvatar(c *fiber.Ctx) error {
	CurrentUser := c.Locals("CurrentUser").(*models.User)

	if !CurrentUser.Avatar.Valid {
		return c.Status(200).JSON(nil)
	}

	return c.SendFile(CurrentUser.Avatar.String, false)
}
