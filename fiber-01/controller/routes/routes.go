package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/paung29/controller/handlers"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/create-account", handlers.CreateAccount)
	api.Post("/login", handlers.Login)

	
}