package db

import (
	"database/sql/driver"
	"time"
)

type UserRole string
type UserAccess []string

const (
	UserRoleAdmin UserRole = "admin"
	UserRoleUser  UserRole = "user"
)

type User struct {
	ID        int        `json:"-"`
	// UID       string     `json:"uid"`
	Email     string     `json:"email"`
	PwHash    string     `json:"-" db:"pw_hash"`
	Role      UserRole   `json:"role"`
	Access    UserAccess `json:"access"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
}

func (v UserAccess) Value() (driver.Value, error) {
	return jsonSqlValuer(v)
}

func (v *UserAccess) Scan(src interface{}) error {
	return jsonSqlScanner(v, src)
}

func createUsersTable() error {

	// uid TEXT UNIQUE NOT NULL,
	// name VARCHAR(255) NOT NULL,
	// email VARCHAR(255) UNIQUE NOT NULL,
	q := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT UNIQUE NOT NULL,
		pw_hash TEXT NOT NULL,
		role TEXT NOT NULL DEFAULT 'user' CHECK ( role IN ('admin','user') ),
		access TEXT NOT NULL DEFAULT "[]",
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	) 
	`
	_, err := DB.Exec(q)
	return err
}

var adminCreated = false

func potentialAdmin() (UserRole, error) {
	if adminCreated {
		return UserRoleUser, nil
	}
	var count int
	err := DB.Get(&count, `SELECT count(*) FROM users`)
	if err != nil {
		return UserRoleUser, err
	}
	if count == 0 {
		adminCreated = true
		return UserRoleAdmin, nil
	}
	return UserRoleUser, nil
}

func CreateUser(u *User) (int64, error) {
	q := `
	INSERT INTO users (email, pw_hash, role) VALUES (?, ?, ?)
	`
	if u.Role == "" {
		role, err := potentialAdmin()
		if err != nil {
			return 0, err
		}
		u.Role = role
	}
	res, err := DB.Exec(q, u.Email, u.PwHash, u.Role)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func FindUserByID(id int) (User, error) {
	q := `
	SELECT * FROM users WHERE id = ?
	`
	var u User
	err := DB.Get(&u, q, id)
	return u, err
}

func FindUserByEmail(email string) (User, error) {
	q := `
	SELECT * FROM users WHERE email = ?
	`
	var u User
	err := DB.Get(&u, q, email)
	return u, err
}

// rows, err := db.Queryx(`SELECT * FROM users`)
// if err != nil {
// 	fmt.Println(err)
// }
// defer rows.Close()

// for rows.Next() {
// 	var u User
// 	err = rows.StructScan(&u)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(u.ID, u.Email, u.CreatedAt.GoString())
// }

// var users []User
// err = db.Select(&users, `SELECT * FROM users`)
// if err != nil {
// 	fmt.Println(err)
// }
// fmt.Println("users count: ", len(users))
