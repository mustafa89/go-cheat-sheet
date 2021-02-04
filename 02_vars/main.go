// Basic Types - As the name impies, int, float ...
	// int --> singned, unsigned, standard --> int int8 int16 int32 int64
	// uint --> uint uint8 uint16uint32 uint64 uintptr
	// byte - alias for uint8
	// rune --> alias for int32 
	// float --> float32 and float64
	// complex --> comples64 and comples128
	// Booleans --> True and false
	// strings --> string
	// constant --> values that should never change.
	// 
// Aggregate Types - These are arrays and structs
// Reference Types - These are your pointers and slices
// Interface Types - These are your standard interfaces

package main  // --> must start with a package
import "fmt"  // --> basic import
func main() {

	var name string = "Keanu Reeves" // --> the sting explicit declaration is not required here.
	var name1 = "Jason Bourne" // --> Go will detect it's a string
	name3 := "Jessica Alba" /// --> the := lets go interpret the type by itself (like interpreter languages do). can only be used in a function.
	num := 1.34543  // --> it's always going to be float64, if we need float32 we have to explicitly set it.
	fname, email := "magnus", "magnus@yahoo.com"
	fmt.Println(name, name1, name3)
	fmt.Printf("%T", num)  // get var type
}
