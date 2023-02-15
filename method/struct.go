// method with struct 
package main

import "fmt"

// workers structure
type workers struct {
	 name	 string
	 branch string
	 salary int
}

// Method with a receiver of author type
func (a workers) show() {

	fmt.Println("Author's Name: ", a.name)
	fmt.Println("Branch Name: ", a.branch)
	fmt.Println("Salary: ", a.salary)
}

// Main function
func main() {

	// Initializing the values
	// of the author structure
	employ := workers{ name:	"Binod",branch: "DevOps",salary: 40000,}

	// Calling the method
	employ.show()
}
