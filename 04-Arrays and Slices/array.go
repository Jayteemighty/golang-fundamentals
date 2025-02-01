package main

import "fmt"

// arrayExample demonstrates working with arrays.
func arrayExample() {
	// Declaring and initializing an array.
	arr := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array:", arr)

	// Iterating over the array.
	for i, v := range arr {
		fmt.Printf("arr[%d] = %d\n", i, v)
	}
}

func main() {
	arrayExample()
}
