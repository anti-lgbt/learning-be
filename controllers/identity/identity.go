package identity

import (
	"github.com/anti-lgbt/learning-be/config"
	"github.com/anti-lgbt/learning-be/controllers/helpers"
	"github.com/anti-lgbt/learning-be/models"
	"github.com/anti-lgbt/learning-be/types"
	"github.com/gofiber/fiber/v2"
)

type AuthPayload struct {
	Email    string `json:"email" form:"email" validate:"email|required"`
	Password string `json:"password" form:"password" validate:"minLength:8|maxLength:26|required"`
}

type LoginPayload struct {
	AuthPayload
}

type RegisterPayload struct {
	FullName string `json:"full_name" form:"full_name"`
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

	if !user.ComparePassword(params.Password) {
		return c.Status(422).JSON(types.Error{
			Error: "Sai mật khẩu",
		})
	}

	session, err := config.SessionStore.Get(c)
	if err != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không thể xác minh session",
		})
	}

	session.Set("email", user.Email)
	if err := session.Save(); err != nil {
		return c.Status(422).JSON(types.Error{
			Error: err.Error(),
		})
	}

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
	if err := session.Save(); err != nil {
		return c.Status(500).JSON(types.Error{
			Error: "Không thể xác minh session",
		})
	}

	return c.Status(201).JSON(user)
}

func Logout(c *fiber.Ctx) error {
	session, err := config.SessionStore.Get(c)
	if err != nil {
		return c.Status(500).JSON(types.Error{
			Error: "Không thể xác minh session",
		})
	}

	session.Delete("email")
	if err := session.Destroy(); err != nil {
		return c.Status(500).JSON(types.Error{
			Error: "Không thể xác minh session",
		})
	}

	if err := session.Save(); err != nil {
		return c.Status(500).JSON(types.Error{
			Error: "Không thể xác minh session",
		})
	}

	return c.Status(200).JSON(200)
}
