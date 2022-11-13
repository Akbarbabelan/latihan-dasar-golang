package main

import "fmt"

func main() {
	name := "Eko"
	counter := 0

	increment := func() {
		name := "Budi"
		fmt.Println("increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)

}
