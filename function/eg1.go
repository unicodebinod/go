/*
Function syntax
----------------
    func   function_name  (Parameter-list)     (Return_type){
eg:	func      area        (length, width int)   int{
    // function body.....
}


*/

package main
import "fmt"


func area(length, width int)int{
	
	A := length* width
	return A
}

// Main function
func main() {
fmt.Printf("Area of rectangle is : %d", area(12, 10))
}
