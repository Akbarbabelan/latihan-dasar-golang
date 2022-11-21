package main

import (
	"fmt"
	"strings"
)

func main() {
	totalBaris := 10 // jika ingin menambah / mengurangi total baris ubah disini

	dividingLine("Latihan 1")
	x := 0
	for x < totalBaris {
		x++
		total := make([]string, x+1)
		fmt.Println(strings.Join(total, "*"))
	}

	dividingLine("Latihan 2")
	y := totalBaris
	for y != 0 {
		total := make([]string, y+1)
		fmt.Println(strings.Join(total, "*"))
		y--
	}

	dividingLine("Latihan 3")
	z := totalBaris
	for z != 0 {
		total := make([]string, z+1)
		space := make([]string, totalBaris-z+1)
		fmt.Print(strings.Join(space, " "))
		fmt.Println(strings.Join(total, "*"))
		z--
	}

	dividingLine("Latihan 4")
	w := 0
	for w < totalBaris {
		w++
		total := make([]string, w+1)
		space := make([]string, totalBaris-w+1)
		fmt.Print(strings.Join(space, " "))
		fmt.Println(strings.Join(total, "*"))
	}

	dividingLine("Latihan 5")
	v := 0
	for v < totalBaris {
		v++
		totalLeft := make([]string, v+1)
		totalRight := make([]string, v)
		space := make([]string, totalBaris-v+1)
		fmt.Print(strings.Join(space, " "))
		fmt.Print(strings.Join(totalLeft, "*"))
		fmt.Println(strings.Join(totalRight, "*"))
	}
}

func dividingLine(label string) {
	equalString := strings.Join(make([]string, 50), "=")
	fmt.Println(fmt.Sprintf("%v %v %v", equalString, label, equalString))
}
