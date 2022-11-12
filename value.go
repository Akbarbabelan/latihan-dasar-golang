package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}
func main() {
	sayGoodBye := getGoodBye

	result := sayGoodBye("Akbar")
	fmt.Println(result)
	fmt.Println(getGoodBye("Akbar"))

}
