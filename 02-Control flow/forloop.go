package main

import "fmt"

func main() {
	// Regular FOR LOOP
	for i := 0; i < 10; i++ {
		fmt.Println("These are the numbers:", i)
	}

	// While loop
	x := 10
	for x < 10 {
		fmt.Println("These are the numbers: ", x)
		x++
	}

	// RANGE-BASED LOOP
	// used to iterate over elements in arrays, slices, maps, and strings
	//syntax:
	//for index, value  := range collection{
	// Use index and value
	//}

	// 1. Iterating Over Arrays and Slices
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for index, value := range numbers {
		fmt.Println("index: ", index, "value: ", value)
	}
}
