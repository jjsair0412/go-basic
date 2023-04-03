package main

import (
	"fmt"
)

// int를 cnt라는 이름으로 재 정의
type cnt int

func testConvertT(i int) {
	fmt.Println("ex3 : (Default Type) : ", i)
}

func testConvertD(i cnt) {
	fmt.Println("ex3 : (Custom Type) : ", i)
}

func main() {
	a := cnt(15)
	fmt.Println("ex1 : ", a)

	var b cnt = 15
	fmt.Println("ex1 : ", b)

	// testConvertT(b) 에러 발생. int를 넣어야지 왜 cnt 타입으로 넣냐 !! 라는 에러..
	// 에러가 발생하지 않으려면 아래처럼 형변환 해야 함.
	testConvertT(int(b))
	testConvertD(a)

	/*
	   결과 :
	   ex1 :  15
	   ex1 :  15
	   ex3 : (Default Type) :  15
	   ex3 : (Custom Type) :  15
	*/
}
