package main

import "fmt"

func main() {
	// 가격 총합을 저장할 변수(final_price)
	var final_price int
	// 물건의 종류(배열의 크기)를 저장할 변수(n)
	var n int
	// 물건의 가격과 수량을 저장할 변수(price, quantity)
	var price, quantity int
	// items의 가격 총합을 저장할 변수(total_price)
	var total_price int

	// 각 변수에 값을 입력받음
	fmt.Scanln(&final_price)
	fmt.Scanln(&n)
	// n만큼 loop하여 물건의 가격과 수량을 입력받은 후 둘을 곱하여 total_price에 더함
	for i := 0; i < n; i++ {
		fmt.Scanln(&price, &quantity)
		total_price += price * quantity
	}

	// total_price가 final_price와 같으면 "Yes"를 출력하고, 아니면 "No"를 출력
	if total_price == final_price {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}
