package identity

import (
	"github.com/anti-lgbt/learning-be/config"
	"github.com/anti-lgbt/learning-be/controllers/helpers"
	"github.com/anti-lgbt/learning-be/models"
	"github.com/anti-lgbt/learning-be/types"
	"github.com/gofiber/fiber/v2"
)

type AuthPayload struct {
	Email    string `json:"email" validate:"email|required"`
	Password string `json:"password" validate:"min:8|max:26|required"`
}

type LoginPayload struct {
	AuthPayload
}

type RegisterPayload struct {
	FullName string `json:"full_name"`
	AuthPayload
}

func Login(c *fiber.Ctx) error {
	var params = new(LoginPayload)
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

func Register(c *fiber.Ctx) error {
	var params = new(RegisterPayload)
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

	hashed, err := models.HashPassword(params.Password)
	if err != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không xác minh được password",
		})
	}

	user := &models.User{
		Email:    params.Email,
		Password: hashed,
		FullName: params.FullName,
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
