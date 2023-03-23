package main

import (
	"fmt"
)

// return 매개변수에서 변수 선언 가능
func multiply(x, y int) (r1 int, r2 int) {
	// r1, r2를 명시적 선언하여 가독성 증대
	r1 = x * 10
	r2 = y * 10

	return r1, r2
}

func main() {
	a, b := multiply(10, 5)

	fmt.Println(a, b)
}
