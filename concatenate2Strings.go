package main

import "fmt"

func main() {

    var str1 string
	str1 = "Welcome!"

	var str2 string
	str2 = "to nepal"

	fmt.Println("first string is : ", str1+str2)

	//2nd way 
	str3 := "hello "
	str4 := "guys"
	
	result := str3 + "for" + str4

	fmt.Println("second string is : ", result)

}
