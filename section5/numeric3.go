package main

import (
	"fmt"
)

func main() {
	var num1 float32 = 0.14
	var num2 float32 = .7 // .7로 하면 , 0.7로 인식된다.
	var num3 float32 = 442.0378373
	var num4 float32 = 10.0

	// 지수표기법
	var num5 float32 = 14e6
	var num6 float32 = .156875e+3
	var num7 float64 = 5.32521e-10

	fmt.Println("ex1 : ", num1)
	fmt.Println("ex2 : ", num2)
	fmt.Println("ex3 : ", num3)
	fmt.Println("ex4 : ", num4)
	fmt.Println("ex5 : ", num4-0.1)
	fmt.Println("ex6 : ", float32(num4-0.1)) // 형변환 .
	fmt.Println("ex7 : ", float64(num4-0.1)) // float64로 형변환 . 값이 이상해짐
	fmt.Println("ex8 : ", num5)
	fmt.Println("ex9 : ", num6)
	fmt.Println("ex10 : ", num7)

	/*
		결과 :
		ex1 :  0.14
		ex2 :  0.7
		ex3 :  442.03784
		ex4 :  10
		ex5 :  9.9
		ex6 :  9.9
		ex6 :  9.9
		ex7 :  9.899999618530273
		ex8 :  1.4e+07
		ex9 :  156.875
		ex10 :  5.32521e-10
	*/
}
