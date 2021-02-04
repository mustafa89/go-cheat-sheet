package main

import(
	"fmt"
)

func main() {
	ids := []int{41,62,53,41,56,76}

	 // loop through

	 for i, id := range ids {   // --> for <index>, <actual placeholder> := range <what we're looping through
		 fmt.Printf("%d - ID: %d\n", i ,id)
	 }

	// Output
	// 0 - ID: 1
	// 1 - ID: 2
	// 2 - ID: 3
	// 3 - ID: 4
	// 4 - ID: 5
	// 5 - ID: 6

	for _, id := range ids {   // --> not using an index
		 fmt.Printf("ID: %d\n", id)

	 }

	// outputs
	// ID: 41
	// ID: 62
	// ID: 53
	// ID: 41
	// ID: 56
	// ID: 76

	// Sum example
	 num := 0
	for _, sum := range ids {
		num += sum
	}
	fmt.Println("Sum", num)

	// outputs
	// Sum 329


	// maps
	emails := map[string]string{ "Bob": "bob@gmail.com", "Keanu": "keanu@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n",k ,v)

	// outputs
	// Bob: bob@gmail.com
	// Keanu: keanu@gmail.com
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
	// Outputs
	// Name: Keanu
	// Name: Bob
}