package api

import (
	"pingoh/handlers"

	"github.com/gofiber/fiber/v2"
)

func addTaskRoutes(api *fiber.Router) {
	r := (*api).Group("/tasks")

	r.Post("/", func(c *fiber.Ctx) error {
		var b handlers.NewTask
		if err := c.BodyParser(&b); err != nil {
			return fiber.ErrBadRequest
		}
		err := handlers.Validator.Struct(&b)
		if err != nil {
			return err
		}
		err = handlers.CreateNewTask(&b)
		if err != nil {
			return err
		}
		return nil
	})

	r.Get("/:task_id/activate", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("task_id")
		if err != nil {
			return err
		}
		return handlers.ActivateTaskByID(id)
	})

	r.Get("/:task_id/deactivate", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("task_id")
		if err != nil {
			return err
		}
		return handlers.DeactivateTaskByID(id)
	})
}
