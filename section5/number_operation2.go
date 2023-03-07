package main

import (
	"fmt"
)

func main() {
	var n1 uint8 = 125
	var n2 uint8 = 90

	fmt.Println("ex1 : ", n1+n2)
	fmt.Println("ex1 : ", n1-n2)
	fmt.Println("ex1 : ", n1*n2)
	fmt.Println("ex1 : ", n1/n2)
	fmt.Println("ex1 : ", n1%n2)
	fmt.Println("ex1 : ", n1<<2) // 왼쪽으로 2비트 이동
	fmt.Println("ex1 : ", n1>>2) // 오른쪽으로 2비트 이동
	fmt.Println("ex1 : ", ^n1)   // 0 > 1, 1 > 0 으로 변경

	/*
		결과 :
		ex1 :  215
		ex1 :  35
		ex1 :  242
		ex1 :  1
		ex1 :  35
		ex1 :  244
		ex1 :  31
		ex1 :  130

	*/

	var n3 int = 12
	var n4 float32 = 8.2
	var n5 uint16 = 1024
	var n6 uint32 = 120000

	// fmt.Println("ex3 : ", n3+n4)
	fmt.Println("ex3 : ", float32(n3)+n4)
	fmt.Println("ex3 : ", n3+int(n4)) // 결과 변경됨. int로 형변환시 소수점 삭제됨

	fmt.Println("ex3 : ", n5+uint16(n6))
	fmt.Println("ex3 : ", uint32(n5)+n6)
}
