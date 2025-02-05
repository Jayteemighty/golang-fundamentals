package main

import (
	"fmt"
	"strings"
)

// Person is a simple struct representing a person with a name and age.
type Person struct {
	Name string
	Age  int
}

// Greet is a method with a value receiver. It returns a greeting message.
// Since it uses a value receiver, it works on a copy of the struct and does not modify the original.
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old.", p.Name, p.Age)
}

// UpdateName is a method with a pointer receiver.
// Using a pointer receiver allows the method to modify the original struct.
func (p *Person) UpdateName(newName string) {
	// Trim spaces and update the Name field.
	p.Name = strings.TrimSpace(newName)
}

// HaveBirthday is another method with a pointer receiver.
// It increments the Age of the person by one.
func (p *Person) HaveBirthday() {
	p.Age++
}

// DisplayInfo is a method with a value receiver that prints the struct information.
func (p Person) DisplayInfo() {
	fmt.Printf("Person Info -> Name: %s, Age: %d\n", p.Name, p.Age)
}

// main demonstrates creating instances of Person, calling methods,
// and understanding the difference between value and pointer receivers.
func main() {
	// Create an instance of Person using a struct literal.
	person1 := Person{
		Name: "Alice",
		Age:  30,
	}

	// Call a method with a value receiver.
	fmt.Println(person1.Greet())

	// Display initial info.
	person1.DisplayInfo()

	// Call a method with a pointer receiver to update the name.
	// Using &person1 to pass a pointer to the struct.
	person1.UpdateName("  Alice Johnson  ")
	fmt.Println("After updating name:")
	person1.DisplayInfo()

	// Call a method to increment the age.
	person1.HaveBirthday()
	fmt.Println("After having a birthday:")
	person1.DisplayInfo()

	// Demonstrating that methods with pointer receivers can also be called on struct values.
	// Go automatically takes the address when needed.
	person1.HaveBirthday()
	fmt.Println("After another birthday:")
	person1.DisplayInfo()

	// Create another instance using the 'new' keyword.
	person2 := new(Person)
	person2.Name = "Bob"
	person2.Age = 25

	fmt.Println("\nAnother person:")
	person2.DisplayInfo()
	fmt.Println(person2.Greet())
}
