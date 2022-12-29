package main

import "fmt"

func main() {
	fmt.Println("Hello, World")

	var whatToSay string

	whatToSay = "Goodbye"
	fmt.Println(whatToSay)

	var i int
	i = 7
	fmt.Println(i)

	whatWasSaid, otherWord := saySomething()
	fmt.Println("the function returned", whatWasSaid, otherWord)
}

func saySomething() (string, string){
	return "something", "else"
}