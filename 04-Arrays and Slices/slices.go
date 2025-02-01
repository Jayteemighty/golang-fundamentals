package main

import "fmt"

// sliceExample demonstrates working with slices.
func sliceExample() {
	// Creating a slice using a literal.
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Initial slice:", numbers)
	fmt.Printf("Length: %d, Capacity: %d\n", len(numbers), cap(numbers))

	// Appending elements to the slice.
	numbers = append(numbers, 6, 7, 8)
	fmt.Println("After appending:", numbers)
	fmt.Printf("Length: %d, Capacity: %d\n", len(numbers), cap(numbers))

	// Creating a sub-slice.
	subset := numbers[2:5] // Elements at indices 2, 3, and 4.
	fmt.Println("Sub-slice (indices 2 to 4):", subset)
}

func main() {
	sliceExample()
}
