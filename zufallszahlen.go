package main

import ( //für mehreren Packages ("")
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	//eingabe forderung für den start eines wertbereichs
	fmt.Println("git bitte den Startwert des Wertebereich an:")
	reader := bufio.NewReader(os.Stdin)
	startRange, _ := reader.ReadString('\r') // _ = anonym platzhalter für
	startRange = startRange[:len(startRange)-1]
	startInt, _ := strconv.Atoi(startRange) //Ascii to intger

	//einegabeaufforderung für Endwert eines Wertbereichs
	fmt.Print("git bitte den EndWert des Wertbereiche an")
	reader = bufio.NewReader(os.Stdin)
	endRange, _ := reader.ReadString('\r')
	endRange = endRange[:len(endRange)-1] //[: nim alles von array]
	endInt, _ := strconv.Atoi(endRange)

	//zufallsgenerator initialisieren
	rand.Seed(time.Now().UnixNano())

	//[startPosition:EndPosition:schritte]
	//(1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

	//Array zur speicherung unserer Zufallszahlen
	var randomNumbers []int

	//schleife
	for i := 0; i <= 10; i++ {
		randomNum := rand.Intn(endInt-startInt+1) + startInt
		randomNumbers = append(randomNumbers, randomNum)

		//Ausgabe von Array randomnumbers(Ausgabe der Zufallszahlen)
		fmt.Print("meine zufallszahlen sind:", randomNumbers)
	}
}

//'\'This allows you to overwrite the current line with new text, rather than printing new text on a new line.