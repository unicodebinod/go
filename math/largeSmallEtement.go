package main

import "fmt"

func main() {
    // Sample slice
    s := []int{3, 5, 2, 7, 1, 8, 4, 6}
    
    // Call the user-defined functions to get the largest and smallest elements
    largest := getLargest(s)
    smallest := getSmallest(s)
    
    fmt.Printf("The largest element in the slice is %d\n", largest)
    fmt.Printf("The smallest element in the slice is %d\n", smallest)
}

// Function to get the largest element in a slice
func getLargest(slice []int) int {
    var largest int
    for _, value := range slice {
        if value > largest {
            largest = value
        }
    }
    return largest
}

// Function to get the smallest element in a slice
func getSmallest(slice []int) int {
    smallest := slice[0]
    for _, value := range slice {
        if value < smallest {
            smallest = value
        }
    }
    return smallest
}

