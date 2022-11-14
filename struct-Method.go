package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Helo", name, "My Name is", customer.Name)
}
func (a Customer) sayHuuu() {
	fmt.Println("Huuu from", a.Name)
}
func main() {
	var akbar Customer
	akbar.Name = "Akbar"
	akbar.Address = "Indonesia"
	akbar.Age = 22

	akbar.sayHi("Joko")
	akbar.sayHuuu()
}
