package admin

import (
	"errors"
	"fmt"

	"github.com/adfharrison1/go-learning-api/structs/user"
)

type Admin struct {
	email    string
	password string
	*user.User
}

func New(email, password, firstName, lastName, birthdate string) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("ERROR: email and password are required")
	}
	user, error := user.New(
		firstName,
		lastName,
		birthdate,
	)
	if error != nil {
		return nil, errors.New(error.Error())
	}
	return &Admin{
		email,
		password,
		user,
	}, nil
}
func (a *Admin) OutputAdminInfo() {
	fmt.Println(a.email, a.password)
}
