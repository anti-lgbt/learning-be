package resource

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/null"

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
		ID:    CurrentUser.ID,
		Email: CurrentUser.Email,
		FullName: null.String{
			String: CurrentUser.FullName.String,
			Valid:  CurrentUser.FullName.Valid,
		},
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

	hashed_old_password, err := models.HashPassword(params.OldPassword)
	if err != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Sai mật khẩu",
		})
	}

	if CurrentUser.Password != hashed_old_password {
		return c.Status(422).JSON(types.Error{
			Error: "Sai mật khẩu",
		})
	}

	if params.OldPassword == params.NewPassword {
		return c.Status(422).JSON(types.Error{
			Error: "Mật khẩu mới phải khác mật khẩu cũ",
		})
	}

	if params.NewPassword != params.ConfirmPassword {
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

	if params.FullName.Valid {
		CurrentUser.FullName = sql.NullString{
			String: params.FullName.String,
			Valid:  true,
		}
	}

	if params.Avatar != nil {
		image_path := fmt.Sprintf("./uploads/%s", params.Avatar.Filename)
		if err := c.SaveFile(params.Avatar, image_path); err != nil {
			return c.Status(422).JSON(types.Error{
				Error: "Không thể upload được ảnh",
			})
		}

		CurrentUser.Avatar = sql.NullString{
			String: image_path,
			Valid:  true,
		}
	}

	return c.Status(200).JSON(entities.User{
		ID:    CurrentUser.ID,
		Email: CurrentUser.Email,
		FullName: null.String{
			String: CurrentUser.FullName.String,
			Valid:  CurrentUser.FullName.Valid,
		},
		State:     CurrentUser.State,
		Role:      CurrentUser.Role,
		CreatedAt: CurrentUser.CreatedAt,
		UpdatedAt: CurrentUser.UpdatedAt,
	})
}