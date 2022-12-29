package main

import (
	"log"

	helpers "github.com/yudisyudis/go_udemy/helpers"
)

func main() {
	log.Println("Hello Go")

	var myVar helpers.acakAngka
	myVar.TypeName = "Any Name"

	log.Println(myVar.TypeName)
}