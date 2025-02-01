package main

import "fmt"

// RANGE-BASED LOOP
// used to iterate over elements in arrays, slices, maps, and strings
//syntax:
//for index, value  := range collection{
// Use index and value
//}

func main() {
	// 1. Iterating Over Arrays and Slices
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for index, value := range numbers {
		fmt.Println("index: ", index, "value: ", value)
	}

	// 2. Iterating Over Maps
	this_map := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
		"key4": "value4",
	}
	for key, value := range this_map {
		fmt.Println("key: ", key, "value: ", value)
	}

	// 3. Iterating Over Strings
	this_str := "crazy is the head that wears the crown"
	for index, runeValue := range this_str {
		fmt.Printf("Index: %d, Rune: %c\n", index, runeValue)
	}

	// 4. Ignoring Index or Value
	numbers2 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	//ignore the index
	for _, value := range numbers2 {
		fmt.Println(value)
	}
	// Ignore the value
	for index := range numbers {
		fmt.Println(index)
	}
}
