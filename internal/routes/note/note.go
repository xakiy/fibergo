package noteRoute

import "github.com/gofiber/fiber/v2"

func SetupNoteRoutes(router fiber.Router) {

	note := router.Group("/note")

	// Create a Note
	note.Post("/", func(c *fiber.Ctx) error {})

	// Read all Notes
	note.Get("/", func(c *fiber.Ctx) error {})

	// Read one Note
	note.Get("/:noteId", func(c *fiber.Ctx) error {})

	// Update one Note
	note.Put("/:noteId", func(c *fiber.Ctx) error {})

	// Delete one Note
	note.Delete("/:noteId", func(c *fiber.Ctx) error {})
}
