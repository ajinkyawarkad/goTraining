package main

import "fmt"

func main() {
	name := "Ajinkya"

	fmt.Println("name before func execution", name)

	changeWithPointer(&name)

	fmt.Println("name after func execution", name)

}

func changeWithPointer(s *string) {
	newName := "Ram"

	*s = newName

}
