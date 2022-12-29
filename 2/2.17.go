package main

import (
	"log"

	"github.com/yudisyudis/go_udemy/helpers"
)


func main() {
	intChan := make(chan int)
	defer close(intChan)

	go hitungNilai(intChan)

	nilai := <- intChan
	log.Println(nilai)
}

const numPool = 10
func hitungNilai(intChan chan int) {
	randomNumber := helpers.AcakNomor(numPool)
	intChan <- randomNumber
}