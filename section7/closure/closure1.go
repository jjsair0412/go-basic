package main

import "fmt"

func main() {
	multiply := func(x, y int) int {
		return x * y
	}

	r1 := multiply(5, 10)

	fmt.Println(r1)

	m, n := 5, 10 // 변수가 캡쳐되는 순간
	sum := func(c int) int {
		return m + n + c
	}

	r2 := sum(10)

	fmt.Println("ex2 : ", r2)
	// 결과 : ex2 :  25
}
