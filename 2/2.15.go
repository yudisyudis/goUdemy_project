package main

import "fmt"

type Animal interface {
	Bunyi() string
	Kaki() int
}

type Dog struct {
	Nama  string
	Jenis string
}

type Gorila struct {
	Nama  string
	Berat int
}

func main() {
	dog := Dog{
		Nama:  "Bleki",
		Jenis: "Helder",
	}

	gorila := Gorila{
		Nama:  "Hendri",
		Berat: 100,
	}
	printInfo(&dog)
	printInfo(&gorila)
}

func printInfo(a Animal) {
	fmt.Println("Hewan ini memiliki bunyi", a.Bunyi(), "dan jumlah kaki", a.Kaki())
}

func (d *Dog) Bunyi()string {
	return "guk"
}
func (d *Dog) Kaki()int {
	return 4
}
func (d *Gorila) Bunyi()string {
	return "aum"
}
func (d *Gorila) Kaki()int {
	return 2
}