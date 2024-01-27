package api

import (
	"pingoh/handlers"

	"github.com/gofiber/fiber/v2"
)

func addTaskRoutes(api *fiber.Router) {
	r := (*api).Group("/task")

	r.Post("/test", func(c *fiber.Ctx) error {
		return c.SendString("hee")
	})

	r.Post("/new", func(c *fiber.Ctx) error {
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

	r.Get("/activate/:task_id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("task_id")
		if err != nil {
			return err
		}
		return handlers.StartTaskByTaskID(id)
	})

	r.Get("/deactivate/:task_id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("task_id")
		if err != nil {
			return err
		}
		return handlers.StopTaskByTaskID(id)
	})
}
