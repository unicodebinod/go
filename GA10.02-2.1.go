//Erweitern Sie den Ihren Code um die Ausgabe der Summe aus dem Schleifendurchlauf
package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("The sum of the numbers from 1 to 100 is", sum)
}
