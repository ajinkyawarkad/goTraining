
package main


import (
	"fmt"
)

type user struct{
	name string;
	email string;
	age int16
}

// Add user type and provide comment.

func main() {

	// Declare variable of type user and init using a struct literal.
	ajinkya := user{"ajinkya", "ajinkyaw@mkcl.org", 22}
	// Display the field values.


	fmt.Println( ajinkya)

	
	
	// Declare a variable using an anonymous struct.
	anonymous := struct{
		purpose string;
	}{"dummyString"}
	
	// Display the field values.
	fmt.Println(anonymous)
}
