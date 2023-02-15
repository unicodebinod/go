/*Erstellen Sie einen Golang Code in VSCode das im Terminal das Ergebnis einer Addition 
von zwei Zahlen ausgibt*/
package main

import "fmt"

func main() {

	var number1 = 10 
	var number2 = 20
	var number3 = number1 +number2	

   // printing the results
	fmt.Println("The addition of ", number1, " and ", number2, " is \n ", number3, "\n(Addition of two integers within the function)")
}
