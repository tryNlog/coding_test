package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scanln(&n)
	n /= 4
	for i := 0; i < n; i++ {
		s += "long "
	}
	fmt.Print(s + "int")
}
