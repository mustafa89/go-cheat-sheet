package main

import (
	"fmt"
)

func main() {
	x := 5
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)  // --> the 5d is a place holder like {} in python
	} else {
		fmt.Println("%d is greater than %d\n", x, y)
	}

	// elseif

	color := "red"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is green")
	}

	// switch 

	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is green")
	}
}