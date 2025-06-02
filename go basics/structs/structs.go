package main

import (
	"fmt"

	"github.com/adfharrison1/go-learning-api/structs/admin"
	"github.com/adfharrison1/go-learning-api/structs/user"
)

func main() {
	// ... do something awesome with that gathered data!
	appUser, userError := user.New(
		getUserData("Please enter your first name: "),
		getUserData("Please enter your last name: "),
		getUserData("Please enter your birthdate (MM/DD/YYYY): "),
	)

	adminUser, adminError := admin.New(
		"me@you.com",
		"password",
		appUser.FirstName,
		appUser.LastName,
		appUser.Birthdate,
	)

	if userError != nil || adminError != nil {
		fmt.Println(userError, adminError)
		return
	}

	appUser.OutputUserInfo()
	adminUser.OutputAdminInfo()
	appUser.ClearUserName()
	appUser.OutputUserInfo()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
