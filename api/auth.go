package api

import (
	"pingoh/handlers"

	"github.com/gofiber/fiber/v2"
)

func addAuthRoutes(api *fiber.Router) {
	auth := (*api).Group("/auth")

	auth.Post("/signup", func(c *fiber.Ctx) error {
		var b handlers.AuthCredentials
		if err := c.BodyParser(&b); err != nil {
			return fiber.ErrBadRequest
		}
		u, err := handlers.Signup(&b)
		if err != nil {
			return err
		}
		return c.JSON(u)
	})

	auth.Post("/signin", func(c *fiber.Ctx) error {
		var b handlers.AuthCredentials
		if err := c.BodyParser(&b); err != nil {
			return fiber.ErrBadRequest
		}
		u, err := handlers.Signin(&b)
		if err != nil {
			return err
		}
		return c.JSON(u)
	})

	auth.Post("/refresh", func(c *fiber.Ctx) error {
		b := struct {
			Token string `json:"token"`
		}{}
		if err := c.BodyParser(&b); err != nil {
			return fiber.ErrBadRequest
		}
		u, err := handlers.RefreshTokens(b.Token)
		if err != nil {
			return err
		}
		return c.JSON(u)
	})
}
