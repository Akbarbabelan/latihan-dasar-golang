package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		x := 0
		for x != i {
			fmt.Print(" ")
			x++
		}
		for j := i; j < 5; j++ {
			fmt.Print("*")
		}

		fmt.Println()

	}
}
