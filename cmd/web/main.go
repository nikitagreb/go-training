package main

import (
	"fmt"
	"github.com/nikitagreb/go-training/pkg/handlers"
	"net/http"
)

const portNumber string = ":8888"

// main is the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
