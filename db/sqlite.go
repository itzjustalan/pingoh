package db

import (
	_ "github.com/glebarez/go-sqlite"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func panicOnErr(err error, reason string) {
	if err != nil {
		panic(reason + err.Error())
	}
}

func ConnectDB() {
	var err error
	DB, err = sqlx.Connect("sqlite", "./data.db")
	panicOnErr(err, "error connecting to db!")
	// defer DB.Close()

	// connection pool options
	DB.SetMaxIdleConns(4)

	// create tables
	err = createUsersTable()
	panicOnErr(err, "err creating table")
	err = createTasksTable()
	panicOnErr(err, "err creating table")
}

// db, err := sqlx.Open("sqlite", "./data.db")
// db := sqlx.MustConnect("sqlite", ":memory:")
// db := sqlx.MustConnect("sqlite", "./data.db")

// err := db.Ping()
// if err != nil {
// 	fmt.Println(err)
// }
