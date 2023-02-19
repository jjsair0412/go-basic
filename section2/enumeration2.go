package main

import "fmt"

func main() {
	// iota 예약어 사용
	const (
		A = iota
		B
		C
	)

	// 결과 : 0 1 2
	fmt.Println(A, B, C)

	const (
		D = iota * 10
		E
		F
	)

	// 결과 : 0 10 20
	fmt.Println(D, E, F)
}
