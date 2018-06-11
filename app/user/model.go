package user

import (
	"database/sql"
	"fmt"

	// needed for sql driver
	_ "github.com/mattn/go-sqlite3"

	uuid "github.com/satori/go.uuid"
)

// User model
type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	uuid      uuid.UUID
}

// Save user to DB
func (u *User) Save() {
	db, err := sql.Open("sqlite3", "./lbmp_users.db")
	if err != nil {
		panic(err)
	}
	stmt, err := db.Prepare("INSERT INTO users(uuid, username) values(?,?)")
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(&u.uuid, &u.Firstname)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println(id)
}
