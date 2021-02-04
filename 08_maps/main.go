package main

import (
	"fmt"
)

func main() {
	// Define maps
	// emails := make(map[string]string)  // --> [key]value

	// Assign KV's
	// emails["Bob"] = "bob@gmail.com"
	// emails["Keanu"] = "keanu@gmail.com"

	// Declare map and add KVs 
	emails := map[string]string{ "Bob": "bob@gmail.com", "Keanu": "keanu@gmail.com"}
	fmt.Println(emails)
	fmt.Println(emails["Keanu"])

	delete(emails, "Bob")
	fmt.Println(emails)


}