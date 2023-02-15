/*Sortieren Sie ein Slice mit mindestens 10 Werten in umgekehrter Reihenfolge. Geben Sie die 
umgekehrte Reihenfolge auf der Konsole aus.*/
package main

import (
    "fmt"
    "sort"
)

func main() {
    s := []int{3, 6, 2, 8, 1, 9, 4, 5, 7, 0} // Beispiel-Slice mit 10 Werten
    fmt.Println("Unsortierter Slice:", s)

    // Sortiere den Slice in umgekehrter Reihenfolge
    sort.Sort(sort.Reverse(sort.IntSlice(s)))

    fmt.Println("Sortierter Slice in umgekehrter Reihenfolge:", s)
}
