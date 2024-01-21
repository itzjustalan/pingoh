package db

import (
	_ "github.com/glebarez/go-sqlite"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func ConnectDB() {
	var err error
	DB, err = sqlx.Connect("sqlite", "./data.db")
	if err != nil {
		panic("error connecting to db!")
	}
	// defer DB.Close()

	// connection pool options
	DB.SetMaxIdleConns(4)

	// create tables
	CreateUsersTable()
}

// db, err := sqlx.Open("sqlite", "./data.db")
// db := sqlx.MustConnect("sqlite", ":memory:")
// db := sqlx.MustConnect("sqlite", "./data.db")

// err := db.Ping()
// if err != nil {
// 	fmt.Println(err)
// }
