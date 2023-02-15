//to find the largest number with append 
package main

import "fmt"

func main() {
    // Sample slice
    s := []int{3, 5, 2, 7, 1, 8, 4, 6}
    
    // Append a new element to the slice
    s = append(s, 9)
    
    // Call the user-defined function to get the largest element
    largest := getLargest(s)
    
    fmt.Printf("The largest element in the slice is %d", largest)
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
