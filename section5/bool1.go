package main

import (
	"fmt"
)

func main() {
	var b1 bool = true
	var b2 bool = false
	b3 := true

	fmt.Println("ex 1 : ", b1)
	fmt.Println("ex 2 : ", b2)
	fmt.Println("ex 3 : ", b3)

	fmt.Println("ex4 : ", b3 == b1)

	fmt.Println("ex5 : ", b1 && b3)

	fmt.Println("ex6 : ", b2 || b3)

	fmt.Println("ex7 : ", !b1)

	/*
		// 암묵적 형 변환은 인정되지 않는다.
		b4 := 1

		if b4 {
			fmt.Println("ex 8 : TRUE ") 에러발생 . 1은 true가 아니다.
		}
	*/

}
