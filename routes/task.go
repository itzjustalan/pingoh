package routes

import (
	"pingoh/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func addTaskRoutes(api *fiber.Router) {
	r := (*api).Group("/tasks")

	r.Post("/", func(c *fiber.Ctx) error {
		var b controllers.NewTask
		if err := c.BodyParser(&b); err != nil {
			return fiber.ErrBadRequest
		}
		err := controllers.Validator.Struct(&b)
		if err != nil {
			return err
		}
		return controllers.CreateNewTask(&b)
	})

	r.Get("/:task_id/activate", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("task_id")
		if err != nil {
			return err
		}
		log.Info().Msgf("Activating task with ID: %d", id)
		return controllers.ActivateTaskByID(id)
	})

	r.Get("/:task_id/deactivate", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("task_id")
		if err != nil {
			return err
		}
		log.Info().Msgf("Deactivating task with ID: %d", id)
		return controllers.DeactivateTaskByID(id)
	})
}
