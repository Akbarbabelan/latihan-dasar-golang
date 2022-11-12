package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpKholifah NoKTP = "3242343376677"
	var marriedStatus Married = false
	fmt.Println(noKtpKholifah)
	fmt.Println(marriedStatus)
}
