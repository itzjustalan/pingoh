package main

import (
	"embed"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"pingoh/api"
	"pingoh/db"
	"pingoh/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//go:embed all:frontend/build
var dashboard embed.FS

func main() {
	setLogger()
	db.ConnectDB()
	go listenAndStop()
	app := fiber.New()
	handlers.StartTasks()

	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

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

func setLogger() {
	// zerolog.ErrorFieldName = "e"
	// zerolog.LevelFieldName = "l"
	// zerolog.CallerFieldName = "c"
	// zerolog.MessageFieldName = "m"
	// zerolog.TimestampFieldName = "t"
	// zerolog.ErrorStackFieldName = "s"
	zerolog.TimeFieldFormat = time.RFC3339
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// todo: console only on dev
	file, _ := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	multiw := zerolog.MultiLevelWriter(consoleWriter, file)
	log.Logger = log.Output(multiw)
	// defer file.Close()
}
