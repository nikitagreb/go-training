package main

import (
	"fmt"
	"net/http"
)

const portNumber string = ":8888"

// main is the main application function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

// Home is the about page handler
func Home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))
}

func Divide(w http.ResponseWriter, r *http.Request) {

}

// addValues adds two integers and return the sum
func addValues(x, y int) int {
	return x + y
}

func divideValues(x, y float32) (float32, error) {
	result := x / y

	return result, nil
}
