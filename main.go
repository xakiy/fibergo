package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xakiy/fibergo/database"
	"github.com/xakiy/fibergo/router"
)

func main() {
	// initiate new fiber app
	app := fiber.New()

	// Connecting to database
	database.ConnectDB()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	err := c.SendString("And the API is UP!")
	// 	return err
	// })

	// Setup the router
	router.SetupRoutes(app)

	// Listen on PORT 3000
	app.Listen(":3000")
}
