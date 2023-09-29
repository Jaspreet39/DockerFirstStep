package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	authenticator()
	ArrayReverse()
	TableGenerator()
	Armstrong()
	greaterNumber()
        fibonacci()
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

func fibonacci() {
	var n int = 5
	t1 := 0
	t2 := 1
	nextTerm := 0

	fmt.Print("Enter the number of terms : ")
	fmt.Print("Fibonacci Series :")
	for i := 1; i <= n; i++ {
		if i == 1 {
			fmt.Print(" ", t1)
			continue
		}
		if i == 2 {
			fmt.Print(" ", t2)
			continue
		}
		nextTerm = t1 + t2
		t1 = t2
		t2 = nextTerm
		fmt.Print(" ", nextTerm)
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

func Armstrong() {
	var number, tempNumber, remainder int
	var result int = 0
	fmt.Print("Enter any three digit number : ")
	fmt.Scan(&number)

	tempNumber = number

	for {
		remainder = tempNumber % 10
		result += remainder * remainder * remainder
		tempNumber /= 10

		if tempNumber == 0 {
			break
		}
	}

	if result == number {
		fmt.Printf("%d is an Armstrong number.", number)
	} else {
		fmt.Printf("%d is not an Armstrong number.", number)
	}
}

func greaterNumber() {
	var num1 float64 = 11.25
	var num2 float64 = 22.14

	var large float64 = 0

	large = math.Max(num1, num2)

	fmt.Println("largest number is : %f", large)
}
