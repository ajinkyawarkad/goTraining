package main

import (
	"fmt"
	"net/http"
)

const portNo string = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is Home")

}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 4)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is abt Page %d", sum))

}

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is Contact Page ")

}

func addValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/contact", Contact)

	fmt.Println("Starting connection on port No.: ", portNo)

	_ = http.ListenAndServe(portNo, nil)
}
