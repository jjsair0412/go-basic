package main

import "fmt"

type shoppingBasket struct {
	cnt,
	price int
}

// 결제 함수
func (b shoppingBasket) purchase() int {
	return b.cnt * b.price
}

// 원본 수정(참조 전달 형식)
func (b *shoppingBasket) rePurchaseP(cnt, price int) {
	b.cnt += cnt
	b.price += price
	/*
		실제 로 발생하는 동작은 아래처럼 동작함
		*b.cnt += cnt
		*b.price += price

	*/

}

// 원본 수정 안됨(값 전달 형식)
func (b shoppingBasket) rePurchaseD(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

func main() {
	bs1 := shoppingBasket{
		3,
		5000,
	}

	fmt.Println("ex1 Total price : ", bs1.purchase())
	// 결과 : ex1 Total price :  15000

	// 원본 값 ( 3에 7 더하고 5000에 5000 더함)
	bs1.rePurchaseP(7, 5000)
	fmt.Println("ex2 Total price : ", bs1.purchase())
	// 결과 : ex1 Total price :  100000
	// 원본값이 수정됐음.
	// 주소값을 넣어주는 & 기호같은걸 넣어주지 않아도 , 자동으로 주소값이 넘어감

	bs1.rePurchaseD(10, 0)
	fmt.Println("ex3 Total price : ", bs1.purchase())
	// 결과 : ex1 Total price :  100000
	// 원본값이 수정되지 않음. 값 전달 방식으로 진행.
	// 받는게 포인터가 아니니까 값 전달 방식이 됨.
}
