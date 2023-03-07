package main

import (
	"fmt"
	"math"
)

func main() {
	var n1 uint8 = math.MaxUint8
	var n2 uint16 = math.MaxUint16
	var n3 uint32 = math.MaxUint32
	var n4 uint64 = math.MaxUint64

	// 최댓값
	fmt.Println("ex1 : ", n1)
	fmt.Println("ex2 : ", n2)
	fmt.Println("ex3 : ", n3)
	fmt.Println("ex4 : ", n4)
	/*
		결과 :
		ex1 :  255
		ex1 :  65535
		ex1 :  4294967295
		ex1 :  18446744073709551615
	*/

	fmt.Println("ex1 : ", math.MaxUint8)
	fmt.Println("ex2 : ", math.MaxUint16)
	fmt.Println("ex3 : ", math.MaxUint32)

	fmt.Println("ex1 : ", math.MaxFloat32)
	fmt.Println("ex2 : ", math.MaxFloat64)

	n5 := 100000       // data type : int
	n6 := int16(10000) // 형변환됨 . 데이터 타입은 int16
	n7 := int8(100)

	// fmt.Println("ex2 : ", n5 + n6) 에러 발생 . 다른타입이라.

	fmt.Println("ex2 : ", n5+int(n6))
	fmt.Println("ex3 : ", n7+int8(n5))

	fmt.Println("ex2 : ", n6 > int16(n7))

	fmt.Println("ex2 : ", n6-int16(n7) > 5000)
}
