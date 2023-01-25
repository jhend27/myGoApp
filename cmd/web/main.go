package main

import (
	"fmt"
	"myapp/cmd/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting application on port:", portNumber)

	_ = http.ListenAndServe(portNumber, nil)
}
