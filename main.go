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
	var idInput int = 500224330
	var passInput int = 123456

	if idInput == userID && passInput == password {
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

	fmt.Print("Fibonacci Series of 5 terms :")
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

	num := []int{50, 90, 30, 10, 50}
	fmt.Println("Initial Integer Array: ")
	fmt.Println(num)
	fmt.Println("Integer Reverse Sort")
	sort.Sort(sort.Reverse(sort.IntSlice(num)))
	fmt.Println(num)

	fmt.Println()

	text := []string{"Japan", "UK", "Germany", "Australia", "Pakistan"}
	fmt.Println("Initial String Array: ")
	fmt.Println(text)
	fmt.Println("String Reverse Sort")
	sort.Sort(sort.Reverse(sort.StringSlice(text)))
	fmt.Println(text)

}

func TableGenerator() {
	var n int = 5
	fmt.Println("Table of 5: ")
	i := 1
	for {
		if i > 10 {
			break
		}
		fmt.Println(n, " X ", i, " = ", n*i)
		i++
	}
}

func Armstrong() {
	var tempNumber, remainder int
	var result int = 0
	var number int = 153
	fmt.Print("To check whether 153 in Armstrong or not: ")

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

	fmt.Printf("larger number among 11.25 & 22.14 is : %f", large)
}
