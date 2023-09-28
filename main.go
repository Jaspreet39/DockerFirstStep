package main

import (
	"fmt"
)

func main() {
	authenticator()
}

// Japsreet's function
func authenticator() {
	var userID int = 500224473
	var password int = 123456
	var username int
	var pass int

	fmt.Println("Please enter the username:")
	fmt.Scanln(&username)

	fmt.Println("Please enter the password:")
	fmt.Scanln(&pass)

	if username == userID && password == pass {
		fmt.Println("you are successfully loggedIn")
	} else {
		fmt.Println("Incorrect credentials")
	}
}
