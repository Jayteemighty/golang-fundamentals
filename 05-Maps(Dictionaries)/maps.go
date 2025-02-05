package main

import (
	"fmt"
)

// basicMapExample demonstrates creating and manipulating a map with string keys and int values.
func basicMapExample() {
	fmt.Println("=== Basic Map Example ===")

	// Creating and initializing a map using a literal.
	fruitCounts := map[string]int{
		"apple":  5,
		"banana": 10,
	}

	// Accessing a value in the map.
	fmt.Println("Initial apple count:", fruitCounts["apple"])

	// Adding a new key-value pair.
	fruitCounts["orange"] = 15
	fmt.Println("After adding orange:", fruitCounts)

	// Updating an existing key-value pair.
	fruitCounts["apple"] = 7
	fmt.Println("After updating apple:", fruitCounts)

	// Deleting an entry using the delete function.
	delete(fruitCounts, "banana")
	fmt.Println("After deleting banana:", fruitCounts)

	// Checking if a key exists using the comma-ok idiom.
	if count, ok := fruitCounts["banana"]; ok {
		fmt.Println("Banana count:", count)
	} else {
		fmt.Println("Banana key does not exist.")
	}
}

// heterogenousMapExample demonstrates how to store multiple data types in a map using interface{}.
func heterogenousMapExample() {
	fmt.Println("=== Heterogeneous Map Example (using interface{}) ===")

	// Creating a map with string keys and values of type interface{} to hold different types.
	data := map[string]interface{}{
		"name":    "Alice",
		"age":     30,
		"married": true,
	}

	// Accessing values with type assertions.
	if name, ok := data["name"].(string); ok {
		fmt.Println("Name:", name)
	}

	// Iterating over the map using a type switch to handle different value types.
	for key, value := range data {
		fmt.Printf("Key: %s, ", key)
		switch v := value.(type) {
		case string:
			fmt.Printf("Value (string): %s\n", v)
		case int:
			fmt.Printf("Value (int): %d\n", v)
		case bool:
			fmt.Printf("Value (bool): %t\n", v)
		default:
			fmt.Println("Value of unknown type")
		}
	}
}

func main() {
	// Run the basic map example.
	basicMapExample()

	fmt.Println() // Print an empty line for better readability.

	// Run the heterogeneous map example.
	heterogenousMapExample()
}
