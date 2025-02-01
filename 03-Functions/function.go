package main

import (
	"errors"
	"fmt"
)

// 1. Defining Functions
func sayHello() {
	fmt.Println("Hello")
}

// 2. Function Arguments and Return Values
func greet(name string) {
	fmt.Println("Hello, " + name)
}

func add(x int, y int) int {
	return x + y
}

// Named Return Values
func addAndSubtract(x int, y int) (sum2 int, difference int) {
	sum2 = x + y
	difference = x - y
	return
}

// 3. Multiple Return Values
func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("Cannot divide by 0")
	}
	return x / y, nil
}

// Main
func main() {
	sayHello()
	greet("Joshua")

	sum := add(3, 5)
	fmt.Println("Sum is = ", sum)

	sum2, difference := addAndSubtract(9, 6)
	fmt.Println("The sum is = ", sum2)
	fmt.Println("The difference is = ", difference)

	result, err := divide(12, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("the division is = ", result)
	}
	// Trying division by zero
	_, error := divide(10, 0)
	if error != nil {
		fmt.Println(error)
	}
}
