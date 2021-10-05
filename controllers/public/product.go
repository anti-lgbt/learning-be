package public

import (
	"strings"

	"github.com/anti-lgbt/learning-be/config"
	"github.com/anti-lgbt/learning-be/controllers/entities"
	"github.com/anti-lgbt/learning-be/controllers/helpers"
	"github.com/anti-lgbt/learning-be/controllers/queries"
	"github.com/anti-lgbt/learning-be/models"
	"github.com/anti-lgbt/learning-be/types"
	"github.com/creasty/defaults"
	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/null"
)

func GetProducts(c *fiber.Ctx) error {
	var params = new(queries.ProductQuery)
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

	tx := config.DataBase.Offset(params.Page*params.Limit - params.Limit).Limit(params.Limit)

	if len(params.OrderBy) > 0 {
		tx = tx.Order(params.OrderBy + " " + string(params.Ordering))
	}

	if params.Discounting {
		tx = tx.Where("discount_percentage > 0")
	}

	if params.Special {
		tx = tx.Where("special = true")
	}

	if len(params.Type) > 0 {
		tx = tx.Where("product_type_id IN (SELECT \"id\" FROM \"product_types\" WHERE \"name\" LIKE ?)", "%"+params.Type+"%")
	}

	if len(params.Name) > 0 {
		tx = tx.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(params.Name)+"%")
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

	product.ViewCount++
	config.DataBase.Save(&product)

	return c.Status(200).JSON(entities.Product{
		ID:   product.ID,
		Type: product.ProductType.Name,
		Name: product.Name,
		Description: null.String{
			String: product.Description,
			Valid:  true,
		},
		Price:              product.Price,
		DiscountPercentage: product.DiscountPercentage,
		StockLeft:          product.StockLeft,
		Special:            product.Special,
		ViewCount:          product.ViewCount,
		CreatedAt:          product.CreatedAt,
		UpdatedAt:          product.UpdatedAt,
	})
}

func GetProductImage(c *fiber.Ctx) error {
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

	return c.SendFile(product.Image, false)
}

func GetProductTypes(c *fiber.Ctx) error {
	var product_types []*models.ProductType

	config.DataBase.Find(&product_types)

	product_type_entities := make([]*entities.ProductType, 0)
	for _, product_type := range product_types {
		product_type_entities = append(product_type_entities, &entities.ProductType{
			ID:        product_type.ID,
			Name:      product_type.Name,
			CreatedAt: product_type.CreatedAt,
			UpdatedAt: product_type.UpdatedAt,
		})
	}

	return c.Status(200).JSON(product_type_entities)
}
