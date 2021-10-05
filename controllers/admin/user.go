package admin

import (
	"time"

	"github.com/anti-lgbt/learning-be/config"
	"github.com/anti-lgbt/learning-be/controllers/admin/entities"
	"github.com/anti-lgbt/learning-be/controllers/admin/queries"
	"github.com/anti-lgbt/learning-be/controllers/helpers"
	"github.com/anti-lgbt/learning-be/models"
	"github.com/anti-lgbt/learning-be/types"
	"github.com/creasty/defaults"
	"github.com/gofiber/fiber/v2"
)

func userToEntity(user *models.User) entities.User {
	return entities.User{
		ID:        user.ID,
		Email:     user.Email,
		FullName:  user.FullName,
		State:     user.State,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func GetUsers(c *fiber.Ctx) error {
	var users []*models.User

	var params = new(queries.UserQuery)
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

	if len(params.Email) > 0 {
		tx = tx.Where("email = ?", params.Email)
	}

	if len(params.FullName) > 0 {
		tx = tx.Where("full_name LIKE ?", "%"+params.FullName+"%")
	}

	if len(params.State) > 0 {
		tx = tx.Where("state = ?", params.State)
	}

	if len(params.Role) > 0 {
		tx = tx.Where("role = ?", params.Role)
	}

	tx.Find(&users)

	user_entities := make([]entities.User, 0)
	for _, user := range users {
		user_entities = append(user_entities, userToEntity(user))
	}

	return c.Status(200).JSON(user_entities)
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không tìm thấy user",
		})
	}

	var user *models.User
	if result := config.DataBase.First(&user, id); result.Error != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không tìm thấy user",
		})
	}

	return c.Status(200).JSON(userToEntity(user))
}

func CreateUser(c *fiber.Ctx) error {
	params := new(queries.UserPayload)
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

	hashed, err := models.HashPassword(params.Password)
	if err != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không xác minh được password",
		})
	}

	user := &models.User{
		Email:    params.Email,
		Password: hashed,
		FullName: params.FullName,
		State:    params.State,
		Role:     params.Role,
	}

	return c.Status(201).JSON(userToEntity(user))
}

func UpdateUser(c *fiber.Ctx) error {
	params := new(queries.UserPayload)
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

	var user *models.User
	if result := config.DataBase.First(&user, params.ID); result.Error != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không tìm thấy user",
		})
	}

	user.Email = params.Email

	if len(params.Password) > 0 {
		hashed, err := models.HashPassword(params.Password)
		if err != nil {
			return c.Status(422).JSON(types.Error{
				Error: "Không xác minh được password",
			})
		}

		user.Password = hashed
	}
	user.FullName = params.FullName
	user.State = params.State
	user.Role = params.Role

	return c.Status(200).JSON(userToEntity(user))
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không tìm thấy user",
		})
	}

	var user *models.User
	if result := config.DataBase.First(&user, id); result.Error != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không tìm thấy user",
		})
	}

	if result := config.DataBase.Delete(&user); result.Error != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không thể xoá user",
		})
	}

	return c.Status(200).JSON(200)
}
