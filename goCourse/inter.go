package main

import "fmt"

type Animals interface {
	says() string
	legs() int
}

type Dog struct {
	name  string
	color string
}

type Gorilla struct {
	name   string
	color  string
	teeths int
}

func main() {
	dogg := Dog{
		name:  "Moti",
		color: "Black",
	}

	gorr := Gorilla{
		name:   "Kong",
		color:  "Grey",
		teeths: 32,
	}

	printInfo(dogg)
	printInfo(gorr)

}

func printInfo(a Animals) {
	fmt.Println("Animal Says:", a.says(), "Having legs", a.legs())

}

func (d Dog) says() string {

	return "bhobho"
}

func (d Dog) legs() int {
	return 4
}

func (g Gorilla) legs() int {
	return 2
}

func (g Gorilla) says() string {

	return "HoHoHo"
}
