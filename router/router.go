package router

import (
	"log",

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api", logger.New())  // Group endpoints with param 'api' and log whenever this endpoit is hit
}
