
// Declare and initialize a variable of type int with the value of 20. Display
// the _address of_ and _value of_ the variable.
//
// Declare and initialize a pointer variable of type int that points to the last
// variable you just created. Display the _address of_ , _value of_ and the
// _value that the pointer points to_.
package main

// Add imports.


import (
	"fmt"
)


func main() {

	a := 20
	fmt.Println("the value of_a",a)
	fmt.Println("the address of_a",&a)

	b := &a

	fmt.Println("the value of_b",b)
	fmt.Println("the address of_b",&b)

	// Declare an integer variable with the value of 20.

	// Display the address of and value of the variable.

	// Declare a pointer variable of type int. Assign the
	// address of the integer variable above.

	// Display the address of, value of and the value the pointer
	// points to.
}
