package main

import (
	"fmt"
 
)

/*func main() {
    // Create a new rectangle with width 5 and height 10
    r := rectangle{width: 5, height: 10}

    Calculate the area and perimeter of the rectangle
    area := r.area()
    perimeter := r.perimeter()

    // Print the results
    fmt.Printf("Rectangle with width %f and height %f\n", r.width, r.height)
    fmt.Printf("Area: %f\n", area)
    fmt.Printf("Perimeter: %f\n", perimeter)
    */

    func main() {
        // Create a new rectangle with width 5 and height 10
        r := Rectangle{Width: 5, Height: 10}
        fmt.Println("Die flache betragt", r.Area())
        fmt.Println("Der umfang betragt:", r.Perimeter())
}
