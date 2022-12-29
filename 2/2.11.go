package main

import "log"

type myStruct struct {
	FirstName string
}

//receiver 
func (m *myStruct) printFirstName()string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Shelby",
	}

	log.Println("myVar is", myVar.printFirstName())
	log.Println("myVar2 is", myVar2.printFirstName())
}