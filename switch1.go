/* Schreiben Sie einen Go Code, dass zwischen zwei Ereignissen unterscheidet:
Fall 1: Der Wert welcher über die Tastatur eingegeben wird ist größer gleich 18 
Auszugebender Text: „Du darfst alles kaufen was du willst!“
Fall 2: Der Wert welcher über die Tastatur eingegeben wird ist kleiner als 18 größer als 7
Auszugebender Text: „Du darfst nur im Rahmen deines Taschengeldes einkaufen und mit 
Erlaubnis deiner Eltern!“
Fall 3: Der Wert welcher über die Tastatur eingegebn wird ist kleiner als 7 und größer als 0
Auszugebender Text: „Du darfst leider nichts kaufen, was dir nicht deine Eltern besorgen!“*/
package main

import "fmt"

func main() {
	var age int
	fmt.Print("Bitte geben Sie Ihr Alter ein: ")
	fmt.Scanf("%d", &age)
	switch {
	case age >= 18:
		fmt.Println("Du darfst alles kaufen was du willst!")
	case age >= 7 && age < 18:
		fmt.Println("Du darfst nur im Rahmen deines Taschengeldes einkaufen und mit Erlaubnis deiner Eltern!")
	case age > 0 && age < 7:
		fmt.Println("Du darfst leider nichts kaufen, was dir nicht deine Eltern besorgen!")
	default:
		fmt.Println("Ungültiger Alterseingabe.")
	}
}
