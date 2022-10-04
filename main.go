package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// initiate new fiber app
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("And the API is UP!")
		return err
	})

	// Listen on PORT 3000
	app.Listen(":3000")
}
