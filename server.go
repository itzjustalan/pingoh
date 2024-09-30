package main

import (
	"embed"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"pingoh/controllers"
	"pingoh/db"
	"pingoh/routes"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// old way -> go:embed all:frontend/dist

//go:embed frontend/dist/*
var dashboard embed.FS

var (
	PINGOH_PORT           = ":3000"
	PINGOH_DB_FILE        = "pingoh.db"
	PINGOH_LOG_FILE       = "pingoh.log"
	PINGOH_ADMIN_EMAIL    = "admin@mail.com"
	PINGOH_ADMIN_PASSWORD = "password"
)

func main() {
	loadEnvs()
	setLogger(&PINGOH_LOG_FILE)
	db.ConnectDB(&PINGOH_DB_FILE)
	go listenAndStop()
	controllers.Signup(&controllers.SignupCredentials{
		Name:  "Admin",
		Email: PINGOH_ADMIN_EMAIL,
		Passw: PINGOH_ADMIN_PASSWORD,
	})
	controllers.StartTasks()
	app := fiber.New(fiber.Config{
		ServerHeader:             "Pingoh",
		AppName:                  "Pingoh",
		EnableSplittingOnParsers: true,
	})

	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	routes.AddRoutes(app)

	app.Use("/", filesystem.New(filesystem.Config{
		Root:         http.FS(dashboard),
		NotFoundFile: "frontend/dist/index.html",
		PathPrefix:   "frontend/dist",
		Browse:       true,
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
	if v := os.Getenv("PINGOH_ADMIN_EMAIL"); v != "" {
		PINGOH_ADMIN_EMAIL = v
	}
	if v := os.Getenv("PINGOH_ADMIN_PASSWORD"); v != "" {
		PINGOH_ADMIN_PASSWORD = v
	}

	// maybe in the future recieve a number and create the port string in run time with fmt.Sprintf(":%d", port)
	// and also check if the port is free or something something like that that lol
	flag.StringVar(&PINGOH_PORT, "port", PINGOH_PORT, "port number in :3000 format")
	flag.StringVar(&PINGOH_DB_FILE, "db", PINGOH_DB_FILE, "db file path")
	flag.StringVar(&PINGOH_LOG_FILE, "log", PINGOH_LOG_FILE, "log file path")
	flag.StringVar(&PINGOH_ADMIN_EMAIL, "email", PINGOH_ADMIN_EMAIL, "admin user email")
	flag.StringVar(&PINGOH_ADMIN_PASSWORD, "password", PINGOH_ADMIN_PASSWORD, "admin user password")
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
