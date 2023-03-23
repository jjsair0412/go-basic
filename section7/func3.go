package main

import "fmt"

// 매개변수 타입이 int로 동일하기에 int 하나 생략
// return할 변수가 두개라서 int, int 두개 작성
func multiply(x, y int) (int, int) {
	// x * 10은 첫번쨰 int , y * 10은 두번째 int
	return x * 10, y * 10
}

func arrayMultiply(a, b, c, d, e int) (int, int, int, int, int) {
	return a * 1, b * 2, c * 3, d * 4, e * 5
}

func main() {
	// a에 100 담기고 , b에 50 담김.
	a, b := multiply(10, 5)

	// c := multiply(10, 5) 함수 리턴타입 개수 안맞아서 에러발생
	c, _ := multiply(10, 5) // _로 뒤에 리턴되는 y * 10 생략
	_, d := multiply(10, 5)

	fmt.Println(a, b)
	fmt.Println(c)
	fmt.Println(d)

	/*
		결과 :
		100 50
		100
		50
	*/

	x1, x2, x3, x4, x5 := arrayMultiply(1, 2, 3, 4, 5)
	y1, _, y3, _, y5 := arrayMultiply(1, 2, 3, 4, 5)

	fmt.Println(x1, x2, x3, x4, x5)
	fmt.Println(y1, y3, y5)
	/*
		결과 :
		1 4 9 16 25
		1 9 25
	*/
}
