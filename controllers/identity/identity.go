package identity

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/anti-lgbt/learning-be/config"
	"github.com/anti-lgbt/learning-be/controllers/helpers"
	"github.com/anti-lgbt/learning-be/models"
	"github.com/anti-lgbt/learning-be/services"
	"github.com/anti-lgbt/learning-be/types"
	"github.com/gofiber/fiber/v2"
)

type AuthPayload struct {
	Email    string `json:"email" form:"email" validate:"email|required"`
	Password string `json:"password" form:"password" validate:"minLength:8|maxLength:26|required"`
}

type LoginPayload struct {
	AuthPayload
}

type RegisterPayload struct {
	FullName string `json:"full_name" form:"full_name" validate:"required"`
	AuthPayload
}

func Login(c *fiber.Ctx) error {
	var params = new(LoginPayload)
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
	config.DataBase.First(&user, "email = ?", params.Email)

	if !user.ComparePassword(params.Password) {
		return c.Status(422).JSON(types.Error{
			Error: "Sai mật khẩu",
		})
	}

	session, err := config.SessionStore.Get(c)
	if err != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không thể xác minh session",
		})
	}

	session.Set("email", user.Email)
	if err := session.Save(); err != nil {
		return c.Status(422).JSON(types.Error{
			Error: err.Error(),
		})
	}

	return c.Status(200).JSON(user.ToJSON())
}

func Register(c *fiber.Ctx) error {
	var params = new(RegisterPayload)
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

	var n_user *models.User
	if result := config.DataBase.First(&n_user, "email = ?", params.Email); result.Error == nil {
		return c.Status(422).JSON(types.Error{
			Error: "Email đã tồn tại",
		})
	}

	if len(params.Password) < 8 {
		return c.Status(422).JSON(types.Error{
			Error: "Password cần ít nhất 8 ký tự",
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
	}

	if result := config.DataBase.Create(&user); result.Error != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Không thể tạo user",
		})
	}

	session, err := config.SessionStore.Get(c)
	if err != nil {
		return c.Status(500).JSON(types.Error{
			Error: "Không thể xác minh session",
		})
	}

	session.Set("email", user.Email)
	if err := session.Save(); err != nil {
		return c.Status(500).JSON(types.Error{
			Error: "Không thể xác minh session",
		})
	}

	return c.Status(201).JSON(user.ToJSON())
}

func Logout(c *fiber.Ctx) error {
	session, err := config.SessionStore.Get(c)
	if err != nil {
		return c.Status(500).JSON(types.Error{
			Error: "Không thể xác minh session",
		})
	}

	if err := session.Destroy(); err != nil {
		return c.Status(500).JSON(types.Error{
			Error: "Không thể xác minh session",
		})
	}

	if err := session.Save(); err != nil {
		return c.Status(500).JSON(types.Error{
			Error: "Không thể xác minh session",
		})
	}

	return c.Status(200).JSON(200)
}

type ForgotPasswordPayload struct {
	Email string `json:"email" form:"email" validate:"email|required"`
}

func randomNumber(min, max int32) int32 {
	rand.Seed(time.Now().UnixNano())
	return min + int32(rand.Intn(int(max-min)))
}

func randomStringGenerator(charSet string, codeLength int32) string {
	code := ""
	charSetLength := int32(len(charSet))
	for i := int32(0); i < codeLength; i++ {
		index := randomNumber(0, charSetLength)
		code += string(charSet[index])
	}

	return code
}

func generateRandomStrongPassword() string {
	// random set of: abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ123456789!@#$&*_
	charSet := "RW!Vx&6fMHyPYSB1da*4LTm3ki@5c2ptgDzZ9Gq8w7Ke$XNE#s_jvrJuQnFCAUbh"
	return randomStringGenerator(charSet, 12)
}

func ForgotPassword(c *fiber.Ctx) error {
	var params = new(ForgotPasswordPayload)
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
	if result := config.DataBase.First(&user, "email = ?", params.Email); result.Error != nil {
		return c.Status(422).JSON(types.Error{
			Error: "Email không tồn tại",
		})
	}

	new_password := generateRandomStrongPassword()
	new_password_hashed, _ := models.HashPassword(new_password)

	user.Password = new_password_hashed
	config.DataBase.Save(&user)

	go services.SendEmail(user.Email, "Tài khoản X-SHOP của bạn vừa được khôi phục mật khẩu", fmt.Sprintf("Mật khẩu mới của bạn là: %s", new_password))

	return c.Status(200).JSON(200)
}
