package main

import (
	"fmt"
	"strings"
)

func main() {
	totalBaris := 5
	x := 0
	for x <= totalBaris {
		x++
		total := make([]string, x)
		fmt.Println(strings.Join(total, "*"))
	}
}
