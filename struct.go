package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var akbar Customer
	akbar.Name = "Akbar"
	akbar.Address = "Indonesia"
	akbar.Age = 22

	fmt.Println(akbar.Name)
	fmt.Println(akbar.Address)
	fmt.Println(akbar.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     22,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 22}
	fmt.Println(budi)
}
