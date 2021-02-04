package main

import(
	"fmt"
)

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// short method
	for i := 0; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	// Fizzbuzz
	for i := 0; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i % 5 == 0 {
			fmt.Println("buzz")
		} else if i % 3 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(i)
		}
	} 
}