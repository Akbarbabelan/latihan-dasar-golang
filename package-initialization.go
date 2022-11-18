package main

import (
	"fmt"

	"latihan-.git/latihan-dasar-golang/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
