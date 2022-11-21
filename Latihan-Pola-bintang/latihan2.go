package main

import (
	"fmt"
	"strings"
)

func main() {
	totalBaris := 0
	x := 7
	for x >= totalBaris {
		x--
		total := make([]string, x)
		fmt.Println(strings.Join(total, "*"))
	}
}
