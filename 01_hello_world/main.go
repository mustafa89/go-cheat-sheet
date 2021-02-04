package main  // --> must start with a package
import "fmt"  // --> basic import

// unsigned integer with 8 bits
var umyint uint8

// signed integer with 8 bits
var myint int8

// Standard integer type, most commonly used.
var simple_int int

func main() {
	fmt.Println("Hello World")

    for i := 0; i < 129; i++ {
        myint += 1
    }
	fmt.Println(myint) // prin
	names()
	myFunc := anonTest()
	fmt.Println(myFunc())
	var employee Employee
    employee.Name = "Elliot"
    employee.UpdateName("Forbsey")
    employee.PrintName()
}




// converting different types to one another

// func type_conversion() {
// var int1 int
// var float1 float32

// float1 := float(int1)

// var int2 int
// var float2 float32

// int2 := int(float2)
// }

func multiReturns(first_name string, last_name string) (string, error) {
	return first_name + " " + last_name, nil
}


func names(){
	first_name := "Keanu"
	last_name := "Reeves"
	res, err := multiReturns(first_name, last_name)
	if err == nil {
		fmt.Println("Handle error case")
	}
	fmt.Println(res, err)
}


func anonTest() func() int {
	var x int

	return func () int{
		x++
		return x
	}
}


type Employee struct {
    Name string
}

func (e *Employee) UpdateName(newName string) {
    e.Name = newName
}

func (e *Employee) PrintName() {
    fmt.Println(e.Name)
}
