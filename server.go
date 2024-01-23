package main

import (
	"embed"
	"net/http"
	"os"
	"os/signal"
	"pingoh/api"
	"pingoh/db"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

//go:embed all:frontend/build
var dashboard embed.FS

func main() {
	db.ConnectDB()
	go listenAndStop()
	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	api.AddRoutes(app)

	app.Use("/", filesystem.New(filesystem.Config{
		// Root:       http.FS(frontend.BuildDir),
		Root:       http.FS(dashboard),
		PathPrefix: "frontend/build",
		Browse:     true,
	}))

	app.Listen(":3000")
}

func listenAndStop() {
	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		syscall.SIGINT, // for *nix systems
		// syscall.SIGKILL, // for nt systems
		// syscall.SIGTERM, // just cause
	)
	<-c
	// app.Shutdown()
	os.Exit(1)
}
