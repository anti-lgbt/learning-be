package routes

import (
	"github.com/anti-lgbt/learning-be/controllers/admin"
	"github.com/anti-lgbt/learning-be/controllers/identity"
	"github.com/anti-lgbt/learning-be/controllers/public"
	"github.com/anti-lgbt/learning-be/controllers/resource"
	"github.com/anti-lgbt/learning-be/routes/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRouter() *fiber.App {
	app := fiber.New()
	app.Use(logger.New(), encryptcookie.New(encryptcookie.Config{
		Key: "session_lmao",
	}))

	api_v2_public := app.Group("/api/v2/public")
	{
		api_v2_public.Get("/products", public.GetProducts)
		api_v2_public.Get("/products/types", public.GetProductTypes)
		api_v2_public.Get("/products/:id", public.GetProduct)
		api_v2_public.Get("/products/:id/image", public.GetProductImage)
		api_v2_public.Get("/products/:product_id/comments", public.GetComments)
	}

	api_v2_identity := app.Group("/api/v2/identity", middlewares.Guest)
	{
		api_v2_identity.Post("/login", identity.Login)
		api_v2_identity.Post("/register", identity.Register)
	}

	api_v2_resource := app.Group("/api/v2/resource", middlewares.Authenticate)
	{
		api_v2_resource.Get("/users/me", resource.GetUser)
		api_v2_resource.Put("/users/password", resource.UpdatePassword)
		api_v2_resource.Put("/users", resource.UpdateUser)
	}

	api_v2_admin := app.Group("/api/v2/admim", middlewares.Authenticate, middlewares.Admin)
	{
		api_v2_admin.Get("/products", admin.GetProducts)
		api_v2_admin.Get("/products/:id", admin.GetProduct)
		api_v2_admin.Post("/products", admin.CreateProduct)
		api_v2_admin.Put("/products", admin.UpdateProduct)
		api_v2_admin.Delete("/products/:id", admin.DeleteProduct)

		api_v2_admin.Get("/comments", admin.GetComments)
		api_v2_admin.Delete("/comments/:id", admin.DeleteComment)

		api_v2_admin.Get("/users", admin.GetUsers)
		api_v2_admin.Post("/users", admin.CreateUser)
		api_v2_admin.Put("/users", admin.UpdateUser)
		api_v2_admin.Delete("/users/:id", admin.DeleteUser)
	}

	return app
}
