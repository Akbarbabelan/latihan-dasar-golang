package main

import "fmt"

func sumALL(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	total := sumALL(10, 90, 30, 50, 40)
	fmt.Println(total)
	slice := []int{10, 20, 30, 40, 50}
	total = sumALL(slice...)
	fmt.Println(total)
}
