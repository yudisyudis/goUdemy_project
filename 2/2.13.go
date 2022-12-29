package main

import "log"

func main() {
	//IF STATEMENT
	// num := 50
	// isTrue := false

	// if num > 99 && isTrue {
	// 	log.Println("benar semua")
	// }else if num<100 && isTrue{
	// 	log.Println("benar satu")
	// }else if num > 99 && !isTrue{
	// 	log.Println("benar satu")
	// }else{
	// 	log.Println("salah semua")
	// }

	//SWITCH STATEMENT

	myVar := "gajah"

	switch myVar{
	case "dog":
		log.Println("var is dog")

	case "cat":
		log.Println("var is cat")

	case "bird":
		log.Println("var is bird")
	
	default:
		log.Println("yang penting binatang")
	}
}