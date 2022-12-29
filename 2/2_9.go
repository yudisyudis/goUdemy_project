package main

import "log"

func main() {
	var myString string
	myString = "Blue"

	log.Println("myString is", myString)
	changeUsingPointer(&myString)
	log.Println("new color is", myString)
}

func changeUsingPointer(s *string)  {
	log.Println("s is", s)
	newValue := "Red"
	*s = newValue
}