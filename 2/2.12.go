package main

import (
	"log"
	"sort"
)

type User struct{
	FirstName string
	LastName string
}

func main() {
	map1 := make(map[string]string)
	map2 := make(map[int]int)
	map3 := make(map[string]User)

	var slice []int

	slice2 := []string{"satu", "dua", "tiga"}

	slice = append(slice, 4)
	slice = append(slice, 9)
	slice = append(slice, 2)
	sort.Ints(slice)

	map1["1"] = "Satu"
	map2[2] = 2

	me := User{
		FirstName: "Arthur",
		LastName: "Shelby",
	}

	map3["me"] = me

	log.Println(map1["1"], map2[2], map3["me"].FirstName, map3["me"].LastName)
	log.Println(slice, slice2)
}