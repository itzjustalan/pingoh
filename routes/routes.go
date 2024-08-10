package routes

import (
	"github.com/gofiber/fiber/v2"
)

func AddRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/hc", func(c *fiber.Ctx) error {
		return c.JSON([]string{"all", "gud"})
	})

	addAuthRoutes(&api)
	addAuthChecks(&api)
	addTaskRoutes(&api)
	addStatsRoutes(&api)
	addSharedRoutes(&api)
}
