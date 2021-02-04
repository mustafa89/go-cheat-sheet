package main

import(
	"fmt"
)

func main() {
	// Arrays
	// var fruitArr [2]string  // --> Delaring first 

	// Assing values

	// fruitArr[0] = "Apple"  // --> assigining values to the array
	// fruitArr[1] = "Orange"

	fruitArr := [2]string{"Apple", "Orange"} // Declaring and the also assigning values to the array
	fmt.Println(fruitArr)
	fmt.Println(fruitArr[0])

	fruitSlice := []string{"Applce", "orange", "grapes"} // --> slice is an array without the limit of items set
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))  // --> length of the array
	fmt.Println(fruitSlice[1:3])  // --< subset of an array

}