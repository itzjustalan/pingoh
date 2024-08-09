package api

import (
	"pingoh/handlers"

	"github.com/gofiber/fiber/v2"
)

func addSharedRoutes(api *fiber.Router) {
	r := (*api).Group("/shared")

	r.Get("/fetch", func(c *fiber.Ctx) error {
		var p handlers.FetchParams
		if err := c.QueryParser(&p); err != nil {
			return fiber.ErrBadRequest
		}
		err := handlers.Validator.Struct(&p)
		if err != nil {
			return err
		}
		p.M = c.Queries()
		res, err := handlers.Fetch(&p)
		if err != nil {
			return err
		}
		return c.JSON(res)
	})
}
