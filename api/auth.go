package api

import (
	"pingoh/handlers"
	"pingoh/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func addAuthMiddle(api *fiber.Router) {
	(*api).Use(func(c *fiber.Ctx) error {
		v := strings.Split(c.Get("Authorization"), " ")
		if len(v) != 2 || v[0] != "Bearer" {
			return fiber.ErrUnauthorized
		}
		claims, err := services.ValidateToken(v[1])
		if err != nil {
			return err
		}
		c.Locals("uid", claims.UID)
		c.Locals("role", claims.Role)
		c.Locals("access", claims.Access)
		return c.Next()
	})
}

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
