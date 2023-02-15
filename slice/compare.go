package main

import "fmt"

func main () {
	slices1 := [] int {5, 8, 10, 11}
	slices2 := [] int {5, 8, 10, 11}	
    slices3 := [] int {5, 8, 10, 12}
	fmt.Println(compareSlices(slices1,slices2))
	fmt.Println(compareSlices(slices1,slices3))
}
//Die Funktion comparesSlices()vergleicht zwie Slices und gibt true zurück,wenn
//sie gleich sind.oder es gibt false ,wenn sie unterschiedlich sind
func compareSlices(slices1, slices2 []int) bool{
//erster vergleich:haben beide Slices die selbe Länge haben
if len(slices1) != len(slices2){
	return false 
}
//zweiter Vergleich:sind die Inhalte identisch...
 for i := range slices1 {
   if slices1[i] != slices2[i]{
	  return false
   }

 }
 //slices sind gleich:generieren keine Rückgabewerte ""return false"
 //returun true wird als rückgabewert für die Funktion compareSlices generiert
return true

}