// Variables
package main
import "fmt"

func main(){
	// Using var to declare variables
	var name string = "Joshua"
	var num int = 24
	var pie float64 = 3.14 // also float32

	// Shorthand syntax (only inside a function)
	name2 := "Joshhy"
	age2 := 24
	grownMan := true

	// Default Values
	var emptyNum int
	var emptyStr string

	// Type Inference
	weight := 70 // Inferred as int
	
	fmt.Println(emptyStr, emptyNum, weight) // every variable or import has to be used or ypu would get errors
	fmt.Println(name, num, pie, name2, age2, grownMan)
}