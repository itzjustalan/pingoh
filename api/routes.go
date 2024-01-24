package api

import (
	"github.com/gofiber/fiber/v2"
)

func AddRoutes(app *fiber.App) {
	api := app.Group("/api")

	addAuthRoutes(&api)
	addAuthMiddle(&api)
	addTaskRoutes(&api)

	api.Get("/h", func(c *fiber.Ctx) error {
		return c.JSON([]string{"key", "val"})
	})
}
