package main

import (
	"fmt"
	"sort"
)

func main() {
	authenticator()
	ArrayReverse()
	TableGenerator()
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

func ArrayReverse() {

	fmt.Println("Interger Reverse Sort")
	num := []int{50, 90, 30, 10, 50}
	sort.Sort(sort.Reverse(sort.IntSlice(num)))
	fmt.Println(num)

	fmt.Println()

	fmt.Println("String Reverse Sort")
	text := []string{"Japan", "UK", "Germany", "Australia", "Pakistan"}
	sort.Sort(sort.Reverse(sort.StringSlice(text)))
	fmt.Println(text)

}

func TableGenerator() {
	var n int
	fmt.Print("Enter any Integer Number : ")
	fmt.Scan(&n)
	i := 1
	/*     For loop as a Go's While     */
	for {
		if i > 10 {
			break
		}
		fmt.Println(n, " X ", i, " = ", n*i)
		i++
	}
}
