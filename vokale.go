/*Schreiben Sie ein Golang Programm, das eine Zeichenkette einliest und die Anzahl der Vokale und 
Konsonanten in dieser Zeichenkette berechnet.*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Gib das zu zählende Vokal ein:")

	//The character string(Zeichenkette) is read from the input
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Counts the total number of vowel and consonant 
	AnzahlVokal := 0
	AnzahlKonstant := 0
	for _, ch := range input {
		switch strings.ToLower(string(ch)) {
		case "a", "e", "i", "o", "u":
			AnzahlVokal++
		default:
			if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
				AnzahlKonstant++
			}
		}
	}

	// output
	fmt.Println("Anzahl der Vokale:", AnzahlVokal)
	fmt.Println("Anzahl der Konsonanten:", AnzahlKonstant)

}
