package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Birthdate string
	CreatedAt time.Time
}

func (u *User) OutputUserInfo() {
	fmt.Println(u.FirstName, u.LastName, u.Birthdate, u.CreatedAt)
}
func (u *User) ClearUserName() {
	u.FirstName = ""
	u.LastName = ""
}
func New(FirstName, LastName, Birthdate string) (*User, error) {
	if FirstName == "" || LastName == "" || Birthdate == "" {
		return nil, errors.New("ERROR: first name, last name and birth date are required")
	}
	CreatedAt := time.Now()
	return &User{
		FirstName,
		LastName,
		Birthdate,
		CreatedAt,
	}, nil
}
