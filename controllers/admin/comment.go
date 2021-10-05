package admin

import (
	"time"

	"github.com/creasty/defaults"
	"github.com/gofiber/fiber/v2"

	"github.com/anti-lgbt/learning-be/config"
	"github.com/anti-lgbt/learning-be/controllers/admin/entities"
	"github.com/anti-lgbt/learning-be/controllers/admin/queries"
	"github.com/anti-lgbt/learning-be/controllers/helpers"
	"github.com/anti-lgbt/learning-be/models"
	"github.com/anti-lgbt/learning-be/types"
)

func commentToEntity(comment *models.Comment) entities.Comment {
	return entities.Comment{
		ID:        comment.ID,
		UserID:    comment.UserID,
		ProductID: comment.ProductID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}
}

func GetComments(c *fiber.Ctx) error {
	var comments []*models.Comment

	var params = new(queries.CommentQuery)
	if err := c.QueryParser(params); err != nil {
		return c.Status(500).JSON(types.Error{
			Error: "Không thể xác minh được query",
		})
	}

	if err := defaults.Set(params); err != nil {
		return c.Status(500).JSON(types.Error{
			Error: "Không thể xác minh được query",
		})
	}

	if err := helpers.Vaildate(params); err != nil {
		return c.Status(422).JSON(types.Error{
			Error: err.Error(),
		})
	}

	tx := config.DataBase.
		Offset(params.Page*params.Limit - params.Limit).
		Limit(params.Limit)

	if params.TimeFrom > 0 {
		tx = tx.Where("created_at >= ?", time.Unix(params.TimeFrom, 0))
	}

	if params.TimeTo > 0 {
		tx = tx.Where("updated_at >= ?", time.Unix(params.TimeTo, 0))
	}

	if params.ProductID > 0 {
		tx.Where("product_id", params.ProductID)
	}

	if params.UserID > 0 {
		tx.Where("user_id", params.UserID)
	}

	tx.Find(&comments)

	var comment_entities = make([]entities.Comment, 0)
	for _, comment := range comments {
		comment_entities = append(comment_entities, commentToEntity(comment))
	}

	return c.Status(200).JSON(comment_entities)
}

func GetComment(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không tìm thấy comment",
		})
	}

	var comment *models.Comment
	if result := config.DataBase.First(&comment, id); result.Error != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không tìm thấy comment",
		})
	}

	return c.Status(200).JSON(commentToEntity(comment))
}

func DeleteComment(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không tìm thấy comment",
		})
	}

	var comment *models.Comment
	if result := config.DataBase.First(&comment, id); result.Error != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không tìm thấy comment",
		})
	}

	if result := config.DataBase.Delete(&comment, id); result.Error != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không xoá được comment",
		})
	}

	return c.Status(200).JSON(200)
}
