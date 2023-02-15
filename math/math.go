/*Schreiben Sie ein Golang Programm, das eine Liste von Zahlen einliest und die Summe, das 
Durchschnitt und die größte und kleinste Zahl in dieser Liste berechnet.*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var numbers []int
	var sum int
	var min int
	var max int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a list of numbers, separated by spaces:")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	strNumbers := strings.Split(input, " ")

	for i, strNum := range strNumbers {
		number, err := strconv.Atoi(strNum)
		if err != nil {
			fmt.Println("Invalid input. Only numbers are allowed.")
			return
		}

		numbers = append(numbers, number)
		sum += number

		if i == 0 {
			min = number
			max = number
		} else {
			if number > max {
				max = number
			}
			if number < min {
				min = number
			}
		}
	}

	avg := float64(sum) / float64(len(numbers))

	fmt.Println("Sum:", sum)
	fmt.Println("Average:", avg)
	fmt.Println("Minimum:", min)
	fmt.Println("Maximum:", max)
}

