package routes

import (
	"github.com/anti-lgbt/learning-be/controllers/identity"
	"github.com/anti-lgbt/learning-be/controllers/public"
	"github.com/anti-lgbt/learning-be/routes/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRouter() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())

	api_v2_public := app.Group("/api/v2/public")
	{
		api_v2_public.Get("/products", public.GetComments)
		api_v2_public.Get("/products/:id", public.GetProduct)
		api_v2_public.Get("/products/:product_id/comments", public.GetComments)
	}

	api_v2_identity := app.Group("/api/v2/identity", middlewares.Guest)
	{
		api_v2_identity.Post("/login", identity.Login)
		api_v2_identity.Post("/register", identity.Register)
	}

	api_v2_admin := app.Group("/api/v2/admim", middlewares.Authenticate, middlewares.Admin)
	{
		_ = api_v2_admin
	}

	return app
}
