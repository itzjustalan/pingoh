package routes

import (
	"pingoh/controllers"

	"github.com/gofiber/fiber/v2"
)

func addSharedRoutes(api *fiber.Router) {
	r := (*api).Group("/shared")

	r.Get("/fetch", func(c *fiber.Ctx) error {
		var p controllers.FetchParams
		if err := c.QueryParser(&p); err != nil {
			return fiber.ErrBadRequest
		}
		err := controllers.Validator.Struct(&p)
		if err != nil {
			return err
		}
		p.M = c.Queries()
		res, err := controllers.Fetch(&p)
		if err != nil {
			return err
		}
		return c.JSON(res)
	})
}
