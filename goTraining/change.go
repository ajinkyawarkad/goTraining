// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
package main

import (
	"fmt"
)

// Declare a type named user.
type user struct{
	name string;
	email string;
	age int16
}


// Create a function that changes the value of one of the user fields.


// func funcName( point /* add pointer parameter, add value parameter */ , val) {

// 	fmt.Println(point)

// 	fmt.Println(val)

// 	// Use the pointer to change the value that the
// 	// pointer points to.

	
// }

func main() {

	// Create a variable of type user and initialize each field.

	ajinkya := user{"ajinkya","ajinkyaw@mkcl.org",22}
	
	fmt.Println(ajinkya)


	// Display the value of the variable.

	// Share the variable with the function you declared above.

	

	// fmt.Println(funcName(*ajinkya , ajinkya))


	// Display the value of the variable.
}
