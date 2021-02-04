package main

import(
	"fmt"
	"strconv"
)
// Structs are pretty important in Go, they are like classes where we can define a structure.

// Methods:
// Pointer recievers --> methods that actually change things
// Value receviers --? methods that don't change anything, they do calculations 


// Define person struct.

type Person struct {
	// firstName string
	// lastName string
	// city string
	// gender string
	// age int

	firstName, lastName, city, gender string   // -- > shorter way
	age int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " year old!"
}

// hasBirthday method() (pointer reciver) -- > we want to change something

func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (Pointer receiver)

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "Male" {
		return
	} else {
		p.lastName = spouseLastName
	}
}
func main() {
	person1 := Person{
		firstName: "Sam",
		lastName: "Smith",
		city: "Boston",
		gender: "Female",
		age: 25,
		}
	// Another way to declare it.
	person2 := Person{"Keanu","Reeves", "Miami", "Male", 55}
		

		fmt.Println(person1.lastName)
		fmt.Println(person1.firstName)
		person1.hasBirthday()		   // --> Hello, my name is Sam Smith and I am 26 year old!v
		person1.getMarried("Williams") // --> Hello, my name is Sam Williams and I am 26 year old!
		person2.getMarried("Thompson")
		fmt.Println(person2.greet())   // --> Hello, my name is Keanu Reeves and I am 55 year old!
		fmt.Println(person1.greet())   // --> o/p : Hello, my name is Sam Smith and I am 25 year old!
}