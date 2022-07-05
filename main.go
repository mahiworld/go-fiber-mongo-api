package main

import (
	"go-fiber-mongo-api/configs"
	"go-fiber-mongo-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.JSON(&fiber.Map{"data": "Hello from Go, Fiber & mongoDB Api"})
	// })

	//connect to database
	configs.ConnectDB()

	//routes
	routes.UserRoutes(app)

	app.Listen(":6000")
}
