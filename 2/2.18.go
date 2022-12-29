package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HasGun    bool   `json:"has_gun"`
}

func main() {
	myJson := `
	[
		{
			"first_name": "Thomas",
			"last_name": "Shelby",
			"has_gun": true
		},
		{
			"first_name": "John",
			"last_name": "Shelby",
			"has_gun": false
		}
	]`

	var unmar []Person

	err := json.Unmarshal([]byte(myJson), &unmar)
	if err != nil{
		log.Println("Error unmarshalling", err)
	} 
	log.Println("unmarshalled: %v", unmar)

	var mySlice []Person

	var m1 Person
	m1.FirstName = "Arthur"
	m1.LastName = "Shelby"
	m1.HasGun = false

	mySlice = append(mySlice, m1)

	newJson, err := json.MarshalIndent(mySlice, "", " ")
	if err != nil{
		log.Println("error marshalling", err)
	}
	fmt.Println(string(newJson))
}