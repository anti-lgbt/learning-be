package identity

import (
	"database/sql"

	"github.com/anti-lgbt/learning-be/config"
	"github.com/anti-lgbt/learning-be/models"
	"github.com/anti-lgbt/learning-be/types"
	"github.com/gofiber/fiber/v2"
)

type LoginPayload struct {
	Email    string
	Password string
}

func Login(c *fiber.Ctx) error {
	var params = new(LoginPayload)
	if err := c.BodyParser(params); err != nil {
		return c.Status(500).JSON(types.Error{
			Error: "Không thể xác minh được body",
		})
	}

	var user *models.User
	config.DataBase.First(&user, "email = ?", params.Email)

	hashed, err := models.HashPassword(params.Password)
	if err != nil || user.Password != hashed {
		return c.Status(500).JSON(types.Error{
			Error: "Sai mật khẩu",
		})
	}

	session, err := config.SessionStore.Get(c)
	if err != nil {
		return c.Status(500).JSON(types.Error{
			Error: "Không thể xác minh session",
		})
	}

	session.Set("email", user.Email)

	return c.Status(200).JSON(user)
}

type RegisterPayload struct {
	FullName sql.NullString
	Email    string
	Password string
}

func Register(c *fiber.Ctx) error {
	var params = new(RegisterPayload)
	if err := c.BodyParser(params); err != nil {
		return c.Status(500).JSON(types.Error{
			Error: "Không thể xác minh được body",
		})
	}

	var n_user *models.User
	if result := config.DataBase.First(&n_user, "email = ?", params.Email); result.Error == nil {
		return c.Status(500).JSON(types.Error{
			Error: "Email đã tồn tại",
		})
	}

	if len(params.Password) < 8 {
		return c.Status(500).JSON(types.Error{
			Error: "Password cần ít nhất 8 ký tự",
		})
	}

	user := &models.User{
		Email:    params.Email,
		Password: params.Password,
		FullName: params.FullName,
		State:    "active",
		Role:     "member",
	}
	if result := config.DataBase.Create(&user); result.Error != nil {
		return c.Status(500).JSON(types.Error{
			Error: "Không thể tạo user",
		})
	}

	session, err := config.SessionStore.Get(c)
	if err != nil {
		return c.Status(500).JSON(types.Error{
			Error: "Không thể xác minh session",
		})
	}

	session.Set("email", user.Email)

	return c.Status(201).JSON(user)
}
