package public

import (
	"github.com/creasty/defaults"
	"github.com/gofiber/fiber/v2"

	"github.com/anti-lgbt/learning-be/config"
	"github.com/anti-lgbt/learning-be/controllers/entities"
	"github.com/anti-lgbt/learning-be/controllers/queries"
	"github.com/anti-lgbt/learning-be/models"
	"github.com/anti-lgbt/learning-be/types"
)

func GetComments(c *fiber.Ctx) error {
	product_id, err := c.ParamsInt("product_id")
	if err != nil {
		return c.Status(500).JSON(types.Error{
			Error: "Product id phải là số nguyên",
		})
	}

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

	var product *models.Product
	if result := config.DataBase.First(&product, "id = ?", product_id); result.Error != nil {
		return c.Status(500).JSON(types.Error{
			Error: "Không tìm thấy sản phẩm",
		})
	}

	var comments []*models.Comment
	config.DataBase.
		Where("product_id = ?", product.ID).
		Offset(params.Page*params.Limit - params.Limit).
		Limit(params.Limit).
		Preload("User").
		Find(&comments)

	comment_entities := make([]*entities.Comment, 0)
	for _, comment := range comments {
		comment_entities = append(comment_entities, &entities.Comment{
			ID:        comment.ID,
			UserID:    comment.UserID,
			FullName:  comment.User.FullName,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		})
	}

	return c.Status(200).JSON(comment_entities)
}
