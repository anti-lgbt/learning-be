package middlewares

import (
	"github.com/gofiber/fiber/v2"

	"github.com/anti-lgbt/learning-be/models"
	"github.com/anti-lgbt/learning-be/types"
)

func Admin(c *fiber.Ctx) error {
	CurrentUser := c.Locals("CurrentUser").(*models.User)

	if CurrentUser.Role != "admin" {
		return c.Status(422).JSON(types.Error{
			Error: "Không đủ quyền hạn!",
		})
	}

	return c.Next()
}
