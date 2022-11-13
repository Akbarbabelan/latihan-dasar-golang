package main

import "fmt"

func endApp() {
	message := recover()
	fmt.Println("Error dengan message", message)
	fmt.Println("Aplikasi seselai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
}
