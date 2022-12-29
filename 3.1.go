package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8888"

func Home(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "This is Home page")
}

func About(w http.ResponseWriter, r *http.Request)  {
	sum := addvalues(2, 2)
	_,_ = fmt.Fprintf(w, fmt.Sprintf("This is page abour and sum 2 + 2 is %d", sum))
}

func addvalues(x, y int)int  {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request)  {
	f, err := divideValues(100.0, 0.0)
	if err != nil{
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, f))
}

func divideValues(x, y float32)(float32, error)  {
	if y == 0{
		err := errors.New("cannot divide by zero")
		return 0,err
	}
	result := x/y
	return result, nil
}

func main() {

	http.HandleFunc("/", Home)

	http.HandleFunc("/about", About)

	http.HandleFunc("/divide", Divide)


	// http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
	// 	n, err := fmt.Fprintf(w, "hello world")
	// 	if err != nil{
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(fmt.Sprintf("bytes: %d", n))
	// })

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_=http.ListenAndServe(":8888", nil)
}