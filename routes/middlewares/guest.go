package middlewares

import (
	"github.com/anti-lgbt/learning-be/config"
	"github.com/anti-lgbt/learning-be/types"
	"github.com/gofiber/fiber/v2"
)

func Guest(c *fiber.Ctx) error {
	session, err := config.SessionStore.Get(c)
	if err != nil {
		return c.Status(500).JSON(types.Error{
			Error: "Không thể xác minh session",
		})
	}

	email := session.Get("email")

	if email != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Vui lòng đăng xuất trước khi sử dụng chức năng này",
		})
	}

	return c.Next()
}
