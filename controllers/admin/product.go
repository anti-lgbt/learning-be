package admin

import (
	"fmt"
	"strings"
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

func productToEntity(product *models.Product) entities.Product {
	return entities.Product{
		ID:                 product.ID,
		ProductTypeID:      product.ProductTypeID,
		Name:               product.Name,
		Description:        product.Description,
		Price:              product.Price,
		DiscountPercentage: product.DiscountPercentage,
		StockLeft:          product.StockLeft,
		Special:            product.Special,
		ViewCount:          product.ViewCount,
		CreatedAt:          product.CreatedAt,
		UpdatedAt:          product.UpdatedAt,
	}
}

func productTypeToEntity(product_type *models.ProductType) entities.ProductType {
	return entities.ProductType{
		ID:        product_type.ID,
		Name:      product_type.Name,
		CreatedAt: product_type.CreatedAt,
		UpdatedAt: product_type.UpdatedAt,
	}
}

func GetProducts(c *fiber.Ctx) error {
	var products []*models.Product

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

	if params.TimeFrom > 0 {
		tx = tx.Where("created_at >= ?", time.Unix(params.TimeFrom, 0))
	}

	if params.TimeTo > 0 {
		tx = tx.Where("updated_at >= ?", time.Unix(params.TimeTo, 0))
	}

	tx.Find(&products)

	product_entities := make([]entities.Product, 0)
	for _, product := range products {
		product_entities = append(product_entities, productToEntity(product))
	}

	return c.Status(200).JSON(product_entities)
}

func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không tìm thấy product",
		})
	}

	var product *models.Product
	if result := config.DataBase.First(&product, id); result.Error != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không tìm thấy product",
		})
	}

	return c.Status(200).JSON(productToEntity(product))
}

func CreateProduct(c *fiber.Ctx) error {
	params := new(queries.ProductPayload)
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

	var product_type *models.ProductType
	if result := config.DataBase.First(&product_type, params.ProductTypeID); result.Error != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Type không xác định",
		})
	}

	product := &models.Product{
		ProductTypeID:      product_type.ID,
		Name:               params.Name,
		Description:        params.Description,
		Price:              params.Price,
		DiscountPercentage: params.DiscountPercentage,
		StockLeft:          params.StockLeft,
		Special:            params.Special,
	}

	image, err := c.FormFile("image")
	if err == nil {
		image_path := fmt.Sprintf("./uploads/%s", image.Filename)
		if err := c.SaveFile(image, image_path); err != nil {
			return c.Status(422).JSON(types.Error{
				Error: "Không thể upload được ảnh",
			})
		}

		product.Image = image_path
	}

	if result := config.DataBase.Create(&product); result.Error != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không thể tạo product",
		})
	}

	return c.Status(201).JSON(productToEntity(product))
}

func UpdateProduct(c *fiber.Ctx) error {
	params := new(queries.ProductPayload)
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

	var product_type *models.ProductType
	if result := config.DataBase.First(&product_type, params.ProductTypeID); result.Error != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Type không xác định",
		})
	}

	var product *models.Product
	if result := config.DataBase.First(&product, params.ID); result.Error != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không tìm thấy product",
		})
	}

	product.ProductTypeID = product_type.ID
	product.Name = params.Name
	product.Description = params.Description
	product.Price = params.Price
	product.DiscountPercentage = params.DiscountPercentage
	product.StockLeft = params.StockLeft
	product.Special = params.Special

	config.DataBase.Save(&product)

	return c.Status(200).JSON(productToEntity(product))
}

func UpdateProductImage(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không tìm thấy product",
		})
	}

	var product *models.Product
	if result := config.DataBase.First(&product, id); result.Error != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không tìm thấy product",
		})
	}

	image, err := c.FormFile("image")
	if err != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không tìm thấy ảnh",
		})
	}

	image_path := fmt.Sprintf("./uploads/%s", image.Filename)
	if err := c.SaveFile(image, image_path); err != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không thể upload được ảnh",
		})
	}

	product.Image = image_path

	config.DataBase.Save(&product)

	return c.Status(200).JSON(200)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không tìm thấy product",
		})
	}

	var product *models.Product
	if result := config.DataBase.First(&product, id); result.Error != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không tìm thấy product",
		})
	}

	if result := config.DataBase.Delete(&product); result.Error != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không thể xoá product",
		})
	}

	return c.Status(200).JSON(200)
}

func GetProductTypes(c *fiber.Ctx) error {
	var product_types []*models.ProductType

	config.DataBase.Find(&product_types)

	product_type_entities := make([]entities.ProductType, 0)
	for _, product_type := range product_types {
		product_type_entities = append(product_type_entities, productTypeToEntity(product_type))
	}

	return c.Status(200).JSON(product_type_entities)
}

func GetProductType(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không tìm thấy product type",
		})
	}

	var product_type *models.ProductType
	if result := config.DataBase.First(&product_type, id); result.Error != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không tìm thấy product type",
		})
	}

	return c.Status(200).JSON(productTypeToEntity(product_type))
}

func CreateProductType(c *fiber.Ctx) error {
	params := new(queries.ProductTypePayload)
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

	var product_type *models.ProductType
	if result := config.DataBase.First(&product_type, "name = ?", params.Name); result.Error == nil {
		return c.Status(422).JSON(types.Error{
			Error: "Loại hàng này đã tồn tại",
		})
	}

	product_type = &models.ProductType{
		Name: params.Name,
	}

	if result := config.DataBase.Create(&product_type); result.Error != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Tạo thất bại",
		})
	}

	return c.Status(201).JSON(productTypeToEntity(product_type))
}

func UpdateProductType(c *fiber.Ctx) error {
	params := new(queries.ProductTypePayload)
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

	var product_type *models.ProductType
	if result := config.DataBase.First(&product_type, params.ID); result.Error != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không tìm thấy loại hàng",
		})
	}

	product_type.Name = params.Name

	config.DataBase.Save(&product_type)

	return c.Status(200).JSON(productTypeToEntity(product_type))
}

func DeleteProductType(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không tìm thấy product type",
		})
	}

	var product_type *models.ProductType
	if result := config.DataBase.First(&product_type, id); result.Error != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không tìm thấy product type",
		})
	}

	if result := config.DataBase.Delete(&product_type); result.Error != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không thể xoá product type",
		})
	}

	return c.Status(200).JSON(200)
}
