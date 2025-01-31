package main

import "fmt"

func main() {
	day := "monday"
	var x int = 4

	// Value Matching
	switch day {
	case "monday":
		fmt.Println("Today is Monday, bitch! Pay me!!!")
	case "tuesday":
		fmt.Println("Today is Tuesday, bitch! Pay me!!!")
	default:
		fmt.Println("I don't know what day it is, bitch!")
	}

	// Expression Matching
	switch {
	case x == 3:
		fmt.Println("The answer is not 3")
	case x == 2:
		fmt.Println("Dumbass, answer is not 2")
	case x == 4:
		fmt.Println("Apparently, you are not a Dumbass")
	default:
		fmt.Println("Give up")
	}
}
