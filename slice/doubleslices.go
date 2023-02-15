package main

import "fmt"

func main (){
 doubleNum := []int{1, 2, 3, 4, 5}
 doubleSlicesElements(doubleNum)  //funktion aufrufen 
 fmt.Println ("Das ist das doppelte Wert von doubleNum:" ,doubleNum)
}

func doubleSlicesElements(doubleNum []int) {
for i := range doubleNum{
doubleNum[i] *= 2
}

}