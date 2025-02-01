package main

import "fmt"

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

// Main
func main() {
	sayHello()
	greet("Joshua")

	sum := add(3, 5)
	fmt.Println("Sum is = ", sum)

	sum2, difference := addAndSubtract(9, 6)
	fmt.Println("The sum is = ", sum2)
	fmt.Println("The difference is = ", difference)
}
