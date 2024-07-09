package db

import (
	_ "github.com/glebarez/go-sqlite"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

var DB *sqlx.DB

func panicOnErr(err error, reason string) {
	if err != nil {
		log.Panic().Msg(reason + " " + err.Error())
	}
}

func ConnectDB() {
	var err error
	DB, err = sqlx.Connect("sqlite", "./pingoh.db")
	panicOnErr(err, "error connecting to db!")
	// defer DB.Close()

	// connection pool options
	DB.SetMaxIdleConns(4)

	// create tables
	err = createUsersTable()
	panicOnErr(err, "err creating table")
	err = createTasksTable()
	panicOnErr(err, "err creating table")
	err = createHttpTasksTable()
	panicOnErr(err, "err creating table")
	err = createHttpAuthsTable()
	panicOnErr(err, "err creating table")
	err = createHttpResultTable()
	panicOnErr(err, "err creating table")
}

// db, err := sqlx.Open("sqlite", "./data.db")
// db := sqlx.MustConnect("sqlite", ":memory:")
// db := sqlx.MustConnect("sqlite", "./data.db")
