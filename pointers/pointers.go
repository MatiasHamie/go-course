package main

import "fmt"

// &variable te da el address como 0x748ecadcc0c0
// *variable te da el valor que tiene guardado esa variable

// ? A Pointer's Null Value
// All values in Go have a so-called "Null Value" - i.e., the value that's set as a default if no value is assigned to a variable.
// For example, the null value of an int variable is 0. Of a float64, it would be 0.0. Of a string, it's "".
// For a pointer, it's nil - a special value built-into Go.
// nil represents the absence of an address value - i.e., a pointer pointing at no address / no value in memory.
func main() {
	age := 32

	var agePointer *int
	agePointer = &age

	fmt.Println("Age", age)
	fmt.Println("AgePointer Address", agePointer)
	fmt.Println("AgePointer Value", *agePointer)
	editAgeToAdultYears(agePointer)
	fmt.Println("Age after calling the function", age)
}

func editAgeToAdultYears(age *int) {
	// return *age - 18
	// esto es para demostrar que con un puntero se puede modificar la variable inicial desde otro lado
	*age = *age - 18
}
