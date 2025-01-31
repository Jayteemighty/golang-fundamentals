package main

import "fmt"

func main() {
	x := 10

	//Basic If statement
	if x > 6 {
		fmt.Println("x is greater than 6")
	}

	// Basic if-else statement
	if x < 6 {
		fmt.Println("x is less than 6")
	} else {
		fmt.Println("x is greater than or equal to 6")
	}

	// if-else else if statement
	if x < 5 {
		fmt.Println("wowww!")
	} else if x == 5 {
		fmt.Println("wazo!")
	} else {
		fmt.Println("werey!")
	}

}
