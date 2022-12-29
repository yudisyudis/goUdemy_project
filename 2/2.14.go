package main

import "log"

func main() {
	// for i := 0; i <= 100; i++ {
	// 	log.Println(i)
	// }
	
	// WITH SLICE
	// animals := []string{"cow", "cat", "dog"}
	
	// //jika tidak ingin menampilkan index
	// for _, animal := range animals{
	// 	log.Println(animal)
	// }
	
	// WITH MAP
	// animals := make(map[string]string)
	// animals["dog"] = "dogi"
	// animals["cat"] = "cattie"

	// for jenis, nama := range animals{
	// 	log.Println(jenis, nama)
	// }

	// WITH STRUCT
	type User struct{
		FirstName string
		LastName string
		Email string
		Age int
	}

	var users []User
	users = append(users, User{"Thomas", "Shelby", "tom@peakyblinders.com", 30})
	users = append(users, User{"Arthur", "Shelby", "arthur@peakyblinders.com", 32})
	users = append(users, User{"John", "Shelby", "john@peakyblinders.com", 28})

	for _, name := range users{
		log.Println(name.FirstName, name.LastName, name.Email, name.Age)
	}
}