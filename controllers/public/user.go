package public

import (
	"github.com/anti-lgbt/learning-be/config"
	"github.com/anti-lgbt/learning-be/controllers/helpers"
	"github.com/anti-lgbt/learning-be/controllers/queries"
	"github.com/anti-lgbt/learning-be/models"
	"github.com/anti-lgbt/learning-be/types"
	"github.com/gofiber/fiber/v2"
)

func GetUserAvatar(c *fiber.Ctx) error {
	params := new(queries.IDQuery)
	if err := c.QueryParser(params); err != nil {
		return c.Status(500).JSON(types.Error{
			Error: "Không thể xác minh được query",
		})
	}

	if err := helpers.Vaildate(params); err != nil {
		return c.Status(422).JSON(types.Error{
			Error: err.Error(),
		})
	}

	var user *models.User
	if result := config.DataBase.First(&user, params.ID); result.Error != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không tìm thấy user",
		})
	}

	if !user.Avatar.Valid {
		return c.Status(200).JSON(nil)
	}

	return c.SendFile(user.Avatar.String, false)
}
