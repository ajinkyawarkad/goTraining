

// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

// Add imports.

import (
	"fmt"
)

func main() {

	// Declare a nil slice of integers.

	// Append numbers to the slice.

	// Display each value in the slice.

	// Declare a slice of strings and populate the slice with names.

	// Display each index position and slice value.

	// Take a slice of index 1 and 2 of the slice of strings.

	// Display each index position and slice values for the new slice.

	a := []int{}
	for i:=0 ; i<=10; i++ {
		a=append(a,i)
	}
	fmt.Println(a)
	
	
	names:= []string{"Ajinkya","Rahul","Pratik","Yash","Ganesh"}
	
	for i, val:=range  names {
	fmt.Println(i,val)
	
	}
	
	newSlice := []string{}
	newSlice = append(newSlice, names[1])
	newSlice = append(newSlice, names[2])
	
	for i, val:=range  newSlice{
	fmt.Println(i,val)
	
	}
	
	
	
}



// [0 1 2 3 4 5 6 7 8 9 10]


// 0 Ajinkya
// 1 Rahul
// 2 Pratik
// 3 Yash
// 4 Ganesh


// 0 Rahul
// 1 Pratik
