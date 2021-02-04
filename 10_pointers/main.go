package main

import (
	"fmt"
)

func main() {
	a :=  5
	b := &a // --> assigning b to pointer of a

	fmt.Println(a,b)
	fmt.Printf("%T\n", b)  // o/p --> *int
	fmt.Println(*&a, *b) // o/p --> 5 5

	// Change value with pointer

	*b = 10
	fmt.Println(a)

	// Passing address instead of data is faster. So we use pointers to do this.
	// Also used to change a value in a different location
}