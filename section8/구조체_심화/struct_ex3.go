package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) CalculateD(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func (a *Account) CalculateP(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func main() {
	// 정리 : 구조체 인스턴스 값 변경시 -> 포인터 전달, 보통의 경우 -> 값 전달
	kim := Account{
		"245-901",
		100000,
		0.015,
	}

	lee := Account{
		"245-901",
		100000,
		0.015,
	}

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", lee)

	/*
		결과 :
		ex1 :  {245-901 100000 0.015}
		ex1 :  {245-901 100000 0.015
	*/
	fmt.Println()

	lee.CalculateD(100000)
	kim.CalculateP(100000)

	fmt.Println("ex2 : ", int(lee.balance))
	fmt.Println("ex2 : ", int(kim.balance))

	/*
		결과 :
		ex2 :  100000
		ex2 :  201500
		값의 의한 전달은 , 실제 객체(kim 인스턴스) 의 주소값을 참조할 수 없기 때문에
		값이 변하지 않았지만 ,

		포인터 리시버는(참조에 의한 전달) 실제 객체의 주소값을  참조할 수 있기 때문에
		값이 변합니다.
	*/
}
