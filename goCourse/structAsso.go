package main

import "fmt"

type Users struct {
	firstName string
	lastName  string
}

func main() {

	user1 := Users{
		firstName: "Ajinkya",
		lastName:  "Warkad",
	}

	user2 := Users{
		firstName: "Ram",
		lastName:  "Rathod",
	}

	fmt.Println("details1:", user1.printDetails())
	fmt.Println("details2:", user2.printDetails())

}

func (u *Users) printDetails() (string string) {
	return u.firstName
	return u.lastName
}
