package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}
func main() {
	firstName := "Kholifah"
	sayHelloTo(firstName, "Akbar")
	sayHelloTo("Budi", "Nugroho")
}
