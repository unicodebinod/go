/* Erweitern Sie den Golang Code so, dass Sie beide Werte der Additon über die Tastatur 
eingeben können.*/
package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)
	result := num1 + num2
	fmt.Println("The result of", num1, "+", num2, "is", result)
}
