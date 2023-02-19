package main

import "fmt"

func main() {
	const a, b int = 10, 99

	// 여러개 연속으로 선언하는데 , 각 다른 타입이 와도 문제가 없습니다.
	const c, d, e = true, 0.84, "TEST"

	const (
		x, y int16 = 50, 90
		i, k       = "Data", 7776
	)

	/*
		결과 :
		10 99 true 0.84 TEST
		50 90 Data 7776
	*/
	fmt.Println(a, b, c, d, e)
	fmt.Println(x, y, i, k)
}
