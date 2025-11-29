package main

import (
	"myfiber/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDatabase()
	app.Listen(":3000")
}
