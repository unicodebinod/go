/*Schreiben Sie ein Golang Programm, das eine Zeichenkette einliest und überprüft, ob es ein Palindrom 
ist*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a string:")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if isPalindrome(input) {
		fmt.Println("The string is a palindrome.")
	} else {
		fmt.Println("The string is not a palindrome.")
	}
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {   //len(s)/2 : modullo= 0 -->true--->arbeitet in hintergrund mit ascii tablle mit werten 	
	if s[i] != s[len(s)-i-1] {
	 return false
		}
	}
	return true
}

