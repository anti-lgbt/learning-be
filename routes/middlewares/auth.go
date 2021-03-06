package middlewares

import (
	"github.com/anti-lgbt/learning-be/config"
	"github.com/anti-lgbt/learning-be/models"
	"github.com/anti-lgbt/learning-be/types"
	"github.com/gofiber/fiber/v2"
)

func Authenticate(c *fiber.Ctx) error {
	session, err := config.SessionStore.Get(c)
	if err != nil {
		return c.Status(500).JSON(types.Error{
			Error: "Không thể xác minh session",
		})
	}

	email := session.Get("email")
	if email == nil {
		return c.Status(401).JSON(types.Error{
			Error: "Session không tồn tại",
		})
	}

	var user *models.User
	if result := config.DataBase.First(&user, "email = ?", email.(string)); result.Error != nil {
		session.Destroy()
		return c.Status(500).JSON(types.Error{
			Error: "Session gặp lỗi",
		})
	}

	c.Locals("CurrentUser", user)

	return c.Next()
}
