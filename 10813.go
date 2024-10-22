package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanln(&n, &m)
	baskets := make([]int, n)

	for i := 0; i < n; i++ {
		baskets[i] = i + 1
	}

	for j := 0; j < m; j++ {
		var x, y int
		var temp int
		fmt.Scanln(&x, &y)
		temp = baskets[x-1]
		baskets[x-1] = baskets[y-1]
		baskets[y-1] = temp
	}

	for k := 0; k < n; k++ {
		fmt.Printf("%d ", baskets[k])
	}
}
