package main

import (
	"embed"
	"flag"
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

//go:embed all:frontend/dist
var dashboard embed.FS

var (
  PINGOH_PORT = ":3000"
  PINGOH_DB_FILE = "pingoh.db"
  PINGOH_LOG_FILE = "pingoh.log"
)

func main() {
  loadEnvs()
	setLogger(&PINGOH_LOG_FILE)
	db.ConnectDB(&PINGOH_DB_FILE)
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
		PathPrefix: "frontend/dist",
		Browse:     true,
	}))

	app.Listen(PINGOH_PORT)
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

func loadEnvs() {
  if v := os.Getenv("PINGOH_PORT"); v != "" {
    PINGOH_PORT = v
  }
  if v := os.Getenv("PINGOH_DB_FILE"); v != "" {
    PINGOH_DB_FILE = v
  }
  if v := os.Getenv("PINGOH_LOG_FILE"); v != "" {
    PINGOH_LOG_FILE = v
  }

  // maybe in the future recieve a number and create the port string in run time with fmt.Sprintf(":%d", port)
  // and also check if the port is free or something something like that that lol
  flag.StringVar(&PINGOH_PORT, "port", PINGOH_PORT, "port number in :3000 format")
  flag.StringVar(&PINGOH_DB_FILE, "db", PINGOH_DB_FILE, "db file path")
  flag.StringVar(&PINGOH_LOG_FILE, "log", PINGOH_LOG_FILE, "log file path")
  flag.Parse()
}

func setLogger(logfile *string) {
	// zerolog.ErrorFieldName = "e"
	// zerolog.LevelFieldName = "l"
	// zerolog.CallerFieldName = "c"
	// zerolog.MessageFieldName = "m"
	// zerolog.TimestampFieldName = "t"
	// zerolog.ErrorStackFieldName = "s"
	zerolog.TimeFieldFormat = time.RFC3339
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// todo: console only on dev
	file, _ := os.OpenFile(*logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	multiw := zerolog.MultiLevelWriter(consoleWriter, file)
	log.Logger = log.Output(multiw)
	// defer file.Close()
}
