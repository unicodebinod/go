package main

import "fmt"

// defining the struct
type Car struct {
	Name, Model, Color string
	WeightInKg		 float64
}

// Main Function
func main() {
	c := Car{Name: "VM", Model: "Golf",
			Color: "Blue", WeightInKg: 2020}

	fmt.Println("Car Name: ", c.Name)
	fmt.Println("Car Color: ", c.Color)

	// Assigning a new value to a struct field
	c.Color = "Black"
	fmt.Println("Car: ", c)
}

