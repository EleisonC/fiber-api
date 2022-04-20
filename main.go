package main

import (
	"github.com/EleisonC/fiber-mongo-api/configs"
	"github.com/EleisonC/fiber-mongo-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	})
	// run database
	configs.ConnectDB()

	// routes
	routes.UserRoute(app)

	app.Listen(":8080")
}
