package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for x := i; x < 5-1; x++ {
			fmt.Print(" ")

		}
		for j := i; j < i+(i+1); j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
