package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanln(&a, &b, &c)
	// 같은 눈이 3개가 나오면 10,000원+(같은 눈)×1,000원의 상금을 받게 된다.
	if a == b && b == c {
		fmt.Print(10000 + a*1000)
	}
	// 같은 눈이 2개만 나오는 경우에는 1,000원+(같은 눈)×100원의 상금을 받게 된다.
	if a == b && b != c {
		fmt.Print(1000 + a*100)
	}
	if a == c && b != c {
		fmt.Print(1000 + a*100)
	}
	if b == c && a != b {
		fmt.Print(1000 + b*100)
	}
	// 모두 다른 눈이 나오는 경우에는 (그 중 가장 큰 눈)×100원의 상금을 받게 된다.
	if a != b && b != c && a != c {
		if a > b && a > c {
			fmt.Print(a * 100)
		}
		if b > a && b > c {
			fmt.Print(b * 100)
		}
		if c > a && c > b {
			fmt.Print(c * 100)
		}
	}
}
