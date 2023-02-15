/*Ihre Schleife soll nicht mehr von 1 – 100 zählen, sondern in einem Wertebereich Ihrer 
Wahl. D.h. dass sie über die Tastatur einen Startwert und einen Endwert für die Schleife 
eingaben.*/
package main

import "fmt"

func main() {
	var start int
	var end int
	fmt.Print("Enter the start value: ")
	fmt.Scan(&start)
	fmt.Print("Enter the end value: ")
	fmt.Scan(&end)
	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}
	fmt.Println("The sum of the numbers from", start, "to", end, "is", sum)
}
