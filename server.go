package main

import (
	"embed"
	"net/http"
	"pingoh/api"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

//go:embed all:frontend/build
var dashboard embed.FS

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://127.0.0.1:3000",
	}))

	app.Get("/h", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api.AddRoutes(app)

	app.Use("/", filesystem.New(filesystem.Config{
		// Root:       http.FS(frontend.BuildDir),
		Root:       http.FS(dashboard),
		PathPrefix: "frontend/build",
		Browse:     true,
	}))

	app.Listen(":3000")
}
