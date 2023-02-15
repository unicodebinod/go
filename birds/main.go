package main

import (
	"fmt"
)

func main (){

	v := Vogel{name : "Owl", alter : 5 }
	fmt.Println("name der Vogel ist:" , v.Name() + "\n alter der Vogel ist :" , v.Alter())
}