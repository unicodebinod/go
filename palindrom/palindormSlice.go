//Legen Sie ein Integer Slice an und prüfen Sie ob es sich um ein Palindrom handelt

package main

import "fmt"

func isPalindrome(slice []int) bool {
    for i := 0; i < len(slice)/2; i++ {
        if slice[i] != slice[len(slice)-1-i] {
            return false
        }
    }
    return true
}

func main() {
    slice1 := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
    if isPalindrome(slice1) {
        fmt.Println("The slice is a palindrome")
    } else {
        fmt.Println("The slice is not a palindrome")
    }

    slice2 := []int{1, 2, 3, 4, 5, 6, 7, 8}
    if isPalindrome(slice2) {
        fmt.Println("The slice is a palindrome")
    } else {
        fmt.Println("The slice is not a palindrome")
    }
}
