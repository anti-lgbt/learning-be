package public

import (
	"github.com/gofiber/fiber/v2"

	"github.com/anti-lgbt/learning-be/config"
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
	if err := c.BodyParser(params); err != nil {
		return c.Status(500).JSON(types.Error{
			Error: "Không thể xác minh được body",
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
		Find(&comments)

	return c.Status(200).JSON(comments)
}
