package main

import "fmt"

// int두개를 매게변수로 받아서 int 한개를 리턴시키는 함수를
// totCost라는 이름으로 재 정의
type totCost func(int, int) int

// totCost 함수도 메게변수로 받음
// 이렇게하면 totCost를 받지 않으면 해당 함수가 실행되지 않기에 실수를 줄일 수 있음.
func describe(cnt, price int, fn totCost) {
	fmt.Printf("cnt : %d, price : %d, orderPrice : %d", cnt, price, fn(cnt, price))

}

func main() {
	var orderPrice totCost
	// orderPrice라는 변수에 totCost라는 함수 타입으로 정의해서 , 함수를 고정시킴
	orderPrice = func(cnt, price int) int {
		return (cnt * price) + 10000
	}

	describe(3, 5000, orderPrice)

}
