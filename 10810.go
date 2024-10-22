package main

import "fmt"

func main() {

	var num, m int
	fmt.Scanln(&num, &m)
	var i, j, k int
	baskets := make([]int, num)
	for x := 0; x < m; x++ {
		fmt.Scanln(&i, &j, &k)
		for y := i - 1; y < j; y++ {
			baskets[y] = k
		}
	}
	for z := 0; z < num; z++ {
		fmt.Printf("%d ", baskets[z])
	}
}
