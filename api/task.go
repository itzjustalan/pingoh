package api

import "github.com/gofiber/fiber/v2"

func addTaskRoutes(api *fiber.Router) {
	r := (*api).Group("/task")

	r.Post("/test", func(c *fiber.Ctx) error {
		return c.SendString("hee")
	})

	r.Post("/new", func(c *fiber.Ctx) error {
		return c.SendString("task created")
	})
}
