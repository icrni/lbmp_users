package user

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// CreateUser based on given data
func CreateUser(name string) {
	user := User{}
	uuid4, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	user.Name = name
	user.uuid = uuid4
	user.Save()
}
