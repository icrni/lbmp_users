package user

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// CreateUser based on given data
func (u *User) CreateUser() {
	uuid4, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	u.uuid = uuid4
	u.Save()
}
