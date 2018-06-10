package user

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// User model
type User struct {
	name string
	uuid string
}

// Save user to DB
func (u *User) Save() {
	db, err := sql.Open("sqlite3", "./lbmp_users.db")
	if err != nil {
		panic(err)
	}
	stmt, err := db.Prepare("INSERT INTO users(uuid, name) values(?,?)")
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec("123", "testtest")
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println(id)
}
