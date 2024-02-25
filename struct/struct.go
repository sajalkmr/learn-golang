package main

import "fmt"

// Define a struct named Person
type Person struct {
	Name string
	Age  int
}

func main() {
	// Initialize a Person struct
	person := Person{
		Name: "John",
		Age:  30,
	}

	// Access and print fields of the struct
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)

	// Update a field of the struct
	person.Age = 35
	fmt.Println("Updated Age:", person.Age)

	// Update a field using the original struct
	person.Name = "Alice"
	fmt.Println("Updated Name:", person.Name)
}
