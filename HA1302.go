/*Nehmen Sie den Code zur Hand, welchen wir heute zusammen im Kurs erstellt haben. Die forSchleife, die das leere Array mit Zufallszahlen befüllen soll, bekommt im Augenblick über die Schleife 
die Anzahl der zu generierenden Zufallszahlen geliefert. Ändern Sie den Code so um, dass der Wert 
der gewünschten Zufallszahlen nicht im Code steht sondern über eine Tastatureingabe eingegeben 
werden muss.*/

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main(){
	// Eingabeaufforderung fuer den Startwert eines Wertebereichs
	
	fmt.Println("Gib bitte den Starwert des Wertebereich an: ")
	reader := bufio.NewReader(os.Stdin)
	startRange, _ := reader.ReadString('\r')
	startRange = startRange[:len(startRange)-1]
	startInt, _ := strconv.Atoi(startRange)

	// Eingabeaufforderung fuer den Endwert im Wertebereichs

	fmt.Println("Gib bitte das Ende des Wertebereichs an: ")
	reader = bufio.NewReader(os.Stdin)
	endRange, _ := reader.ReadString('\r')
	endRange = endRange[:len(endRange)-1]
	endInt, _ := strconv.Atoi(endRange)

	// Zufallsgenerator initialisieren 
	rand.Seed(time.Now().UnixNano())

	// Array zur Speicherung unserer Zufallszahlen 
	var randomNumbers []int


/*
	Entweder so direkt über die Scan funktion einscannen und durchlaufen lassen

	// Anzahl der Zufallszahlen 	
	fmt.Println("gib an wie Viele Zufallszahlen")
	var howmuch int
	fmt.Scan(&howmuch)

	// Schleife zum Speichern und erzeugen der Zufallszahlen 
		for i := 0 ; i < howmuch ; i++{
		randomNum := rand.Intn(endInt-startInt+1) + startInt
		randomNumbers = append(randomNumbers, randomNum)
	} 
*/

////// Oder Alternativ was wir im Unterricht gemacht haben  mit dem bufio und der speicherung ////////// 

	// Eingabeaufforderung fuer den Endwert im Wertebereichs

	fmt.Println("gib an wie Viele Zufallszahlen : ")
	reader = bufio.NewReader(os.Stdin)
	countRange, _ := reader.ReadString('\r')
	countRange = countRange[:len(countRange)-1]
	countNumber, _ := strconv.Atoi(countRange)


	// Schleife zum Speichern und erzeugen der Zufallszahlen 
	for i := 0 ; i < countNumber ; i++{
		randomNum := rand.Intn(endInt-startInt+1) + startInt
		randomNumbers = append(randomNumbers, randomNum)
	} 
	

	// Ausgabe von Array randomNumbers (Ausgabe der Zufallszahlen)
	fmt.Println("Meine Zufallszahlen sind : ", randomNumbers)


}