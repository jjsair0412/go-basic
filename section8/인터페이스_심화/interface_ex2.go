package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{} = 15

	b := a
	c := a.(int) // 다 int지만 이렇게 타입을 변환해서 사용해야 합니다.

	fmt.Println("ex1 : ", a)
	fmt.Println("ex1 : ", reflect.TypeOf(a))
	fmt.Println("ex1 : ", b)
	fmt.Println("ex1 : ", reflect.TypeOf(b))
	fmt.Println("ex1 : ", c)
	fmt.Println("ex1 : ", reflect.TypeOf(c))

	fmt.Println()

	// 데이터타입이 맞는지 true, false로 검사
	if v, ok := a.(int); ok {
		fmt.Println("ex2 : ", v, ok)
		// 결과 : ex2 :  15 true
	}

}
