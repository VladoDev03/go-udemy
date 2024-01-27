package main

import (
	"fmt"
	"structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (DD/MM/YYYY): ")

	var appUser user.User

	// appUser = user{}

	// appUser = user{
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthdate,
	// 	time.Now(),
	// }

	// appUser = user.User{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthdate: userBirthdate,
	// 	createdAt: time.Now(),
	// }

	// outputUserDetails(appUser)
	// outputUserDetailsPointer(&appUser)
	appUser.OutputUserDetails()

	appUser.ClearUserName()
	appUser.OutputUserDetails()

	// newAppUser, err := user.NewUser(userFirstName, userLastName, userBirthdate)
	newAppUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	newAppUser.OutputUserDetails()

	admin := user.NewAdmin("test@test.test", "test")
	admin.User.OutputUserDetails()
	admin.OutputUserDetails()
}

// func outputUserDetails(u user.User) {
// 	fmt.Println(u.firstName, u.lastName, u.birthdate)
// }

// func outputUserDetailsPointer(u *user.User) {
// 	fmt.Println((*u).firstName, (*u).lastName, (*u).birthdate)
// 	fmt.Println(u.firstName, u.lastName, u.birthdate)
// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
