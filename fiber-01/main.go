package main

import (
	"log"

	"github.com/paung29/controller/routes"
	"github.com/paung29/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDatabase()
	routes.RegisterRoutes(app)

	log.Println("🚀 Server running on http://localhost:3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
