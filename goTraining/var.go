// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

// Add imports
import (
	"fmt"
)

// main is the entry point for the application.
func main() {
	var a string
	var b int 
	var c bool 

	// Declare variables that are set to their zero value.
	fmt.Printf("%T[%v], %T[%v], %T[%v] \n", a, a, b, b, c, c)
	// Display the value of those variables.

	// Declare variables and initialize.
	d,e,f := "asd", 1, true  
	// Using the short variable declaration operator.

	// Display the value of those variables.
	fmt.Printf("%v %v %v \n",d,e,f)
	// Perform a type conversion.
	var g float32 = 3.14;
	// Display the value of that variable.
	fmt.Printf("%v %v\n", g, int32(g))
}
