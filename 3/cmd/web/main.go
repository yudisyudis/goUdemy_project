package main

import (
	"fmt"
	"net/http"

	"github.com/yudisyudis/go-udemy/pkg/handlers"
)

const portNumber = ":8800"

func main() {

	http.HandleFunc("/", handlers.Home)

	http.HandleFunc("/about", handlers.About)


	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_=http.ListenAndServe(portNumber, nil)
}