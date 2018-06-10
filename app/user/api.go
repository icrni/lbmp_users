package user

// CreateUser based on given data
func CreateUser() {
	user := User{}
	user.name = "test"
	user.uuid = "1231"
	user.Save()
}
