package api

import "github.com/gofiber/fiber/v2"

func addMonitorRoutes(api *fiber.Router) {
	r := (*api).Group("/monitor")

	r.Post("/test", func(c *fiber.Ctx) error {
		return c.SendString("hee")
	})
}
