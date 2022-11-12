package main

import "fmt"

func main() {
	var name1 = "Kholifah"
	var name2 = "Kholifah"

	var result bool = name1 == name2
	fmt.Println(result)

	var value1 = 400
	var value2 = 600

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
	fmt.Println(value1 >= value2)
	fmt.Println(value1 <= value2)
}
