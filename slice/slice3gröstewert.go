/*Legen Sie ein Slice mit mindestens 10 Werten an und 
geben auf der Konsole die drei größten Werte aus.*/

package main

import (
    "fmt"
    "sort"
)

func main() {
	// Slice mit 10 Werten
    s := []int{5, 8, 2, 4, 9, 12, 1, 0, 15, 3} 
    
    // Sortiere den Slice in absteigender Reihenfolge
    sort.Sort(sort.Reverse(sort.IntSlice(s)))

	//Ausgabe der 3 höchstenn Werte in der Konsole
    //for i := 0; i < len(s) -1 ; i++ { 
     for i := 0; i < 3 ; i++ {
    fmt.Println("die drei größten Werte:", s[i])
    }
}