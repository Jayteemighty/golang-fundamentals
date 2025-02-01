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
}
