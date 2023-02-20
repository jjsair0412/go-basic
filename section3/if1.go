package main

import "fmt"

func main() {
	// 제어문 ( 조건문 )

	var a int = 20
	b := 20

	if a >= 15 {
		fmt.Println("15 이상입니다.")
	}

	if b >= 25 {
		fmt.Println("25 이상입니다.")
	}

	/* 에러 발생 1

	if b >= 25
	{

	}
	*/

	/* 에러 발생
	if b >= 25
		fmt.Println("25 이상")
	*/

	if c := 40; c >= 35 {
		fmt.Println("35 이상")
	}
}
