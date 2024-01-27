package api

import (
	"pingoh/handlers"

	"github.com/gofiber/fiber/v2"
)

func addStatsRoutes(api *fiber.Router) {
	r := (*api).Group("/stats")

	r.Get("/task/:task_id", func(c *fiber.Ctx) error {
		tid, err := c.ParamsInt("task_id")
		if err != nil {
			return err
		}
		res, err := handlers.HttpResultsByTaskID(tid)
		if err != nil {
			return err
		}
		return c.JSON(res)
	})
}
