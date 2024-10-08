package routes

import (
	"pingoh/controllers"
	"pingoh/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func addAuthChecks(api *fiber.Router) {
	(*api).Use(func(c *fiber.Ctx) error {
		token := c.Query("token")
		if token == "" {
			v := strings.Split(c.Get("Authorization"), " ")
			if len(v) != 2 || v[0] != "Bearer" {
				return fiber.ErrUnauthorized
			}
			token = v[1]
		}
		claims, err := services.ValidateToken(token)
		if err != nil {
			return fiber.ErrUnauthorized
		}
		c.Locals("uid", claims.ID)
		c.Locals("role", claims.Role)
		c.Locals("access", claims.Access)
		return c.Next()
	})
}

func addAuthRoutes(api *fiber.Router) {
	auth := (*api).Group("/auth")

	auth.Post("/signup", func(c *fiber.Ctx) error {
		var b controllers.SignupCredentials
		if err := c.BodyParser(&b); err != nil {
			return fiber.ErrBadRequest
		}
		err := controllers.Validator.Struct(&b)
		if err != nil {
			return err
		}
		u, err := controllers.Signup(&b)
		if err != nil {
			return err
		}
		return c.JSON(u)
	})

	auth.Post("/signin", func(c *fiber.Ctx) error {
		var b controllers.SigninCredentials
		if err := c.BodyParser(&b); err != nil {
			return fiber.ErrBadRequest
		}
		err := controllers.Validator.Struct(&b)
		if err != nil {
			return err
		}
		u, err := controllers.Signin(&b)
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
		u, err := controllers.RefreshTokens(b.Token)
		if err != nil {
			return err
		}
		return c.JSON(u)
	})
}
