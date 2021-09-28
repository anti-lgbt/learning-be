package resource

import (
	"github.com/anti-lgbt/learning-be/config"
	"github.com/anti-lgbt/learning-be/controllers/entities"
	"github.com/anti-lgbt/learning-be/controllers/helpers"
	"github.com/anti-lgbt/learning-be/controllers/queries"
	"github.com/anti-lgbt/learning-be/models"
	"github.com/anti-lgbt/learning-be/types"
	"github.com/gofiber/fiber/v2"
)

func CreateProductComment(c *fiber.Ctx) error {
	CurrentUser := c.Locals("CurrentUser").(*models.User)

	product_id, err := c.ParamsInt("product_id")
	if err != nil {
		return c.Status(500).JSON(types.Error{
			Error: "Product id phải là số nguyên",
		})
	}

	params := new(queries.CommentPayload)
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

	comment := &models.Comment{
		UserID:    CurrentUser.ID,
		ProductID: uint64(product_id),
		Content:   params.Content,
	}
	if result := config.DataBase.Preload("User").Create(&comment); result.Error != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Tạo comment thất bại",
		})
	}

	return c.Status(201).JSON(entities.Comment{
		ID:        comment.ID,
		UserID:    comment.UserID,
		FullName:  CurrentUser.FullName,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	})
}
