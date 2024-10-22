package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanln(&a, &b)
	fmt.Scanln(&c)

	var hours int
	var minutes int
	hours = a + ((b + c) / 60)
	minutes = (b + c) % 60
	if hours >= 24 {
		hours -= 24
	}
	fmt.Print(hours, minutes)
}
