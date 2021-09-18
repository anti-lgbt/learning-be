package public

import (
	"github.com/anti-lgbt/learning-be/config"
	"github.com/anti-lgbt/learning-be/controllers/entities"
	"github.com/anti-lgbt/learning-be/controllers/helpers"
	"github.com/anti-lgbt/learning-be/controllers/queries"
	"github.com/anti-lgbt/learning-be/models"
	"github.com/anti-lgbt/learning-be/types"
	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	var params = new(queries.ProductQuery)
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

	tx := config.DataBase

	if len(params.Type) > 0 {
		tx = tx.Where("product_type_id IN (SELECT \"id\" FROM \"product_type\" WHERE \"name\" LIKE ?)", "%"+params.Name+"%")
	}

	if len(params.Name) > 0 {
		tx = tx.Where("name LIKE ?", "%"+params.Name+"%")
	}

	var products []*models.Product
	tx.Preload("ProductType").Find(&products)

	product_entities := make([]*entities.Product, 0)

	for _, product := range products {
		product_entities = append(product_entities, &entities.Product{
			ID:                 product.ID,
			Type:               product.ProductType.Name,
			Name:               product.Name,
			Price:              product.Price,
			DiscountPercentage: product.DiscountPercentage,
			StockLeft:          product.StockLeft,
			Special:            product.Special,
			ViewCount:          product.ViewCount,
			CreatedAt:          product.CreatedAt,
			UpdatedAt:          product.UpdatedAt,
		})
	}

	return c.Status(200).JSON(product_entities)
}

func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(500).JSON(types.Error{
			Error: "Product id phải là số nguyên",
		})
	}

	var product *models.Product
	if result := config.DataBase.Preload("ProductType").First(&product, "id = ?", id); result.Error != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Product tìm thấy product",
		})
	}

	return c.Status(200).JSON(entities.Product{
		ID:                 product.ID,
		Type:               product.ProductType.Name,
		Name:               product.Name,
		Description:        product.Description,
		Price:              product.Price,
		DiscountPercentage: product.DiscountPercentage,
		StockLeft:          product.StockLeft,
		Special:            product.Special,
		ViewCount:          product.ViewCount,
		CreatedAt:          product.CreatedAt,
		UpdatedAt:          product.UpdatedAt,
	})
}
